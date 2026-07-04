package schema

type schemaGeneric struct {
	baseSchema[any]
}

type SchemaGeneric interface {
	Schema[any]

	Custom(fn Validator[any]) SchemaGeneric
	Clone() SchemaGeneric
}

func Generic() SchemaGeneric {
	return &schemaGeneric{
		baseSchema: newBaseSchema[any](),
	}
}

func (sg *schemaGeneric) Clone() SchemaGeneric {
	return &schemaGeneric{
		baseSchema: sg.baseSchema.clone(),
	}
}

func (sg *schemaGeneric) Custom(fn Validator[any]) SchemaGeneric {
	sg.appendValidator(fn)

	return sg
}
