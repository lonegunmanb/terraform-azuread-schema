package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationRegistration = `{
  "block": {
    "attributes": {
      "client_id": {
        "computed": true,
        "description": "The Client ID (also called Application ID)",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the application as shown to end users",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled_by_microsoft": {
        "computed": true,
        "description": "If the application has been disabled by Microsoft, this shows the status or reason",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The display name for the application",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_membership_claims": {
        "description": "Configures the ` + "`" + `groups` + "`" + ` claim that the app expects issued in a user or OAuth access token",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "homepage_url": {
        "description": "URL of the home page for the application",
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
      "implicit_access_token_issuance_enabled": {
        "description": "Whether this application can request an access token using OAuth implicit flow",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "implicit_id_token_issuance_enabled": {
        "description": "Whether this application can request an ID token using OAuth implicit flow",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "logout_url": {
        "description": "URL of the logout page for the application, where the session is cleared for single sign-out",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "marketing_url": {
        "description": "URL of the marketing page for the application",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notes": {
        "description": "User-specified notes relevant for the management of the application",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the application within the tenant",
        "description_kind": "plain",
        "type": "string"
      },
      "privacy_statement_url": {
        "description": "URL of the privacy statement for the application",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publisher_domain": {
        "computed": true,
        "description": "The verified publisher domain for the application",
        "description_kind": "plain",
        "type": "string"
      },
      "requested_access_token_version": {
        "description": "The access token version expected by this resource",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "service_management_reference": {
        "description": "References application or contact information from a service or asset management database",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sign_in_audience": {
        "description": "The Microsoft account types that are supported for the current application",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "support_url": {
        "description": "URL of the support page for the application",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "terms_of_service_url": {
        "description": "URL of the terms of service statement for the application",
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

func AzureadApplicationRegistrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationRegistration), &result)
	return &result
}
