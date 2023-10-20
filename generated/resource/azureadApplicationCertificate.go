package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplicationCertificate = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The resource ID of the application for which this certificate should be created",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "application_object_id": {
        "computed": true,
        "deprecated": true,
        "description": "The object ID of the application for which this certificate should be created",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encoding": {
        "description": "Specifies the encoding used for the supplied certificate data",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_date": {
        "computed": true,
        "description": "The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_date_relative": {
        "description": "A relative duration for which the certificate is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `",
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
        "description": "A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_date": {
        "computed": true,
        "description": "The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date and time are use",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of key/certificate",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "value": {
        "description": "The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the ` + "`" + `encoding` + "`" + ` argument",
        "description_kind": "plain",
        "required": true,
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

func AzureadApplicationCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplicationCertificate), &result)
	return &result
}
