package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationOptionalClaims = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "The resource ID of the application to which these optional claims belong",
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
      "access_token": {
        "block": {
          "attributes": {
            "additional_properties": {
              "description": "List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "essential": {
              "description": "Whether the claim specified by the client is necessary to ensure a smooth authorization experience",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description": "The name of the optional claim",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source": {
              "description": "The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "id_token": {
        "block": {
          "attributes": {
            "additional_properties": {
              "description": "List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "essential": {
              "description": "Whether the claim specified by the client is necessary to ensure a smooth authorization experience",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description": "The name of the optional claim",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source": {
              "description": "The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "saml2_token": {
        "block": {
          "attributes": {
            "additional_properties": {
              "description": "List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "essential": {
              "description": "Whether the claim specified by the client is necessary to ensure a smooth authorization experience",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description": "The name of the optional claim",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source": {
              "description": "The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func AzureadApplicationOptionalClaimsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationOptionalClaims), &result)
	return &result
}
