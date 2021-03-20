terraform {
  required_providers {
    segfault = {
      version = "0.0.1"
      source = "nspain.dev/test/segfault"
    }
  }
}

provider "segfault" {}

data "segfault" "test" {}

