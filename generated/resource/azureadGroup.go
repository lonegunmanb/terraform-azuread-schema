package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadGroup = `{
  "block": {
    "attributes": {
      "administrative_unit_ids": {
        "description": "The administrative unit IDs in which the group should be. If empty, the group will be created at the tenant level.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "assignable_to_role": {
        "description": "Indicates whether this group can be assigned to an Azure Active Directory role. This property can only be ` + "`" + `true` + "`" + ` for security-enabled groups.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_subscribe_new_members": {
        "computed": true,
        "description": "Indicates whether new members added to the group will be auto-subscribed to receive email notifications.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "behaviors": {
        "description": "The group behaviours for a Microsoft 365 group",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "description": {
        "description": "The description for the group",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name for the group",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "external_senders_allowed": {
        "computed": true,
        "description": "Indicates whether people external to the organization can send messages to the group.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hide_from_address_lists": {
        "computed": true,
        "description": "Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hide_from_outlook_clients": {
        "computed": true,
        "description": "Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mail": {
        "computed": true,
        "description": "The SMTP address for the group",
        "description_kind": "plain",
        "type": "string"
      },
      "mail_enabled": {
        "description": "Whether the group is a mail enabled, with a shared group mailbox. At least one of ` + "`" + `mail_enabled` + "`" + ` or ` + "`" + `security_enabled` + "`" + ` must be specified. A group can be mail enabled _and_ security enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "mail_nickname": {
        "computed": true,
        "description": "The mail alias for the group, unique in the organisation",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "members": {
        "computed": true,
        "description": "A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the group",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_domain_name": {
        "computed": true,
        "description": "The on-premises FQDN, also called dnsDomainName, synchronized from the on-premises directory when Azure AD Connect is used",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_group_type": {
        "computed": true,
        "description": "Indicates the target on-premise group type the group will be written back as",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "onpremises_netbios_name": {
        "computed": true,
        "description": "The on-premises NetBIOS name, synchronized from the on-premises directory when Azure AD Connect is used",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_sam_account_name": {
        "computed": true,
        "description": "The on-premises SAM account name, synchronized from the on-premises directory when Azure AD Connect is used",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_security_identifier": {
        "computed": true,
        "description": "The on-premises security identifier (SID), synchronized from the on-premises directory when Azure AD Connect is used",
        "description_kind": "plain",
        "type": "string"
      },
      "onpremises_sync_enabled": {
        "computed": true,
        "description": "Whether this group is synchronized from an on-premises directory (true), no longer synchronized (false), or has never been synchronized (null)",
        "description_kind": "plain",
        "type": "bool"
      },
      "owners": {
        "computed": true,
        "description": "A set of owners who own this group. Supported object types are Users or Service Principals",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "preferred_language": {
        "computed": true,
        "description": "The preferred language for a Microsoft 365 group, in ISO 639-1 notation",
        "description_kind": "plain",
        "type": "string"
      },
      "prevent_duplicate_names": {
        "description": "If ` + "`" + `true` + "`" + `, will return an error if an existing group is found with the same name",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "provisioning_options": {
        "description": "The group provisioning options for a Microsoft 365 group",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "proxy_addresses": {
        "computed": true,
        "description": "Email addresses for the group that direct to the same group mailbox",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "security_enabled": {
        "description": "Whether the group is a security group for controlling access to in-app resources. At least one of ` + "`" + `security_enabled` + "`" + ` or ` + "`" + `mail_enabled` + "`" + ` must be specified. A group can be security enabled _and_ mail enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "theme": {
        "description": "The colour theme for a Microsoft 365 group",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "types": {
        "description": "A set of group types to configure for the group. ` + "`" + `Unified` + "`" + ` specifies a Microsoft 365 group. Required when ` + "`" + `mail_enabled` + "`" + ` is true",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "visibility": {
        "computed": true,
        "description": "Specifies the group join policy and group content visibility",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "writeback_enabled": {
        "description": "Whether this group should be synced from Azure AD to the on-premises directory when Azure AD Connect is used",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "dynamic_membership": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "rule": {
              "description": "Rule to determine members for a dynamic group. Required when ` + "`" + `group_types` + "`" + ` contains 'DynamicMembership'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "An optional block to configure dynamic membership for the group. Cannot be used with ` + "`" + `members` + "`" + `",
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

func AzureadGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadGroup), &result)
	return &result
}
