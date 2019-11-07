package ovh

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ovh/go-ovh/ovh"
)

func resourceSshKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceSshKeyCreate,
		Read:   resourceSshKeyRead,
		Update: resourceSshKeyUpdate,
		Delete: resourceSshKeyDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"key_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of this public Ssh key",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ASCII encoded public Ssh key",
			},
			"default": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True when this public Ssh key is used for rescue mode and reinstallations",
			},
		},
	}
}

func resourceSshKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	keyName := d.Get("key_name").(string)
	key := d.Get("key").(string)
	params := &SshKeyCreateOpts{
		KeyName: keyName,
		Key:     key,
	}

	log.Printf("[DEBUG] Will create Ssh key: %s", params)

	err := config.OVHClient.Post("/me/sshKey", params, nil)
	if err != nil {
		return fmt.Errorf("Error creating SSH Key with params %s:\n\t %q", params, err)
	}

	return resourceSshKeyRead(d, meta)
}

func resourceSshKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	resultKey := &SshKeyResponse{}

	keyName := d.Get("key_name").(string)
	err := config.OVHClient.Get(
		fmt.Sprintf("/me/sshKey/%s", keyName),
		resultKey,
	)
	if err != nil {
		return nil, fmt.Errorf("Unable to find SSH key named %s:\n\t %q", keyName, err)
	}

	d.Set("key_name", resultKey.KeyName)
	d.Set("key", resultKey.Key)
	d.Set("default", resultKey.Default)

	return nil
}

func resourceSshKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	keyName := d.Get("key_name").(string)
	params := &SshKeyUpdateOpts{
		Default: d.Get("default").(string),
	}
	err := config.OVHClient.Put(
		fmt.Sprintf("/me/sshKey/%s", keyName),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Unable to update SSH key named %s:\n\t %q", keyName, err)
	}

	log.Printf("[DEBUG] Updated SSH Key %s", keyName)
	return resourceSshKeyRead(d, meta)
}

func resourceSshKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	keyName := d.Get("key_name").(string)
	err := config.OVHClient.Delete(
		fmt.Sprintf("/me/sshKey/%s", keyName),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Unable to delete SSH key named %s:\n\t %q", keyName, err)
	}

	log.Printf("[DEBUG] Deleted SSH Key %s", keyName)
	return nil
}
