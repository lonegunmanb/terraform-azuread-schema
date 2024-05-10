package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadGroupRoleManagementPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Description of the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "group_id": {
        "description": "ID of the group to which this policy is assigned",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_id": {
        "description": "The ID of the role of this policy to the group",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func AzureadGroupRoleManagementPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadGroupRoleManagementPolicy), &result)
	return &result
}
