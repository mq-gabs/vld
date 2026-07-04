package schema

type schemaGeneric struct {
	baseSchema[any]
}

type SchemaGeneric interface {
	Schema[any]

	Custom(fn Validator[any]) SchemaGeneric
}

func Generic() SchemaGeneric {
	return &schemaGeneric{
		baseSchema: newBaseSchema[any](),
	}
}

func (sg *schemaGeneric) Custom(fn Validator[any]) SchemaGeneric {
	sg.appendValidator(fn)

	return sg
}
