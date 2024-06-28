package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadDirectoryRoleAssignment = `{
  "block": {
    "attributes": {
      "app_scope_id": {
        "computed": true,
        "description": "Identifier of the app-specific scope when the assignment scope is app-specific",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "app_scope_object_id": {
        "computed": true,
        "deprecated": true,
        "description": "Identifier of the app-specific scope when the assignment scope is app-specific",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directory_scope_id": {
        "computed": true,
        "description": "Identifier of the directory object representing the scope of the assignment",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directory_scope_object_id": {
        "computed": true,
        "description": "Identifier of the directory object representing the scope of the assignment",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "principal_object_id": {
        "description": "The object ID of the member principal",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_id": {
        "description": "The object ID of the directory role for this assignment",
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

func AzureadDirectoryRoleAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadDirectoryRoleAssignment), &result)
	return &result
}
