package main

import (
	"fmt"
	"github.com/linode/terraform-provider-linode/linode"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestLinodeInstance(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"linode": linode.Provider()},
		CheckDestroy: testCheckLinodeInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testLinodeInstanceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testCheckLinodeInstanceIsRunning("linode_instance.foobar"),
					resource.TestCheckResourceAttr("linode_instance.foobar", "label", "test-instance"),
					resource.TestCheckResourceAttr("linode_instance.foobar", "region", "us-east"),
				),
			},
		},
	})
}

func testLinodeInstanceConfig() string {
	return `
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
`
}

func testCheckLinodeInstanceIsRunning(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		instanceResource, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Instance resource %s not found", resourceName)
		}

		instanceStatus := instanceResource.Primary.Attributes["status"]
		if instanceStatus != "running" {
			return fmt.Errorf("Instance %s is not in running state. Current state: %s", resourceName, instanceStatus)
		}

		return nil
	}
}

func testCheckLinodeInstanceDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "linode_instance" {
			continue
		}

		instanceStatus := rs.Primary.Attributes["status"]
		if instanceStatus != "absent" {
			return fmt.Errorf("Instance %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
