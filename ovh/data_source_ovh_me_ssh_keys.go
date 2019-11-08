package ovh

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMeSshKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMeSshKeysRead,
		Schema: map[string]*schema.Schema{
			"names": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func dataSourceMeSshKeysRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	names := make([]string, 0)
	err := config.OVHClient.Get("/me/sshKey", &names)

	if err != nil {
		return fmt.Errorf("Error calling /me/sshKey:\n\t %q", err)
	}

	d.Set("names", names)

	log.Printf("[DEBUG] Read SSH Keys names %s", names)
	return nil
}