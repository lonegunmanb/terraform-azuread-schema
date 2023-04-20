package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationTemplate = `{
  "block": {
    "attributes": {
      "categories": {
        "computed": true,
        "description": "List of categories for this templated application",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "computed": true,
        "description": "The display name for the application template",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "homepage_url": {
        "computed": true,
        "description": "Home page URL of the templated application",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logo_url": {
        "computed": true,
        "description": "URL to retrieve the logo for this templated application",
        "description_kind": "plain",
        "type": "string"
      },
      "publisher": {
        "computed": true,
        "description": "Name of the publisher for this templated application",
        "description_kind": "plain",
        "type": "string"
      },
      "supported_provisioning_types": {
        "computed": true,
        "description": "The provisioning modes supported by this templated application",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supported_single_sign_on_modes": {
        "computed": true,
        "description": "The single sign on modes supported by this templated application",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "template_id": {
        "computed": true,
        "description": "The application template's ID",
        "description_kind": "plain",
        "optional": true,
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

func AzureadApplicationTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationTemplate), &result)
	return &result
}
