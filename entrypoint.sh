#!/bin/bash
set -e

# Optionally, start the mockserver (uncomment the line below if you want it to start in the background)
# echo "Starting mockserver..."
# /app/mockserver &

# Run the Terraform provider as the main process
echo "Starting the Terraform provider..."
exec /app/terraform-provider-jsonendpoint
