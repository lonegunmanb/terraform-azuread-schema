package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplication = `{
  "block": {
    "attributes": {
      "api": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "known_client_applications": [
                "list",
                "string"
              ],
              "mapped_claims_enabled": "bool",
              "oauth2_permission_scopes": [
                "list",
                [
                  "object",
                  {
                    "admin_consent_description": "string",
                    "admin_consent_display_name": "string",
                    "enabled": "bool",
                    "id": "string",
                    "type": "string",
                    "user_consent_description": "string",
                    "user_consent_display_name": "string",
                    "value": "string"
                  }
                ]
              ],
              "requested_access_token_version": "number"
            }
          ]
        ]
      },
      "app_role_ids": {
        "computed": true,
        "description": "Mapping of app role names to UUIDs",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "app_roles": {
        "computed": true,
        "description": "List of app roles published by the application",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allowed_member_types": [
                "list",
                "string"
              ],
              "description": "string",
              "display_name": "string",
              "enabled": "bool",
              "id": "string",
              "value": "string"
            }
          ]
        ]
      },
      "client_id": {
        "computed": true,
        "description": "The Client ID (also called Application ID)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the application as shown to end users",
        "description_kind": "plain",
        "type": "string"
      },
      "device_only_auth_enabled": {
        "computed": true,
        "description": "Specifies whether this application supports device authentication without a user.",
        "description_kind": "plain",
        "type": "bool"
      },
      "disabled_by_microsoft": {
        "computed": true,
        "description": "Whether Microsoft has disabled the registered application",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name for the application",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fallback_public_client_enabled": {
        "computed": true,
        "description": "The fallback application type as public client, such as an installed application running on a mobile device",
        "description_kind": "plain",
        "type": "bool"
      },
      "feature_tags": {
        "computed": true,
        "description": "Block of features configured for this application using tags",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_single_sign_on": "bool",
              "enterprise": "bool",
              "gallery": "bool",
              "hide": "bool"
            }
          ]
        ]
      },
      "group_membership_claims": {
        "computed": true,
        "description": "The ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identifier_uri": {
        "computed": true,
        "description": "One of the application's identifier URIs",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identifier_uris": {
        "computed": true,
        "description": "A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "logo_url": {
        "computed": true,
        "description": "CDN URL to the application's logo",
        "description_kind": "plain",
        "type": "string"
      },
      "marketing_url": {
        "computed": true,
        "description": "URL of the application's marketing page",
        "description_kind": "plain",
        "type": "string"
      },
      "notes": {
        "computed": true,
        "description": "User-specified notes relevant for the management of the application",
        "description_kind": "plain",
        "type": "string"
      },
      "oauth2_permission_scope_ids": {
        "computed": true,
        "description": "Mapping of OAuth2.0 permission scope names to UUIDs",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "oauth2_post_response_required": {
        "computed": true,
        "description": "Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests.",
        "description_kind": "plain",
        "type": "bool"
      },
      "object_id": {
        "computed": true,
        "description": "The application's object ID",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "optional_claims": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_token": [
                "list",
                [
                  "object",
                  {
                    "additional_properties": [
                      "list",
                      "string"
                    ],
                    "essential": "bool",
                    "name": "string",
                    "source": "string"
                  }
                ]
              ],
              "id_token": [
                "list",
                [
                  "object",
                  {
                    "additional_properties": [
                      "list",
                      "string"
                    ],
                    "essential": "bool",
                    "name": "string",
                    "source": "string"
                  }
                ]
              ],
              "saml2_token": [
                "list",
                [
                  "object",
                  {
                    "additional_properties": [
                      "list",
                      "string"
                    ],
                    "essential": "bool",
                    "name": "string",
                    "source": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "owners": {
        "computed": true,
        "description": "A list of object IDs of principals that are assigned ownership of the application",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "privacy_statement_url": {
        "computed": true,
        "description": "URL of the application's privacy statement",
        "description_kind": "plain",
        "type": "string"
      },
      "public_client": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "redirect_uris": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "publisher_domain": {
        "computed": true,
        "description": "The verified publisher domain for the application",
        "description_kind": "plain",
        "type": "string"
      },
      "required_resource_access": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resource_access": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "type": "string"
                  }
                ]
              ],
              "resource_app_id": "string"
            }
          ]
        ]
      },
      "service_management_reference": {
        "computed": true,
        "description": "References application or service contact information from a Service or Asset Management database",
        "description_kind": "plain",
        "type": "string"
      },
      "sign_in_audience": {
        "computed": true,
        "description": "The Microsoft account types that are supported for the current application",
        "description_kind": "plain",
        "type": "string"
      },
      "single_page_application": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "redirect_uris": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "support_url": {
        "computed": true,
        "description": "URL of the application's support page",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A set of tags applied to the application",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "terms_of_service_url": {
        "computed": true,
        "description": "URL of the application's terms of service statement",
        "description_kind": "plain",
        "type": "string"
      },
      "web": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "homepage_url": "string",
              "implicit_grant": [
                "list",
                [
                  "object",
                  {
                    "access_token_issuance_enabled": "bool",
                    "id_token_issuance_enabled": "bool"
                  }
                ]
              ],
              "logout_url": "string",
              "redirect_uris": [
                "list",
                "string"
              ]
            }
          ]
        ]
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

func AzureadApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplication), &result)
	return &result
}
