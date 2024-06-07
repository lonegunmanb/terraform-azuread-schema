package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadUsers = `{
  "block": {
    "attributes": {
      "employee_ids": {
        "computed": true,
        "description": "The employee identifier assigned to the user by the organisation",
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
        "description": "Ignore missing users and return users that were found. The data source will still fail if no users are found",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "mail_nicknames": {
        "computed": true,
        "description": "The email aliases of the users",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "mails": {
        "computed": true,
        "description": "The SMTP address of the users",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "object_ids": {
        "computed": true,
        "description": "The object IDs of the users",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "return_all": {
        "description": "Fetch all users with no filter and return all that were found. The data source will still fail if no users are found.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_principal_names": {
        "computed": true,
        "description": "The user principal names (UPNs) of the users",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "users": {
        "computed": true,
        "description": "A list of users",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_enabled": "bool",
              "display_name": "string",
              "employee_id": "string",
              "mail": "string",
              "mail_nickname": "string",
              "object_id": "string",
              "onpremises_immutable_id": "string",
              "onpremises_sam_account_name": "string",
              "onpremises_user_principal_name": "string",
              "usage_location": "string",
              "user_principal_name": "string"
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

func AzureadUsersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadUsers), &result)
	return &result
}
