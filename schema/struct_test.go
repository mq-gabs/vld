package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		Field("name", String().MinLength(4).MaxLength(32)).
		Field("age", Number[int]().Max(44))
	invalidSchema := Struct[User]().
		Field("name", String().MaxLength(4)).
		Field("age", Number[int]().Min(44))

	err1 := validSchema.Validate(u)
	t.Log(err1)
	err2 := invalidSchema.Validate(u)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
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
		Field("name", String().MaxLength(32)).
		Field("age", Number[int]().Min(18))

	invalidUserSchema := Struct[User]().
		Field("name", String().MinLength(32)).
		Field("age", Number[int]().Max(21))

	validSchema := Struct[Group]().
		Field("name", String().MaxLength(12)).
		Field("member", validUserSchema)
	invalidSchema := Struct[Group]().
		Field("name", String().MaxLength(4)).
		Field("member", invalidUserSchema)

	err1 := validSchema.Validate(g)
	t.Log(err1)
	err2 := invalidSchema.Validate(g)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}
