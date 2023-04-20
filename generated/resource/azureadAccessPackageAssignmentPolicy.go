package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azureadAccessPackageAssignmentPolicy = `{
  "block": {
    "attributes": {
      "access_package_id": {
        "description": "The ID of the access package that will contain the policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of the policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "duration_in_days": {
        "description": "How many days this assignment is valid for",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "expiration_date": {
        "description": "The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "extension_enabled": {
        "description": "When enabled, users will be able to request extension of their access to this package before their access expires",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "approval_settings": {
        "block": {
          "attributes": {
            "approval_required": {
              "description": "Whether an approval is required",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "approval_required_for_extension": {
              "description": "Whether an approval is required to grant extension. Same approval settings used to approve initial access will apply",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "requestor_justification_required": {
              "description": "Whether requestor are required to provide a justification to request an access package. Justification is visible to other approvers and the requestor",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "approval_stage": {
              "block": {
                "attributes": {
                  "alternative_approval_enabled": {
                    "description": "If no action taken, forward to alternate approvers?",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "approval_timeout_in_days": {
                    "description": "Decision must be made in how many days? If a request is not approved within this time period after it is made, it will be automatically rejected",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "approver_justification_required": {
                    "description": "Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_alternative_approval_in_days": {
                    "description": "Forward to alternate approver(s) after how many days?",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "alternative_approver": {
                    "block": {
                      "attributes": {
                        "backup": {
                          "description": "For a user in an approval stage, this property indicates whether the user is a backup fallback approver",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "object_id": {
                          "description": "The object ID of the subject",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "subject_type": {
                          "description": "Type of users",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "primary_approver": {
                    "block": {
                      "attributes": {
                        "backup": {
                          "description": "For a user in an approval stage, this property indicates whether the user is a backup fallback approver",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "object_id": {
                          "description": "The object ID of the subject",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "subject_type": {
                          "description": "Type of users",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The process to obtain an approval",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Settings of whether approvals are required and how they are obtained",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "assignment_review_settings": {
        "block": {
          "attributes": {
            "access_recommendation_enabled": {
              "description": "Whether to show Show reviewer decision helpers. If enabled, system recommendations based on users' access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "access_review_timeout_behavior": {
              "description": "What actions the system takes if reviewers don't respond in time",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "approver_justification_required": {
              "description": "Whether a reviewer need provide a justification for their decision. Justification is visible to other reviewers and the requestor",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "duration_in_days": {
              "description": "How many days each occurrence of the access review series will run",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enabled": {
              "description": "Whether to enable assignment review",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "review_frequency": {
              "description": "This will determine how often the access review campaign runs",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "review_type": {
              "description": "Self review or specific reviewers",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "starting_on": {
              "description": "This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "reviewer": {
              "block": {
                "attributes": {
                  "backup": {
                    "description": "For a user in an approval stage, this property indicates whether the user is a backup fallback approver",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "object_id": {
                    "description": "The object ID of the subject",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject_type": {
                    "description": "Type of users",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The settings of whether assignment review is needed and how it's conducted",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "question": {
        "block": {
          "attributes": {
            "required": {
              "description": "Whether this question is required",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sequence": {
              "description": "The sequence number of this question",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "choice": {
              "block": {
                "attributes": {
                  "actual_value": {
                    "description": "The actual value of this choice",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "display_value": {
                    "block": {
                      "attributes": {
                        "default_text": {
                          "description": "The default text of this question",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "localized_text": {
                          "block": {
                            "attributes": {
                              "content": {
                                "description": "The localized content of this question",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "language_code": {
                                "description": "The language code of this question content",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "The localized text of this question",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The display text of this choice",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Configuration of a choice to the question",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "text": {
              "block": {
                "attributes": {
                  "default_text": {
                    "description": "The default text of this question",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "localized_text": {
                    "block": {
                      "attributes": {
                        "content": {
                          "description": "The localized content of this question",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "language_code": {
                          "description": "The language code of this question content",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The localized text of this question",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The content of this question",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "One or more questions to the requestor",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "requestor_settings": {
        "block": {
          "attributes": {
            "requests_accepted": {
              "description": "Whether to accept requests now, when disabled, no new requests can be made using this policy",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "scope_type": {
              "description": "Specify the scopes of the requestors",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "requestor": {
              "block": {
                "attributes": {
                  "backup": {
                    "description": "For a user in an approval stage, this property indicates whether the user is a backup fallback approver",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "object_id": {
                    "description": "The object ID of the subject",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject_type": {
                    "description": "Type of users",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "This block configures the users who can request access",
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

func AzureadAccessPackageAssignmentPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azureadAccessPackageAssignmentPolicy), &result)
	return &result
}
