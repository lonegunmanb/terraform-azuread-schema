package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azuread-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzureadApplicationRegistrationSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzureadApplicationRegistrationSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
