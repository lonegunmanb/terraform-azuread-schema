package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadServicePrincipal = `{
  "block": {
    "attributes": {
      "account_enabled": {
        "computed": true,
        "description": "Whether or not the service principal account is enabled",
        "description_kind": "plain",
        "type": "bool"
      },
      "alternative_names": {
        "computed": true,
        "description": "A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "app_role_assignment_required": {
        "computed": true,
        "description": "Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application",
        "description_kind": "plain",
        "type": "bool"
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
      "application_tenant_id": {
        "computed": true,
        "description": "The tenant ID where the associated application is registered",
        "description_kind": "plain",
        "type": "string"
      },
      "client_id": {
        "computed": true,
        "description": "The client ID of the application associated with this service principal",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the service principal provided for internal end-users",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the application associated with this service principal",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "feature_tags": {
        "computed": true,
        "description": "Block of features configured for this service principal using tags",
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
      "features": {
        "computed": true,
        "deprecated": true,
        "description": "Block of features configured for this service principal using tags",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_single_sign_on_app": "bool",
              "enterprise_application": "bool",
              "gallery_application": "bool",
              "visible_to_users": "bool"
            }
          ]
        ]
      },
      "homepage_url": {
        "computed": true,
        "description": "Home page or landing page of the application",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "login_url": {
        "computed": true,
        "description": "The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps",
        "description_kind": "plain",
        "type": "string"
      },
      "logout_url": {
        "computed": true,
        "description": "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
        "description_kind": "plain",
        "type": "string"
      },
      "notes": {
        "computed": true,
        "description": "Free text field to capture information about the service principal, typically used for operational purposes",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_email_addresses": {
        "computed": true,
        "description": "List of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
      "oauth2_permission_scopes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
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
        ]
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the service principal",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_single_sign_on_mode": {
        "computed": true,
        "description": "The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps",
        "description_kind": "plain",
        "type": "string"
      },
      "redirect_uris": {
        "computed": true,
        "description": "The URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "saml_metadata_url": {
        "computed": true,
        "description": "The URL where the service exposes SAML metadata for federation",
        "description_kind": "plain",
        "type": "string"
      },
      "saml_single_sign_on": {
        "computed": true,
        "description": "Settings related to SAML single sign-on",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "relay_state": "string"
            }
          ]
        ]
      },
      "service_principal_names": {
        "computed": true,
        "description": "A list of identifier URI(s), copied over from the associated application",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sign_in_audience": {
        "computed": true,
        "description": "The Microsoft account types that are supported for the associated application",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A set of tags to apply to the service principal",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description": "Identifies whether the service principal represents an application or a managed identity",
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

func AzureadServicePrincipalSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadServicePrincipal), &result)
	return &result
}
