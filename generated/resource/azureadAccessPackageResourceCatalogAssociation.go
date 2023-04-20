package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadAccessPackageResourceCatalogAssociation = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "description": "The unique ID of the access package catalog",
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
      "resource_origin_id": {
        "description": "The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_origin_system": {
        "description": "The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup",
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

func AzureadAccessPackageResourceCatalogAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadAccessPackageResourceCatalogAssociation), &result)
	return &result
}
