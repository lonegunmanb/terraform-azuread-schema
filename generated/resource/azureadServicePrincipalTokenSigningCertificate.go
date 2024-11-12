package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadServicePrincipalTokenSigningCertificate = `{
  "block": {
    "attributes": {
      "display_name": {
        "computed": true,
        "description": "A friendly name for the certificate",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_date": {
        "computed": true,
        "description": "The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Default is 3 years from current date.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_id": {
        "computed": true,
        "description": "A UUID used to uniquely identify the verify certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_principal_id": {
        "description": "The ID of the service principal for which this certificate should be created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_date": {
        "computed": true,
        "description": "The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).",
        "description_kind": "plain",
        "type": "string"
      },
      "thumbprint": {
        "computed": true,
        "description": "The thumbprint of the certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "value": {
        "computed": true,
        "description": "The certificate data, which is PEM encoded but does not include the header/footer",
        "description_kind": "plain",
        "sensitive": true,
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

func AzureadServicePrincipalTokenSigningCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadServicePrincipalTokenSigningCertificate), &result)
	return &result
}
