package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azuread-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzureadAccessPackageCatalogRoleAssignmentSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzureadAccessPackageCatalogRoleAssignmentSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
