package schema

import (
	"testing"
)

type User struct {
	name string
	age  int
}

func Test_Struct(t *testing.T) {

	u := User{
		name: "John Doe",
		age:  33,
	}

	validSchema := Struct(func(b Builder, u *User) {
		b.F(u.name, String().LengthMin(4).LengthMax(32))
		b.F(u.age, Number[int]().Max(44))
	})
	invalidSchema := Struct(func(b Builder, u *User) {
		b.F(u.name, String().LengthMax(4))
		b.F(u.age, Number[int]().Min(44))
	})

	if err := validSchema.Validate(&u); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(&u); err == nil {
		t.Error(errExpectedError)
	}
}

type Group struct {
	name   string
	member *User
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

	validUserSchema := Struct(func(b Builder, u *User) {
		b.F(u.name, String().LengthMax(32))
		b.F(u.age, Number[int]().Min(18))
	})

	invalidUserSchema := Struct(func(b Builder, u *User) {
		b.F(u.name, String().LengthMin(32))
		b.F(u.age, Number[int]().Max(21))
	})

	validSchema := Struct(func(b Builder, g *Group) {
		b.F(g.name, String().LengthMax(12))
		b.F(g.member, validUserSchema)
	})

	invalidSchema := Struct(func(b Builder, g *Group) {
		b.F(g.name, String().LengthMax(4))
		b.F(g.member, invalidUserSchema)
	})

	if err := validSchema.Validate(&g); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(&g); err == nil {
		t.Error(errExpectedError)
	}
}
