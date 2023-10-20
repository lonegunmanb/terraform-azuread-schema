package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadServicePrincipals = `{
  "block": {
    "attributes": {
      "application_ids": {
        "computed": true,
        "deprecated": true,
        "description": "The application IDs (client IDs) of the applications associated with the service principals",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "client_ids": {
        "computed": true,
        "description": "The client IDs of the applications associated with the service principals",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_names": {
        "computed": true,
        "description": "The display names of the applications associated with the service principals",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignore_missing": {
        "description": "Ignore missing service principals and return the service principals that were found. The data source will still fail if no service principals are found",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "object_ids": {
        "computed": true,
        "description": "The object IDs of the service principals",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "return_all": {
        "description": "Fetch all service principals with no filter and return all that were found. The data source will still fail if no service principals are found.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "service_principals": {
        "computed": true,
        "description": "A list of service_principals",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_enabled": "bool",
              "app_role_assignment_required": "bool",
              "application_id": "string",
              "application_tenant_id": "string",
              "client_id": "string",
              "display_name": "string",
              "object_id": "string",
              "preferred_single_sign_on_mode": "string",
              "saml_metadata_url": "string",
              "service_principal_names": [
                "list",
                "string"
              ],
              "sign_in_audience": "string",
              "tags": [
                "list",
                "string"
              ],
              "type": "string"
            }
          ]
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

func AzureadServicePrincipalsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadServicePrincipals), &result)
	return &result
}
