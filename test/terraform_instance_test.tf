terraform {
  required_providers {
    linode = {
      source  = "linode/linode"
      version = "2.4.0"
    }
  }
}

provider "linode" {
  token = "cab2575d2ee51f52d7b4fb3b40c0dbc43c59544b0a614d30cc0f717808df8c2b"
}

data "terraform_remote_state" "example" {
  backend = "mock"
}

resource "linode_instance" "foobar" {
  label           = "test-instance"
  group           = "tf_test"
  type            = "g6-nanode-1"
  image           = "linode/ubuntu18.04"
  region          = "us-east"
  root_pass       = "myr00tp@ssw0rd!!!"
  swap_size       = 256
  authorized_keys = ["~/.ssh/id_rsa.pub"]
}

data "linode_instance" "foobar" {
  instance_id = linode_instance.foobar.id
}

data "terraform_remote_state" "mock" {
  backend = "mock"
}
