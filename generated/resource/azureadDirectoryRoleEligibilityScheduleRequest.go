package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadDirectoryRoleEligibilityScheduleRequest = `{
  "block": {
    "attributes": {
      "directory_scope_id": {
        "description": "Identifier of the directory object representing the scope of the role eligibility schedule request",
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
      "justification": {
        "description": "Justification for why the role is assigned",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_id": {
        "description": "The object ID of the member principal",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_definition_id": {
        "description": "The object ID of the directory role for this role eligibility schedule request",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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

func AzureadDirectoryRoleEligibilityScheduleRequestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadDirectoryRoleEligibilityScheduleRequest), &result)
	return &result
}
