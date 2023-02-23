provider "google" {
  project     = "dev-playground-378601"
  region      = "us-east1"
}

terraform {
  backend "gcs" {
    bucket  = "terraform-dguillermo-dev"
    prefix  = "gqlgo/state"
  }
}
