package provider

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceJSONEndpoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,

		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "API path to send requests to, relative to base_url.",
			},
			"payload": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "JSON payload for POST/PUT operations.",
			},
			"response": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Raw response body from the endpoint.",
			},
		},
	}
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*Config)
	endpoint := d.Get("endpoint").(string)
	payload := d.Get("payload").(string)
	url := config.BaseURL + endpoint

	resp, err := http.Post(url, "application/json", bytes.NewBufferString(payload))
	if err != nil {
		return diag.Errorf("POST failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return diag.Errorf("Failed to read POST response: %s", err)
	}

	// Set ID and response from POST
	d.SetId(endpoint)
	if err := d.Set("response", string(body)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*Config)
	id := d.Id()
	url := config.BaseURL + id

	resp, err := http.Get(url)
	if err != nil {
		return diag.Errorf("GET failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return diag.Errorf("Failed to read GET response: %s", err)
	}

	// Update response field with GET result
	if err := d.Set("response", string(body)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*Config)
	url := config.BaseURL + d.Id()
	payload := d.Get("payload").(string)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBufferString(payload))
	if err != nil {
		return diag.Errorf("Failed to create PUT request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return diag.Errorf("PUT failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return diag.Errorf("Failed to read PUT response: %s", err)
	}

	// Set updated response from PUT
	if err := d.Set("response", string(body)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*Config)
	dId := d.Id()
	url := config.BaseURL + dId

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return diag.Errorf("Failed to create DELETE request: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return diag.Errorf("DELETE failed: %s", err)
	}
	defer resp.Body.Close()

	// Gracefully handle 404
	if resp.StatusCode == http.StatusNotFound {
		log.Printf("[WARN] Resource %s not found during delete; assuming already deleted", dId)
		d.SetId("")
		return nil
	}

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return diag.Errorf("DELETE failed: %d - %s", resp.StatusCode, string(body))
	}

	d.SetId("")
	return nil
}
