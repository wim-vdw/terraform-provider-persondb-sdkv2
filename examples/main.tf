provider "myprovider" {}

resource "myprovider_dashboard" "example" {
  name        = "example-dashboard"
  description = "This is an example dashboard."
}
