package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadServicePrincipal = `{
  "block": {
    "attributes": {
      "account_enabled": {
        "description": "Whether or not the service principal account is enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "alternative_names": {
        "description": "A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "app_role_assignment_required": {
        "description": "Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application",
        "description_kind": "plain",
        "optional": true,
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
      "application_id": {
        "description": "The application ID (client ID) of the application for which to create a service principal",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "application_tenant_id": {
        "computed": true,
        "description": "The tenant ID where the associated application is registered",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the service principal provided for internal end-users",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the application associated with this service principal",
        "description_kind": "plain",
        "type": "string"
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
        "description": "The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logout_url": {
        "computed": true,
        "description": "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
        "description_kind": "plain",
        "type": "string"
      },
      "notes": {
        "description": "Free text field to capture information about the service principal, typically used for operational purposes",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notification_email_addresses": {
        "description": "List of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
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
        "type": "string"
      },
      "owners": {
        "description": "A list of object IDs of principals that will be granted ownership of the service principal",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "preferred_single_sign_on_mode": {
        "description": "The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description": "Identifies whether the service principal represents an application or a managed identity",
        "description_kind": "plain",
        "type": "string"
      },
      "use_existing": {
        "description": "When true, the resource will return an existing service principal instead of failing with an error",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "feature_tags": {
        "block": {
          "attributes": {
            "custom_single_sign_on": {
              "description": "Whether this service principal represents a custom SAML application",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enterprise": {
              "description": "Whether this service principal represents an Enterprise Application",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "gallery": {
              "description": "Whether this service principal represents a gallery application",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hide": {
              "description": "Whether this app is invisible to users in My Apps and Office 365 Launcher",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Block of features to configure for this service principal using tags",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "features": {
        "block": {
          "attributes": {
            "custom_single_sign_on_app": {
              "description": "Whether this service principal represents a custom SAML application",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enterprise_application": {
              "description": "Whether this service principal represents an Enterprise Application",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "gallery_application": {
              "description": "Whether this service principal represents a gallery application",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "visible_to_users": {
              "description": "Whether this app is visible to users in My Apps and Office 365 Launcher",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "deprecated": true,
          "description": "Block of features to configure for this service principal using tags",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "saml_single_sign_on": {
        "block": {
          "attributes": {
            "relay_state": {
              "description": "The relative URI the service provider would redirect to after completion of the single sign-on flow",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Settings related to SAML single sign-on",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AzureadServicePrincipalSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadServicePrincipal), &result)
	return &result
}
