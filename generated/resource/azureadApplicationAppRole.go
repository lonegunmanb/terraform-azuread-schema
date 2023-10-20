package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationAppRole = `{
  "block": {
    "attributes": {
      "allowed_member_types": {
        "description": "Specifies whether this app role definition can be assigned to users and groups by setting to ` + "`" + `User` + "`" + `, or to other applications (that are accessing this application in a standalone scenario) by setting to ` + "`" + `Application` + "`" + `, or to both",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "application_id": {
        "description": "The resource ID of the application to which this app role should be applied",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name for the app role that appears during app role assignment and in consent experiences",
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
      "role_id": {
        "description": "The unique identifier of the app role",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "value": {
        "description": "The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth access tokens that are authenticating an assigned service or user principal",
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

func AzureadApplicationAppRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationAppRole), &result)
	return &result
}
