---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "consul_acl_role_policy_attachment Resource - terraform-provider-consul"
subcategory: ""
description: |-
  The consul_acl_role_policy_attachment resource links a Consul ACL role and an ACL policy. The link is implemented through an update to the Consul ACL role.
  ~> NOTE: This resource is only useful to attach policies to an ACL role that has been created outside the current Terraform configuration. If the ACL role you need to attach a policy to has been created in the current Terraform configuration and will only be used in it, you should use the policies attribute of consul_acl_role.
---

# consul_acl_role_policy_attachment (Resource)

The `consul_acl_role_policy_attachment` resource links a Consul ACL role and an ACL policy. The link is implemented through an update to the Consul ACL role.

~> **NOTE:** This resource is only useful to attach policies to an ACL role that has been created outside the current Terraform configuration. If the ACL role you need to attach a policy to has been created in the current Terraform configuration and will only be used in it, you should use the `policies` attribute of [`consul_acl_role`](/docs/providers/consul/r/acl_role.html).

## Example Usage

```terraform
data "consul_acl_role" "my_role" {
  name = "my_role"
}

resource "consul_acl_policy" "read_policy" {
  name        = "read-policy"
  rules       = "node \"\" { policy = \"read\" }"
  datacenters = ["dc1"]
}

resource "consul_acl_role_policy_attachment" "my_role_read_policy" {
  role_id = data.consul_acl_role.test.id
  policy  = consul_acl_policy.read_policy.name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy` (String) The policy name.
- `role_id` (String) The id of the role.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import consul_acl_role_policy_attachment.my_role_read_policy 624d94ca-bc5c-f960-4e83-0a609cf588be:policy_name
```