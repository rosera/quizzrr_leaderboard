# Create Firewall rulebase
# https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall
resource "google_compute_firewall" "network-allow-rule" {
  name          = var.fw_gce_name 
  network       = google_compute_network.dev_network.id
  source_ranges = var.fw_gce_source_ranges
  # Enable INGRESS
  direction     = var.fw_gce_direction
  project       = var.gcp_project_id

  allow {
    protocol    = var.fw_gce_protocol
    ports       = var.fw_gce_ports
  }

  # Apply the rule to target_tags
  source_tags = gce_instance_tags 
  target_tags = gce_instance_tags 

}
