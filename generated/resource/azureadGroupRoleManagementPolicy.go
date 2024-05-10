package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadGroupRoleManagementPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Description of the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "group_id": {
        "description": "ID of the group to which this policy is assigned",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_id": {
        "description": "The ID of the role of this policy to the group",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "activation_rules": {
        "block": {
          "attributes": {
            "maximum_duration": {
              "computed": true,
              "description": "The time after which the an activation can be valid for",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "require_approval": {
              "computed": true,
              "description": "Whether an approval is required for activation",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "require_justification": {
              "computed": true,
              "description": "Whether a justification is required during activation",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "require_multifactor_authentication": {
              "computed": true,
              "description": "Whether multi-factor authentication is required during activation",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "require_ticket_info": {
              "computed": true,
              "description": "Whether ticket information is required during activation",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "required_conditional_access_authentication_context": {
              "computed": true,
              "description": "Whether a conditional access context is required during activation",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "approval_stage": {
              "block": {
                "block_types": {
                  "primary_approver": {
                    "block": {
                      "attributes": {
                        "object_id": {
                          "description": "The ID of the object to act as an approver",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "The type of object acting as an approver",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The IDs of the users or groups who can approve the activation",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description": "The approval stages for the activation",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The activation rules of the policy",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "active_assignment_rules": {
        "block": {
          "attributes": {
            "expiration_required": {
              "computed": true,
              "description": "Must the assignment have an expiry date",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "expire_after": {
              "computed": true,
              "description": "The duration after which assignments expire",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "require_justification": {
              "computed": true,
              "description": "Whether a justification is required to make an assignment",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "require_multifactor_authentication": {
              "computed": true,
              "description": "Whether multi-factor authentication is required to make an assignment",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "require_ticket_info": {
              "computed": true,
              "description": "Whether ticket information is required to make an assignment",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "The rules for active assignment of the policy",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "eligible_assignment_rules": {
        "block": {
          "attributes": {
            "expiration_required": {
              "computed": true,
              "description": "Must the assignment have an expiry date",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "expire_after": {
              "computed": true,
              "description": "The duration after which assignments expire",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The rules for eligible assignment of the policy",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "notification_rules": {
        "block": {
          "block_types": {
            "active_assignments": {
              "block": {
                "block_types": {
                  "admin_notifications": {
                    "block": {
                      "attributes": {
                        "additional_recipients": {
                          "computed": true,
                          "description": "The additional recipients to notify",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "default_recipients": {
                          "description": "Whether the default recipients are notified",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "notification_level": {
                          "description": "What level of notifications are sent",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Admin notification settings",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "approver_notifications": {
                    "block": {
                      "attributes": {
                        "additional_recipients": {
                          "computed": true,
                          "description": "The additional recipients to notify",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "default_recipients": {
                          "description": "Whether the default recipients are notified",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "notification_level": {
                          "description": "What level of notifications are sent",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Approver notification settings",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "assignee_notifications": {
                    "block": {
                      "attributes": {
                        "additional_recipients": {
                          "computed": true,
                          "description": "The additional recipients to notify",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "default_recipients": {
                          "description": "Whether the default recipients are notified",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "notification_level": {
                          "description": "What level of notifications are sent",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Assignee notification settings",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Notifications about active assignments",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "eligible_activations": {
              "block": {
                "block_types": {
                  "admin_notifications": {
                    "block": {
                      "attributes": {
                        "additional_recipients": {
                          "computed": true,
                          "description": "The additional recipients to notify",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "default_recipients": {
                          "description": "Whether the default recipients are notified",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "notification_level": {
                          "description": "What level of notifications are sent",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Admin notification settings",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "approver_notifications": {
                    "block": {
                      "attributes": {
                        "additional_recipients": {
                          "computed": true,
                          "description": "The additional recipients to notify",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "default_recipients": {
                          "description": "Whether the default recipients are notified",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "notification_level": {
                          "description": "What level of notifications are sent",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Approver notification settings",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "assignee_notifications": {
                    "block": {
                      "attributes": {
                        "additional_recipients": {
                          "computed": true,
                          "description": "The additional recipients to notify",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "default_recipients": {
                          "description": "Whether the default recipients are notified",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "notification_level": {
                          "description": "What level of notifications are sent",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Assignee notification settings",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Notifications about activations of eligible assignments",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "eligible_assignments": {
              "block": {
                "block_types": {
                  "admin_notifications": {
                    "block": {
                      "attributes": {
                        "additional_recipients": {
                          "computed": true,
                          "description": "The additional recipients to notify",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "default_recipients": {
                          "description": "Whether the default recipients are notified",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "notification_level": {
                          "description": "What level of notifications are sent",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Admin notification settings",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "approver_notifications": {
                    "block": {
                      "attributes": {
                        "additional_recipients": {
                          "computed": true,
                          "description": "The additional recipients to notify",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "default_recipients": {
                          "description": "Whether the default recipients are notified",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "notification_level": {
                          "description": "What level of notifications are sent",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Approver notification settings",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "assignee_notifications": {
                    "block": {
                      "attributes": {
                        "additional_recipients": {
                          "computed": true,
                          "description": "The additional recipients to notify",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "default_recipients": {
                          "description": "Whether the default recipients are notified",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "notification_level": {
                          "description": "What level of notifications are sent",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Assignee notification settings",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Notifications about eligible assignments",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The notification rules of the policy",
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

func AzureadGroupRoleManagementPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadGroupRoleManagementPolicy), &result)
	return &result
}
