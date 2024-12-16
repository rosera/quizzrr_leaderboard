## PROJECT VARIABLES

# Default value passed in
variable "gcp_project_id" {
  type        = string
  description = "The GCP project ID to create resources in."
}

# Default value passed in
variable "gcp_region" {
  type        = string
  description = "Region to create resources in."
}

# Default value passed in
variable "gcp_zone" {
  type        = string
  description = "Zone to create resources in."
}


## GOOGLEAPIS VARIABLES

# Default value passed in
variable "api_services" {
  type = list(string)
  default = [
    "compute.googleapis.com",
  ]
}

# Default value passed in
variable "api_service_deny" {
  type        = bool 
  description = "Deny disable dependent services"
  default     = false
}

# Default value passed in
variable "api_create_duration" {
  type        = string
  description = "Duration delay to apply post API enablement"
  default     = "60s"
}

## VPC/SUBNETWORK VARIABLES

# Default value passed in
variable "vpc_network" {
  type        = string
  description = "Name of the VPC network to create."
  default     = "dev_network" 
}

# Default value passed in
variable "vpc_network_description" {
  type        = string
  description = "Custom VPC network."
  default     = "Lab network" 
}

# Default value passed in
variable "vpc_subnet" {
  type        = string
  description = "Name of the VPC subnetwork to create."
  default     = "dev_subnet" 
}

# Default value passed in
variable "vpc_subnet_cidr" {
  type        = string
  description = "VPC subnetwork to cidr."
  default     = "10.1.0.0/16" 
}

# Default value passed in
variable "vpc_private_google_access" {
  type        = bool 
  description = "VPC Private Google Access."
  default     = true 
}

# Default value passed in
variable "vpc_flow_logs" {
  type        = bool 
  description = "VPC Flow Logs."
  default     = false 
}

## GCE VARIABLES

# Default value passed in
variable "gce_image_project" {
  type        = string 
  description = "The project hosting the compute image."
  default     = "debian-cloud" 
}

# Default value passed in
variable "gce_image_revision" {
  type        = string 
  description = "The compute image revision."
  default     = "debian-12" 
}

# Custom properties with defaults 
variable "gce_instance_name" {
  type        = string
  description = "GCE instance name."
  default     = "lab-vm"
}

# Custom properties with defaults 
variable "gce_instance_machine_type" {
  type        = string
  description = "Machine type to use for GCE"
  default     = "e2-micro"
}

# Custom properties with defaults 
variable "gce_instance_tags" {
  type        = list(string)
  description = "GCE virtual machine tags"
  default     = ["lab-vm", "dev"]
}

# Custom properties with defaults 
## The default Metadata setting 
variable "gce_instance_metadata" {
  type        = map(string)
  description = "GCE Metadata object"
  default     = { "foo" = "bar" }
}

# Custom properties with defaults 
variable "gce_instance_scopes" {
  type        = list(string)
  description = "GCE service account scope"
  default     = ["cloud-platform"]
}

