provider "jsonendpoint" {
  base_url = "http://localhost:9000"  # Replace with your actual base URL
}

resource "jsonendpoint_item" "item" {
  endpoint = "/api/item/123"
  payload  = jsonencode({
    name  = "Terraform"
    # value = 100
  })
}

# # resource "jsonendpoint_item" "item" {
# #   endpoint = "/api/item/123"
# #   payload  = jsonencode({
# #     name  = "Terraform Updated"
# #     value = 200
# #   })
# # }
#
#
# # resource "jsonendpoint_item" "item" {
# #   endpoint = "/api/item/123"
# #   payload  = jsonencode({
# #     name  = "Terraform"
# #     value = 100
# #   })
# # }
#
data "jsonendpoint_fetch" "item" {
  endpoint = "/api/item/123"
}
# -----------------------------------

# This below code is to test the provider with other server

# ------------------------------------

# provider "jsonendpoint" {
# base_url = "http://localhost:9000"  # Ensure this matches your server's base URL
# }
#
# # Resource to create a Todo item
# resource "jsonendpoint_item" "create_todo" {
# endpoint = "/todo"
# payload  = jsonencode({
# title     = "New Todo Item"
# completed = false
# })
# }
#
# # Data source to fetch all Todo items
# data "jsonendpoint_fetch" "get_todos" {
# endpoint = "/todo"
