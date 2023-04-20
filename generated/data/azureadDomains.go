package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadDomains = `{
  "block": {
    "attributes": {
      "admin_managed": {
        "description": "Set to ` + "`" + `true` + "`" + ` to only return domains whose DNS is managed by Microsoft 365",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "domains": {
        "computed": true,
        "description": "A list of tenant domains",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "admin_managed": "bool",
              "authentication_type": "string",
              "default": "bool",
              "domain_name": "string",
              "initial": "bool",
              "root": "bool",
              "supported_services": [
                "list",
                "string"
              ],
              "verified": "bool"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_unverified": {
        "description": "Set to ` + "`" + `true` + "`" + ` if unverified Azure AD domains should be included",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "only_default": {
        "description": "Set to ` + "`" + `true` + "`" + ` to only return the default domain",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "only_initial": {
        "description": "Set to ` + "`" + `true` + "`" + ` to only return the initial domain, which is your primary Azure Active Directory tenant domain",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "only_root": {
        "description": "Set to ` + "`" + `true` + "`" + ` to only return verified root domains. Excludes subdomains and unverified domains",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "supports_services": {
        "description": "A list of supported services that must be supported by a domain",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzureadDomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadDomains), &result)
	return &result
}
