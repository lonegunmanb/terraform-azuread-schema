package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadServicePrincipalClaimsMappingPolicyAssignment = `{
  "block": {
    "attributes": {
      "claims_mapping_policy_id": {
        "description": "ID of the claims mapping policy to assign",
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
        "description": "Object ID of the service principal for which to assign the policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzureadServicePrincipalClaimsMappingPolicyAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadServicePrincipalClaimsMappingPolicyAssignment), &result)
	return &result
}
