package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationFederatedIdentityCredential = `{
  "block": {
    "attributes": {
      "application_object_id": {
        "description": "The object ID of the application for which this federated identity credential should be created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "audiences": {
        "description": "List of audiences that can appear in the external token. This specifies what should be accepted in the ` + "`" + `aud` + "`" + ` claim of incoming tokens.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "credential_id": {
        "computed": true,
        "description": "A UUID used to uniquely identify this federated identity credential",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description for the federated identity credential",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "A unique display name for the federated identity credential",
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
      "issuer": {
        "description": "The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subject": {
        "description": "The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.",
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

func AzureadApplicationFederatedIdentityCredentialSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationFederatedIdentityCredential), &result)
	return &result
}
