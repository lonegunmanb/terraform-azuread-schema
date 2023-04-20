package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadGroups = `{
  "block": {
    "attributes": {
      "display_name_prefix": {
        "computed": true,
        "description": "Common display name prefix of the groups",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_names": {
        "computed": true,
        "description": "The display names of the groups",
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
        "description": "Ignore missing groups and return groups that were found. The data source will still fail if no groups are found",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "mail_enabled": {
        "computed": true,
        "description": "Whether the groups are mail-enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "object_ids": {
        "computed": true,
        "description": "The object IDs of the groups",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "return_all": {
        "description": "Retrieve all groups with no filter",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "security_enabled": {
        "computed": true,
        "description": "Whether the groups are security-enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func AzureadGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadGroups), &result)
	return &result
}
