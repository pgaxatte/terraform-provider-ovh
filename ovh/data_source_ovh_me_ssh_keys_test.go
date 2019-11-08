package ovh

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccMeSshKeysDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCredentials(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMeSshKeysDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ovh_me_ssh_keys.keys", "names"),
				),
			},
		},
	})
}

const testAccMeSshKeysDatasourceConfig = `
data "ovh_me_ssh_keys" "keys" {}
`
