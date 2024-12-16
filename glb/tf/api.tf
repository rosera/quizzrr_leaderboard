# Enable the Googleapi service

resource "google_project_service" "tlf" {
  ## https://developer.hashicorp.com/terraform/language/functions/toset
  ## Removes duplicates
  for_each = toset(var.api_services)

  project = var.gcp_project_id
  service = each.key

  disable_dependent_services = var.api_service_deny
}

# Introduce a JIT delay for API enablement
## Add a Delay before creating a Workbench instance
resource "time_sleep" "wait_api_delay" {
  depends_on = [google_project_service.tlf]
  ## create_duration = "60s"
  create_duration = var.api_create_duration
}
