package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azuread-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzureadApplicationTemplateSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzureadApplicationTemplateSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
