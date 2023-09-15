package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadGroup = `{
  "block": {
    "attributes": {
      "assignable_to_role": {
        "computed": true,
        "description": "Indicates whether this group can be assigned to an Azure Active Directory role",
        "description_kind": "plain",
        "type": "bool"
      },
      "auto_subscribe_new_members": {
        "computed": true,
        "description": "Indicates whether new members added to the group will be auto-subscribed to receive email notifications.",
        "description_kind": "plain",
        "type": "bool"
      },
      "behaviors": {
        "computed": true,
        "description": "The group behaviors for a Microsoft 365 group",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "The optional description of the group",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name for the group",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dynamic_membership": {
        "computed": true,
        "description": "An optional block to configure dynamic membership for the group. Cannot be used with ` + "`" + `members` + "`" + `",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "rule": "string"
            }
          ]
        ]
      },
      "external_senders_allowed": {
        "computed": true,
        "description": "Indicates whether people external to the organization can send messages to the group.",
        "description_kind": "plain",
        "type": "bool"
      },
      "hide_from_address_lists": {
        "computed": true,
        "description": "Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups.",
        "description_kind": "plain",
        "type": "bool"
      },
      "hide_from_outlook_clients": {
        "computed": true,
        "description": "Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "Whether the group is mail-enabled",
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
        "description": "The object IDs of the group members",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the group",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The object IDs of the group owners",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "preferred_language": {
        "computed": true,
        "description": "The preferred language for a Microsoft 365 group, in ISO 639-1 notation",
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_options": {
        "computed": true,
        "description": "The group provisioning options for a Microsoft 365 group",
        "description_kind": "plain",
        "type": [
          "list",
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
        "computed": true,
        "description": "Whether the group is a security group",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "theme": {
        "computed": true,
        "description": "The colour theme for a Microsoft 365 group",
        "description_kind": "plain",
        "type": "string"
      },
      "types": {
        "computed": true,
        "description": "A list of group types configured for the group. The only supported type is ` + "`" + `Unified` + "`" + `, which specifies a Microsoft 365 group",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "visibility": {
        "computed": true,
        "description": "Specifies the group join policy and group content visibility",
        "description_kind": "plain",
        "type": "string"
      },
      "writeback_enabled": {
        "computed": true,
        "description": "Whether this group is synced from Azure AD to the on-premises directory when Azure AD Connect is used",
        "description_kind": "plain",
        "type": "bool"
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

func AzureadGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadGroup), &result)
	return &result
}
