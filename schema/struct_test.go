package schema

import (
	"testing"
)

type User struct {
	name string
	age  int
}

func (u User) SchemaJSON() map[string]any {
	return map[string]any{
		"name": u.name,
		"age":  u.age,
	}
}

func Test_Struct(t *testing.T) {

	u := User{
		name: "John Doe",
		age:  33,
	}

	validSchema := Struct[User]().
		Field("name", String().LengthMin(4).LengthMax(32)).
		Field("age", Number[int]().Max(44))
	invalidSchema := Struct[User]().
		Field("name", String().LengthMax(4)).
		Field("age", Number[int]().Min(44))

	if err := validSchema.Validate(u); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(u); err != nil {
		t.Error(err)
	}
}

type Group struct {
	name   string
	member *User
}

func (g Group) SchemaJSON() map[string]any {
	return map[string]any{
		"name":   g.name,
		"member": g.member,
	}
}

func Test_StructInsideStruct(t *testing.T) {
	u := User{
		name: "Bob Smith",
		age:  23,
	}

	g := Group{
		name:   "Worker",
		member: &u,
	}

	validUserSchema := Struct[User]().
		Field("name", String().LengthMax(32)).
		Field("age", Number[int]().Min(18))

	invalidUserSchema := Struct[User]().
		Field("name", String().LengthMin(32)).
		Field("age", Number[int]().Max(21))

	validSchema := Struct[Group]().
		Field("name", String().LengthMax(12)).
		Field("member", validUserSchema)
	invalidSchema := Struct[Group]().
		Field("name", String().LengthMax(4)).
		Field("member", invalidUserSchema)

	if err := validSchema.Validate(g); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(g); err != nil {
		t.Error(err)
	}
}
