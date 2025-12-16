package template_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/binaryYuki/error-pages/internal/template"
)

func TestProps_Values(t *testing.T) {
	t.Parallel()

	assert.Equal(t, template.Props{
		Code:               1,
		Message:            "b",
		Description:        "c",
		RequestID:          "d",
		Host:               "e",
		ShowRequestDetails: false,
		L10nDisabled:       true,
	}.Values(), map[string]any{
		"code":          uint16(1),
		"message":       "b",
		"description":   "c",
		"request_id":    "d",
		"host":          "e",
		"show_details":  false,
		"l10n_disabled": true,
	})
}
