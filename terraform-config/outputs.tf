output "api_response" {
  description = "Raw JSON returned by the mock server"
  value       = data.jsonendpoint_fetch.item.response
}


# output "create_todo_response" {
#   value = jsondecode(jsonencode(jsonendpoint_item.create_todo.response))
# }
