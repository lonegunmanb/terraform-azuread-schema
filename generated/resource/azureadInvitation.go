package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadInvitation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redeem_url": {
        "computed": true,
        "description": "The URL the user can use to redeem their invitation",
        "description_kind": "plain",
        "type": "string"
      },
      "redirect_url": {
        "description": "The URL that the user should be redirected to once the invitation is redeemed",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_display_name": {
        "description": "The display name of the user being invited",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_email_address": {
        "description": "The email address of the user being invited",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_id": {
        "computed": true,
        "description": "Object ID of the invited user",
        "description_kind": "plain",
        "type": "string"
      },
      "user_type": {
        "description": "The user type of the user being invited",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "message": {
        "block": {
          "attributes": {
            "additional_recipients": {
              "description": "Email addresses of additional recipients the invitation message should be sent to",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "body": {
              "description": "Customized message body you want to send if you don't want to send the default message",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "language": {
              "description": "The language you want to send the default message in",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Customize the message sent to the invited user",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func AzureadInvitationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadInvitation), &result)
	return &result
}
