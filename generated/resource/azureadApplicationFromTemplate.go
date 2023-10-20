package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationFromTemplate = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The resource ID for this application",
        "description_kind": "plain",
        "type": "string"
      },
      "application_object_id": {
        "computed": true,
        "description": "The object ID for this application",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The display name for the application",
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
      "service_principal_id": {
        "computed": true,
        "description": "The resource ID for this service principal",
        "description_kind": "plain",
        "type": "string"
      },
      "service_principal_object_id": {
        "computed": true,
        "description": "The object ID for this service principal",
        "description_kind": "plain",
        "type": "string"
      },
      "template_id": {
        "description": "The UUID of the template to instantiate for this application",
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

func AzureadApplicationFromTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationFromTemplate), &result)
	return &result
}
