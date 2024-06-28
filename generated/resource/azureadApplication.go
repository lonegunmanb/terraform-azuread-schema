package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadApplication = `{
  "block": {
    "attributes": {
      "app_role_ids": {
        "computed": true,
        "description": "Mapping of app role names to UUIDs",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "application_id": {
        "computed": true,
        "deprecated": true,
        "description": "The Application ID (also called Client ID)",
        "description_kind": "plain",
        "type": "string"
      },
      "client_id": {
        "computed": true,
        "description": "The Client ID (also called Application ID)",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the application as shown to end users",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "device_only_auth_enabled": {
        "description": "Specifies whether this application supports device authentication without a user.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disabled_by_microsoft": {
        "computed": true,
        "description": "Whether Microsoft has disabled the registered application",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The display name for the application",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fallback_public_client_enabled": {
        "description": "Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "group_membership_claims": {
        "description": "Configures the ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identifier_uris": {
        "description": "The user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "logo_image": {
        "description": "Base64 encoded logo image in gif, png or jpeg format",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logo_url": {
        "computed": true,
        "description": "CDN URL to the application's logo",
        "description_kind": "plain",
        "type": "string"
      },
      "marketing_url": {
        "description": "URL of the application's marketing page",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notes": {
        "description": "User-specified notes relevant for the management of the application",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "object_id": {
        "computed": true,
        "description": "The application's object ID",
        "description_kind": "plain",
        "type": "string"
      },
      "owners": {
        "description": "A list of object IDs of principals that will be granted ownership of the application",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "prevent_duplicate_names": {
        "description": "If ` + "`" + `true` + "`" + `, will return an error if an existing application is found with the same name",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "privacy_statement_url": {
        "description": "URL of the application's privacy statement",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publisher_domain": {
        "computed": true,
        "description": "The verified publisher domain for the application",
        "description_kind": "plain",
        "type": "string"
      },
      "service_management_reference": {
        "description": "References application or service contact information from a Service or Asset Management database",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sign_in_audience": {
        "description": "The Microsoft account types that are supported for the current application",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "support_url": {
        "description": "URL of the application's support page",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A set of tags to apply to the application",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "template_id": {
        "computed": true,
        "description": "Unique ID of the application template from which this application is created",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "terms_of_service_url": {
        "description": "URL of the application's terms of service statement",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "api": {
        "block": {
          "attributes": {
            "known_client_applications": {
              "description": "Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "mapped_claims_enabled": {
              "description": "Allows an application to use claims mapping without specifying a custom signing key",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "requested_access_token_version": {
              "description": "The access token version expected by this resource",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "oauth2_permission_scope": {
              "block": {
                "attributes": {
                  "admin_consent_description": {
                    "description": "Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "admin_consent_display_name": {
                    "description": "Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "description": "Determines if the permission scope is enabled",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "id": {
                    "description": "The unique identifier of the delegated permission",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_consent_description": {
                    "description": "Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_consent_display_name": {
                    "description": "Display name for the delegated permission that appears in the end user consent experience",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "One or more ` + "`" + `oauth2_permission_scope` + "`" + ` blocks to describe delegated permissions exposed by the web API represented by this application",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "app_role": {
        "block": {
          "attributes": {
            "allowed_member_types": {
              "description": "Specifies whether this app role definition can be assigned to users and groups by setting to ` + "`" + `User` + "`" + `, or to other applications (that are accessing this application in a standalone scenario) by setting to ` + "`" + `Application` + "`" + `, or to both",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "description": {
              "description": "Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "display_name": {
              "description": "Display name for the app role that appears during app role assignment and in consent experiences",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "enabled": {
              "description": "Determines if the app role is enabled",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "id": {
              "description": "The unique identifier of the app role",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "feature_tags": {
        "block": {
          "attributes": {
            "custom_single_sign_on": {
              "description": "Whether this application represents a custom SAML application for linked service principals",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enterprise": {
              "description": "Whether this application represents an Enterprise Application for linked service principals",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "gallery": {
              "description": "Whether this application represents a gallery application for linked service principals",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hide": {
              "description": "Whether this application is invisible to users in My Apps and Office 365 Launcher",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Block of features to configure for this application using tags",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "optional_claims": {
        "block": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "password": {
        "block": {
          "attributes": {
            "display_name": {
              "description": "A display name for the password",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "end_date": {
              "computed": true,
              "description": "The end date until which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `)",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_id": {
              "computed": true,
              "description": "A UUID used to uniquely identify this password credential",
              "description_kind": "plain",
              "type": "string"
            },
            "start_date": {
              "computed": true,
              "description": "The start date from which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The password for this application, which is generated by Azure Active Directory",
              "description_kind": "plain",
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "App password definition",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "public_client": {
        "block": {
          "attributes": {
            "redirect_uris": {
              "description": "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "required_resource_access": {
        "block": {
          "attributes": {
            "resource_app_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "resource_access": {
              "block": {
                "attributes": {
                  "id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "single_page_application": {
        "block": {
          "attributes": {
            "redirect_uris": {
              "description": "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
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
      },
      "web": {
        "block": {
          "attributes": {
            "homepage_url": {
              "description": "Home page or landing page of the application",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "logout_url": {
              "description": "The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redirect_uris": {
              "description": "The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "implicit_grant": {
              "block": {
                "attributes": {
                  "access_token_issuance_enabled": {
                    "description": "Whether this web application can request an access token using OAuth 2.0 implicit flow",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "id_token_issuance_enabled": {
                    "description": "Whether this web application can request an ID token using OAuth 2.0 implicit flow",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 2
}`

func AzureadApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadApplication), &result)
	return &result
}
