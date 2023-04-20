package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azuread-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzureadGroupsSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzureadGroupsSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
