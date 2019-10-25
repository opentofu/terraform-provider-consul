package consul

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataACLPolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceACLPolicyConfigNotFound,
				ExpectError: regexp.MustCompile("Could not find policy 'not-found'"),
			},
			{
				Config: testAccDataSourceACLPolicyConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceValue("data.consul_acl_policy.test", "name", "test"),
					testAccCheckDataSourceValue("data.consul_acl_policy.test", "description", "foo"),
					testAccCheckDataSourceValue("data.consul_acl_policy.test", "rules", "node_prefix \"\" { policy = \"read\" }"),
					testAccCheckDataSourceValue("data.consul_acl_policy.test", "datacenters.#", "1"),
					testAccCheckDataSourceValue("data.consul_acl_policy.test", "datacenters.0", "dc1"),
				),
			},
		},
	})
}

const testAccDataSourceACLPolicyConfigNotFound = `
data "consul_acl_policy" "test" {
	name = "not-found"
}
`

const testAccDataSourceACLPolicyConfigBasic = `
resource "consul_acl_policy" "test" {
	name = "test"
	description = "foo"
	rules = "node_prefix \"\" { policy = \"read\" }"
	datacenters = [ "dc1" ]
}

data "consul_acl_policy" "test" {
	name = consul_acl_policy.test.name
}
`