resource "google_compute_network" "dev_network" {
  name                    = var.vpc_network
  description             = var.vpc_network_description
  auto_create_subnetworks = false
  project                 = var.gcp_project_id

  # Give API time to be provisioned 
  depends_on = [time_sleep.wait_api_delay]
}

resource "time_sleep" "wait_vpc_delay" {
  create_duration = "60s"
  depends_on      = [google_compute_network.dev_network]
}

# Reference:
# https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_subnetwork
resource "google_compute_subnetwork" "dev_subnet" {
  name                     = var.vpc_subnet
  region                   = var.gcp_region
  ip_cidr_range            = var.vpc_subnet_cidr
  network                  = google_compute_network.dev_network.id
  project                  = var.gcp_project_id
  private_ip_google_access = var.vpc_private_google_access

  # Give network time to be created
  depends_on = [time_sleep.wait_vpc_delay]
}
