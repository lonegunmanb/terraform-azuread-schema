package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadAppRoleAssignment = `{
  "block": {
    "attributes": {
      "app_role_id": {
        "description": "The ID of the app role to be assigned",
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
      "principal_display_name": {
        "computed": true,
        "description": "The display name of the principal to which the app role is assigned",
        "description_kind": "plain",
        "type": "string"
      },
      "principal_object_id": {
        "description": "The object ID of the user, group or service principal to be assigned this app role",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_type": {
        "computed": true,
        "description": "The object type of the principal to which the app role is assigned",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_display_name": {
        "computed": true,
        "description": "The display name of the application representing the resource",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_object_id": {
        "description": "The object ID of the service principal representing the resource",
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
  "version": 1
}`

func AzureadAppRoleAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadAppRoleAssignment), &result)
	return &result
}
