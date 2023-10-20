package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationPermissionScope = `{
  "block": {
    "attributes": {
      "admin_consent_description": {
        "description": "Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "admin_consent_display_name": {
        "description": "Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "application_id": {
        "description": "The resource ID of the application to which this permission scope should be applied",
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
      "scope_id": {
        "description": "The unique identifier of the permission scope",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_consent_description": {
        "description": "Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_consent_display_name": {
        "description": "Display name for the delegated permission that appears in the end user consent experience",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "value": {
        "description": "The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth access tokens",
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

func AzureadApplicationPermissionScopeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationPermissionScope), &result)
	return &result
}
