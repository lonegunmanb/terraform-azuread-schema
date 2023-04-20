package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadUserFlowAttribute = `{
  "block": {
    "attributes": {
      "attribute_type": {
        "computed": true,
        "description": "The type of the user flow attribute",
        "description_kind": "plain",
        "type": "string"
      },
      "data_type": {
        "description": "The data type of the user flow attribute",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of the user flow attribute that is shown to the user at the time of sign-up",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the user flow attribute.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AzureadUserFlowAttributeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadUserFlowAttribute), &result)
	return &result
}
