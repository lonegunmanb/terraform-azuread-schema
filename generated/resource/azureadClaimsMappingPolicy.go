package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadClaimsMappingPolicy = `{
  "block": {
    "attributes": {
      "definition": {
        "description": "A string collection containing a JSON string that defines the rules and settings for this policy",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "description": "Display name for this policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzureadClaimsMappingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadClaimsMappingPolicy), &result)
	return &result
}
