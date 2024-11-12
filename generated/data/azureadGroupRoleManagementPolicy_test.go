package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azuread-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzureadGroupRoleManagementPolicySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzureadGroupRoleManagementPolicySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
