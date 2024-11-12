package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationPreAuthorized = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "The resource ID of the application to which this pre-authorized application should be added",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authorized_client_id": {
        "description": "The client ID of the pre-authorized application",
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
      "permission_ids": {
        "description": "The IDs of the permission scopes required by the pre-authorized application",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
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

func AzureadApplicationPreAuthorizedSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationPreAuthorized), &result)
	return &result
}
