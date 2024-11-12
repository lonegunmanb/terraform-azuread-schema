package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azuread-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzureadGroupSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzureadGroupSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
