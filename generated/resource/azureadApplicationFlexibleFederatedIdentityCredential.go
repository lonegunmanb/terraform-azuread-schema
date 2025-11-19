package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationFlexibleFederatedIdentityCredential = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "The resource ID of the application for which this flexible federated identity credential should be created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "audience": {
        "description": "The audience that can appear in the external token. This specifies what should be accepted in the ` + "`" + `aud` + "`" + ` claim of incoming tokens.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "claims_matching_expression": {
        "description": "The expression to match for claims.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "credential_id": {
        "computed": true,
        "description": "A UUID used to uniquely identify this flexible federated identity credential",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description for the flexible federated identity credential",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "A unique display name for the flexible federated identity credential",
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

func AzureadApplicationFlexibleFederatedIdentityCredentialSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationFlexibleFederatedIdentityCredential), &result)
	return &result
}
