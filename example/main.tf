terraform {
  required_providers {
    pacman = {
      source = "danskespil/pacman"
    }
  }
}

provider "pacman" {
  uri      = "https://pacman.address"
  username = "username"
  password = "password"
}

resource "pacman_asset" "this" {
  asset_name  = "test-1"
  domain_name = "domain.address"
  asset_type  = "TEST"
  ip_address  = "1.2.3.4"
}
