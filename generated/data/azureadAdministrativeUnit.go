package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadAdministrativeUnit = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description for the administrative unit",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name for the administrative unit",
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
      "members": {
        "computed": true,
        "description": "A list of object IDs of members who are be present in this administrative unit.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the administrative unit",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "visibility": {
        "computed": true,
        "description": "Whether the administrative unit and its members are hidden or publicly viewable in the directory",
        "description_kind": "plain",
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
  "version": 0
}`

func AzureadAdministrativeUnitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadAdministrativeUnit), &result)
	return &result
}
