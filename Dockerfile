FROM golang:1.21-buster

# Install necessary tools and Terraform
ENV TERRAFORM_VERSION=1.6.6
RUN apt-get update && apt-get install -y unzip curl && \
    curl -o terraform.zip https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform.zip && mv terraform /usr/local/bin/ && rm terraform.zip

WORKDIR /app

# Copy necessary files
COPY terraform-provider-jsonendpoint ./terraform-provider-jsonendpoint
COPY mockserver ./mockserver
COPY terraform-config ./terraform-config
COPY docker-entrypoint.sh .
COPY terraform.rc /root/.terraformrc

# Build mock server and provider
RUN cd mockserver && go mod init && go build -o mockserver
RUN cd terraform-provider-jsonendpoint && go mod tidy && go build -o terraform-provider-jsonendpoint

# Set up the plugin directory
RUN mkdir -p /usr/local/share/terraform/plugins && \
    mv terraform-provider-jsonendpoint /usr/local/share/terraform/plugins/

# Set up entrypoint
ENTRYPOINT ["/app/docker-entrypoint.sh"]
