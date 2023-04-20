package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadDirectoryRoles = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_ids": {
        "computed": true,
        "description": "The object IDs of the roles",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "roles": {
        "computed": true,
        "description": "A list of roles",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "display_name": "string",
              "object_id": "string",
              "template_id": "string"
            }
          ]
        ]
      },
      "template_ids": {
        "computed": true,
        "description": "The template IDs of the roles",
        "description_kind": "plain",
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

func AzureadDirectoryRolesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadDirectoryRoles), &result)
	return &result
}
