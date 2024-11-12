package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadSynchronizationJob = `{
  "block": {
    "attributes": {
      "enabled": {
        "description": "Whether or not the synchronization job is enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "expiration": "string",
              "interval": "string",
              "state": "string"
            }
          ]
        ]
      },
      "service_principal_id": {
        "description": "The ID of the service principal for which this synchronization job should be created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_id": {
        "description": "Identifier of the synchronization template this job is based on.",
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
  "version": 1
}`

func AzureadSynchronizationJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadSynchronizationJob), &result)
	return &result
}
