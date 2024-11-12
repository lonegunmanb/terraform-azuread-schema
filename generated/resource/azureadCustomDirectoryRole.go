package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadCustomDirectoryRole = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The description of the custom directory role",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the custom directory role",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description": "Indicates whether the role is enabled for assignment",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the directory role",
        "description_kind": "plain",
        "type": "string"
      },
      "template_id": {
        "computed": true,
        "description": "Custom template identifier that is typically used if one needs an identifier to be the same across different directories.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "description": "The version of the role definition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "permissions": {
        "block": {
          "attributes": {
            "allowed_resource_actions": {
              "description": "Set of tasks that can be performed on a resource",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "List of permissions that are included in the custom directory role",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
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
            },
            "update": {
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

func AzureadCustomDirectoryRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadCustomDirectoryRole), &result)
	return &result
}
