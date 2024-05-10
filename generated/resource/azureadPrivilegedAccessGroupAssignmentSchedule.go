package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadPrivilegedAccessGroupAssignmentSchedule = `{
  "block": {
    "attributes": {
      "assignment_type": {
        "description": "The ID of the assignment to the group",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "duration": {
        "description": "The duration of the assignment, formatted as an ISO8601 duration string (e.g. P3D for 3 days)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "expiration_date": {
        "computed": true,
        "description": "The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_id": {
        "description": "The ID of the Group representing the scope of the assignment",
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
      "justification": {
        "description": "The justification for the assignment",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "permanent_assignment": {
        "computed": true,
        "description": "Is the assignment permanent",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "principal_id": {
        "description": "The ID of the Principal assigned to the schedule",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_date": {
        "computed": true,
        "description": "The date that this assignment starts, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the schedule",
        "description_kind": "plain",
        "type": "string"
      },
      "ticket_number": {
        "description": "The ticket number authorising the assignment",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ticket_system": {
        "description": "The ticket system authorising the assignment",
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

func AzureadPrivilegedAccessGroupAssignmentScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadPrivilegedAccessGroupAssignmentSchedule), &result)
	return &result
}
