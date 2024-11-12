package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadAdministrativeUnit = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The description for the administrative unit",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name for the administrative unit",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hidden_membership_enabled": {
        "description": "Whether the administrative unit and its members are hidden or publicly viewable in the directory",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "members": {
        "computed": true,
        "description": "A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the administrative unit",
        "description_kind": "plain",
        "type": "string"
      },
      "prevent_duplicate_names": {
        "description": "If ` + "`" + `true` + "`" + `, will return an error if an existing administrative unit is found with the same name",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
  "version": 1
}`

func AzureadAdministrativeUnitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadAdministrativeUnit), &result)
	return &result
}
