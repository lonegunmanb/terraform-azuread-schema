package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadAccessPackageResourcePackageAssociation = `{
  "block": {
    "attributes": {
      "access_package_id": {
        "description": "The ID of access package this resource association is configured to",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "access_type": {
        "description": "The role of access type to the specified resource, valid values are ` + "`" + `Member` + "`" + ` and ` + "`" + `Owner` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "catalog_resource_association_id": {
        "description": "The ID of the access package catalog association",
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

func AzureadAccessPackageResourcePackageAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadAccessPackageResourcePackageAssociation), &result)
	return &result
}
