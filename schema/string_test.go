package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	name := "John"

	validSchema := String().LengthMin(2).LengthMax(8)
	invalidSchema := String().LengthMin(5).LengthMax(8)

	if err := validSchema.Validate(name); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(name); err != nil {
		t.Error(err)
	}
}

func Test_StringEnum(t *testing.T) {
	status := "PENDING"

	validSchema := String().Enum([]string{"PENDING", "DONE", "IN_PROGRESS"})
	invalidSchema := String().Enum([]string{"DONE", "CANCELLED"})

	if err := validSchema.Validate(status); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(status); err != nil {
		t.Error(err)
	}
}

type regexTestCase struct {
	value  string
	valid  bool
	schema *stringSchema
}

func Test_Regex(t *testing.T) {
	emailSchema := String().Email()
	uuidSchema := String().UUID()
	urlSchema := String().URL()

	var tests = []regexTestCase{
		// Email tests
		{"user@example.com", true, emailSchema},
		{"first.last@domain.co.uk", true, emailSchema},
		{"user+tag@sub.domain.org", true, emailSchema},
		{"user@", false, emailSchema},
		{"@example.com", false, emailSchema},
		{"user@.com", false, emailSchema},
		{"user@domain", false, emailSchema},
		{"plainaddress", false, emailSchema},
		{"", false, emailSchema},

		// UUID tests
		{"123e4567-e89b-12d3-a456-426614174000", true, uuidSchema},
		{"550e8400-e29b-41d4-a716-446655440000", true, uuidSchema},
		{"550e8400e29b41d4a716446655440000", false, uuidSchema},
		{"g23e4567-e89b-12d3-a456-426614174000", false, uuidSchema},
		{"123e4567-e89b-12d3-a456", false, uuidSchema},
		{"", false, uuidSchema},

		// URL tests
		{"http://example.com", true, urlSchema},
		{"https://sub.domain.com/path?query=1#fragment", true, urlSchema},
		{"ftp://files.example.com/resource.zip", true, urlSchema},
		{"http://localhost:8080/path", true, urlSchema},
		{"https://192.168.1.1:8080", true, urlSchema},
		{"example.com", false, urlSchema},
		{"http:/example.com", false, urlSchema},
		{"://example.com", false, urlSchema},
		{"", false, urlSchema},
	}

	for _, tt := range tests {
		isValid := tt.schema.Validate(tt.value) == nil

		assert.Equal(t, isValid, tt.valid)
	}
}
