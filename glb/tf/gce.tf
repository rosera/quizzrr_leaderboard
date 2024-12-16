data "google_compute_image" "debian" {
  project = var.gce_image_project
  family  = var.gce_image_revision
}

# Create Google Compute Engine instance
resource "google_compute_instance" "free_tier_vm" {
  name         = var.gce_instance_name
  machine_type = var.gce_instance_machine_type # Free tier eligible machine type
  zone         = var.gcp_zone # Choose a zone within the selected region

  boot_disk {
    initialize_params {
      # https://cloud.google.com/compute/docs/images/os-details
      # image = var.gce_public_image
      image = data.google_compute_image.debian.self_link
    }
  }

  network_interface {
    # network = "default"
    network = google_compute_network.dev_network.id
    subnetwork = google_compute_subnetwork.dev_subnet.id 
    ## access_config {
    ##   # Allow external access to the VM
    ##   # nat_ip = "EXTERNAL"
    ## }
  }

  tags = var.gce_instance_tags 

  metadata = {
    startup-script = <<EOF
      #!/bin/bash
      # STARTUP-START
      apt-get update -y
      # Update package lists and install required packages
      apt-get install -y curl jq git

      # Download Go
      curl -LO https://go.dev/dl/go1.23.4.linux-amd64.tar.gz 

      # Install Go
      rm -rf /usr/local/go && tar -C /tmp -xzf go1.23.4.linux-amd64.tar.gz
      # cp /tmp/go/bin/go* /usr/local/bin

      # Env Var
      export USER="api-dev"

      # Create a user
      useradd $USER -m -p Password01 -s /bin/bash -c 'Developer Account'
    EOF
  }

  service_account {
    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    # email  = google_service_account.default.email
    scopes = var.gce_instance_scopes 
  }
}
