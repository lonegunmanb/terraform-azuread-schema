package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadServicePrincipalDelegatedPermissionGrant = `{
  "block": {
    "attributes": {
      "claim_values": {
        "description": "A set of claim values for delegated permission scopes which should be included in access tokens for the resource",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_service_principal_object_id": {
        "description": "The object ID of the service principal representing the resource to be accessed",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_principal_object_id": {
        "description": "The object ID of the service principal for which this delegated permission grant should be created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_object_id": {
        "description": "The object ID of the user on behalf of whom the service principal is authorized to access the resource",
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
  "version": 1
}`

func AzureadServicePrincipalDelegatedPermissionGrantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadServicePrincipalDelegatedPermissionGrant), &result)
	return &result
}
