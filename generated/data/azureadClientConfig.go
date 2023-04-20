package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadClientConfig = `{
  "block": {
    "attributes": {
      "client_id": {
        "computed": true,
        "description": "The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the authenticated principal",
        "description_kind": "plain",
        "type": "string"
      },
      "tenant_id": {
        "computed": true,
        "description": "The tenant ID of the authenticated principal",
        "description_kind": "plain",
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

func AzureadClientConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadClientConfig), &result)
	return &result
}
