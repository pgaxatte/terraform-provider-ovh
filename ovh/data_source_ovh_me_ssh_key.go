package ovh

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMeSshKey() *schema.Resource {
	return &schema.Resource{
		Read: readMeSshKey,
		Schema: map[string]*schema.Schema{
			"key_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
