package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadSynchronizationJobProvisionOnDemand = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_principal_id": {
        "description": "The object ID of the service principal for which this synchronization job should be provisioned",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "synchronization_job_id": {
        "description": "The identifier for the synchronization jop.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "triggers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "parameter": {
        "block": {
          "attributes": {
            "rule_id": {
              "description": "The identifier of the synchronization rule to be applied. This rule ID is defined in the schema for a given synchronization job or template.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "subject": {
              "block": {
                "attributes": {
                  "object_id": {
                    "description": "The identifier of an object to which a synchronization job is to be applied. Can be one of the following: (1) An onPremisesDistinguishedName for synchronization from Active Directory to Azure AD. (2) The user ID for synchronization from Azure AD to a third-party. (3) The Worker ID of the Workday worker for synchronization from Workday to either Active Directory or Azure AD.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "object_type_name": {
                    "description": "The type of the object to which a synchronization job is to be applied. Can be one of the following: ` + "`" + `user` + "`" + ` for synchronizing between Active Directory and Azure AD, ` + "`" + `User` + "`" + ` for synchronizing a user between Azure AD and a third-party application, ` + "`" + `Worker` + "`" + ` for synchronization a user between Workday and either Active Directory or Azure AD, ` + "`" + `Group` + "`" + ` for synchronizing a group between Azure AD and a third-party application.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The identifiers of one or more objects to which a synchronizationJob is to be applied.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Represents the objects that will be provisioned and the synchronization rules executed. The resource is primarily used for on-demand provisioning.",
          "description_kind": "plain"
        },
        "min_items": 1,
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

func AzureadSynchronizationJobProvisionOnDemandSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadSynchronizationJobProvisionOnDemand), &result)
	return &result
}
