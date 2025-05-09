# Terraform Provider: `jsonendpoint`

A custom Terraform provider for interacting with JSON-based HTTP endpoints. Includes a mock server for testing and example Terraform configurations.


## 🚀 Features

- **Resource Management**  
  Define, create, read, update, and delete resources exposed via JSON APIs.

- **Mock Server**  
  Included mock server (`mockserver/main.go`) for local testing and development.

- **Terraform Configurations**  
  Example Terraform configurations in the `terraform-config` folder to demonstrate usage.

## 🛠️ Installation

### 1. Clone the Repository

```bash
git clone https://github.com/poornatejav/tf-provider-jsonendpoint.git
cd tf-provider-jsonendpoint
```

### 2. Build the Provider

```bash
go build -o terraform-provider-jsonendpoint
```

### 3. Install the Provider

```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/USERNAME/jsonendpoint/1.0.0/linux_amd64/
mv terraform-provider-jsonendpoint ~/.terraform.d/plugins/registry.terraform.io/USERNAME/jsonendpoint/1.0.0/linux_amd64/
```
📝 Replace linux_amd64 with your system's OS/architecture if different.

### 4. Navigate to the Mock Server Directory

```bash
cd mockserver

go run main.go
```
### 5. Navigate to the Terraform Configuration Directory

First, navigate to the directory containing your Terraform configuration files.

```bash
cd terraform-config
```

Initialize Terraform

```bash
terraform init
```

Apply the Configuration

```bash
terraform apply
```

Verify the Resource

```bash
# Terraform will show what actions it will take
# Type 'yes' to apply the configuration
```
