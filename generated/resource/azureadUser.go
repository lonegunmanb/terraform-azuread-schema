package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadUser = `{
  "block": {
    "attributes": {
      "about_me": {
        "computed": true,
        "description": "A freeform field for the user to describe themselves",
        "description_kind": "plain",
        "type": "string"
      },
      "account_enabled": {
        "description": "Whether or not the account should be enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "age_group": {
        "description": "The age group of the user",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "business_phones": {
        "computed": true,
        "description": "The telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "city": {
        "description": "The city in which the user is located",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "company_name": {
        "description": "The company name which the user is associated. This property can be useful for describing the company that an external user comes from",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "consent_provided_for_minor": {
        "description": "Whether consent has been obtained for minors",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cost_center": {
        "description": "The cost center associated with the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "country": {
        "description": "The country/region in which the user is located, e.g. ` + "`" + `US` + "`" + ` or ` + "`" + `UK` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_type": {
        "computed": true,
        "description": "Indicates whether the user account was created as a regular school or work account (` + "`" + `null` + "`" + `), an external account (` + "`" + `Invitation` + "`" + `), a local account for an Azure Active Directory B2C tenant (` + "`" + `LocalAccount` + "`" + `) or self-service sign-up using email verification (` + "`" + `EmailVerified` + "`" + `)",
        "description_kind": "plain",
        "type": "string"
      },
      "department": {
        "description": "The name for the department in which the user works",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_password_expiration": {
        "description": "Whether the users password is exempt from expiring",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_strong_password": {
        "description": "Whether the user is allowed weaker passwords than the default policy to be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "The name to display in the address book for the user",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "division": {
        "description": "The name of the division in which the user works.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "employee_id": {
        "description": "The employee identifier assigned to the user by the organisation",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "employee_type": {
        "description": "Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "external_user_state": {
        "computed": true,
        "description": "For an external user invited to the tenant, this property represents the invited user's invitation status",
        "description_kind": "plain",
        "type": "string"
      },
      "fax_number": {
        "description": "The fax number of the user",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_password_change": {
        "description": "Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "given_name": {
        "description": "The given name (first name) of the user",
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
        "description": "The user’s job title",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mail": {
        "computed": true,
        "description": "The SMTP address for the user. Cannot be unset.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mail_nickname": {
        "computed": true,
        "description": "The mail alias for the user. Defaults to the user name part of the user principal name (UPN)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manager_id": {
        "description": "The object ID of the user's manager",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mobile_phone": {
        "description": "The primary cellular telephone number for the user",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the user",
        "description_kind": "plain",
        "type": "string"
      },
      "office_location": {
        "description": "The office location in the user's place of business",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's ` + "`" + `user_principal_name` + "`" + ` property when creating a new user account",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Additional email addresses for the user",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "password": {
        "computed": true,
        "description": "The password for the user. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters. This property is required when creating a new user",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "postal_code": {
        "description": "The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_language": {
        "description": "The user's preferred language, in ISO 639-1 notation",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Whether or not the Outlook global address list should include this user",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "state": {
        "description": "The state or province in the user's address",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "street_address": {
        "description": "The street address of the user's place of business",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "surname": {
        "description": "The user's surname (family name or last name)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "usage_location": {
        "description": "The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: ` + "`" + `NO` + "`" + `, ` + "`" + `JP` + "`" + `, and ` + "`" + `GB` + "`" + `. Cannot be reset to null once set",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_principal_name": {
        "description": "The user principal name (UPN) of the user",
        "description_kind": "plain",
        "required": true,
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

func AzureadUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadUser), &result)
	return &result
}
