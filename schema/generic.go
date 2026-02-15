package schema

type SchemaGeneric struct {
	baseSchema[any]
}

func Generic() *SchemaGeneric {
	return &SchemaGeneric{
		baseSchema: newBaseSchema[any](),
	}
}

func (sg *SchemaGeneric) Custom(fn Validator[any]) *SchemaGeneric {
	sg.appendValidator(fn)

	return sg
}
