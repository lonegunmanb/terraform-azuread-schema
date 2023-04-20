package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadAccessPackageCatalog = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the access package catalog",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the access package catalog",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "externally_visible": {
        "computed": true,
        "description": "Whether the access packages in this catalog can be requested by users outside the tenant",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_id": {
        "computed": true,
        "description": "The ID of this access package catalog",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "published": {
        "computed": true,
        "description": "Whether the access packages in this catalog are available for management",
        "description_kind": "plain",
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

func AzureadAccessPackageCatalogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadAccessPackageCatalog), &result)
	return &result
}
