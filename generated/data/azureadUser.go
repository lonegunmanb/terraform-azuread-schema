package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadUser = `{
  "block": {
    "attributes": {
      "account_enabled": {
        "computed": true,
        "description": "Whether or not the account is enabled",
        "description_kind": "plain",
        "type": "bool"
      },
      "age_group": {
        "computed": true,
        "description": "The age group of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "business_phones": {
        "computed": true,
        "description": "The telephone numbers for the user",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "city": {
        "computed": true,
        "description": "The city in which the user is located",
        "description_kind": "plain",
        "type": "string"
      },
      "company_name": {
        "computed": true,
        "description": "The company name which the user is associated. This property can be useful for describing the company that an external user comes from",
        "description_kind": "plain",
        "type": "string"
      },
      "consent_provided_for_minor": {
        "computed": true,
        "description": "Whether consent has been obtained for minors",
        "description_kind": "plain",
        "type": "string"
      },
      "cost_center": {
        "computed": true,
        "description": "The cost center associated with the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "country": {
        "computed": true,
        "description": "The country/region in which the user is located, e.g. ` + "`" + `US` + "`" + ` or ` + "`" + `UK` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_type": {
        "computed": true,
        "description": "Indicates whether the user account was created as a regular school or work account (` + "`" + `null` + "`" + `), an external account (` + "`" + `Invitation` + "`" + `), a local account for an Azure Active Directory B2C tenant (` + "`" + `LocalAccount` + "`" + `) or self-service sign-up using email verification (` + "`" + `EmailVerified` + "`" + `)",
        "description_kind": "plain",
        "type": "string"
      },
      "department": {
        "computed": true,
        "description": "The name for the department in which the user works",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "division": {
        "computed": true,
        "description": "The name of the division in which the user works.",
        "description_kind": "plain",
        "type": "string"
      },
      "employee_id": {
        "computed": true,
        "description": "The employee identifier assigned to the user by the organisation",
        "description_kind": "plain",
        "type": "string"
      },
      "employee_type": {
        "computed": true,
        "description": "Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.",
        "description_kind": "plain",
        "type": "string"
      },
      "external_user_state": {
        "computed": true,
        "description": "For an external user invited to the tenant, this property represents the invited user's invitation status",
        "description_kind": "plain",
        "type": "string"
      },
      "fax_number": {
        "computed": true,
        "description": "The fax number of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "given_name": {
        "computed": true,
        "description": "The given name (first name) of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "im_addresses": {
        "computed": true,
        "description": "The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "job_title": {
        "computed": true,
        "description": "The user’s job title",
        "description_kind": "plain",
        "type": "string"
      },
      "mail": {
        "computed": true,
        "description": "The SMTP address for the user",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mail_nickname": {
        "computed": true,
        "description": "The email alias of the user",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manager_id": {
        "computed": true,
        "description": "The object ID of the user's manager",
        "description_kind": "plain",
        "type": "string"
      },
      "mobile_phone": {
        "computed": true,
        "description": "The primary cellular telephone number for the user",
        "description_kind": "plain",
        "type": "string"
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the user",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "office_location": {
        "computed": true,
        "description": "The office location in the user's place of business",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_distinguished_name": {
        "computed": true,
        "description": "The on-premise Active Directory distinguished name (DN) of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_domain_name": {
        "computed": true,
        "description": "The on-premise FQDN (i.e. dnsDomainName) of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_immutable_id": {
        "computed": true,
        "description": "The value used to associate an on-premise Active Directory user account with their Azure AD user object",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_sam_account_name": {
        "computed": true,
        "description": "The on-premise SAM account name of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_security_identifier": {
        "computed": true,
        "description": "The on-premise security identifier (SID) of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_sync_enabled": {
        "computed": true,
        "description": "Whether this user is synchronized from an on-premises directory (true), no longer synchronized (false), or has never been synchronized (null)",
        "description_kind": "plain",
        "type": "bool"
      },
      "onpremises_user_principal_name": {
        "computed": true,
        "description": "The on-premise user principal name of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "other_mails": {
        "computed": true,
        "description": "Additional email addresses for the user",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "postal_code": {
        "computed": true,
        "description": "The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_language": {
        "computed": true,
        "description": "The user's preferred language, in ISO 639-1 notation",
        "description_kind": "plain",
        "type": "string"
      },
      "proxy_addresses": {
        "computed": true,
        "description": "Email addresses for the user that direct to the same mailbox",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "show_in_address_list": {
        "computed": true,
        "description": "Whether or not the Outlook global address list should include this user",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The state or province in the user's address",
        "description_kind": "plain",
        "type": "string"
      },
      "street_address": {
        "computed": true,
        "description": "The street address of the user's place of business",
        "description_kind": "plain",
        "type": "string"
      },
      "surname": {
        "computed": true,
        "description": "The user's surname (family name or last name)",
        "description_kind": "plain",
        "type": "string"
      },
      "usage_location": {
        "computed": true,
        "description": "The usage location of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "user_principal_name": {
        "computed": true,
        "description": "The user principal name (UPN) of the user",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_type": {
        "computed": true,
        "description": "The user type in the directory. Possible values are ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `",
        "description_kind": "plain",
        "type": "string"
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

func AzureadUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadUser), &result)
	return &result
}
