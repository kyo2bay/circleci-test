provider "google" {
  project = var.project.name
  region  = "asia-northeast1"
}

resource "google_pubsub_topic" "slack_notify" {
  name    = "slack-notify"
  project = var.project.name
}
