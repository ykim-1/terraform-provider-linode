package main

import (
	"github.com/linode/terraform-provider-linode/linode"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
}

func teardown() {
}

var testAccProviders map[string]*schema.Provider

func init() {
	testAccProviders = map[string]*schema.Provider{
		"linode": linode.Provider(),
	}
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"linode_instance": resourceLinodeInstance(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}

func resourceLinodeInstance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
		Create: resourceLinodeInstanceCreate,
		Read:   resourceLinodeInstanceRead,
		Update: resourceLinodeInstanceUpdate,
		Delete: resourceLinodeInstanceDelete,
	}
}

func resourceLinodeInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	// Implement the Create method for the Linode instance resource here
	return nil
}

func resourceLinodeInstanceRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the Read method for the Linode instance resource here
	return nil
}

func resourceLinodeInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	// Implement the Update method for the Linode instance resource here
	return nil
}

func resourceLinodeInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the Delete method for the Linode instance resource here
	return nil
}
