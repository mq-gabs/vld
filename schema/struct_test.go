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

	validSchema := Struct[User](func(u *User) []Tuple[any] {
		return []Tuple[any]{
			T(u.name, String().LengthMin(4).LengthMax(32)),
			T(u.age, Number[int]().Max(44)),
		}
	})
	invalidSchema := Struct[User](func(u *User) []Tuple[any] {
		return []Tuple[any]{
			T(u.name, String().LengthMax(4)),
			T(u.age, Number[int]().Min(44)),
		}
	})

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

func Test_StructInsideStruct(t *testing.T) {
	u := User{
		name: "Bob Smith",
		age:  23,
	}

	g := Group{
		name:   "Worker",
		member: &u,
	}

	validUserSchema := Struct[User](func(u *User) []Tuple[any] {
		return []Tuple[any]{
			T(u.name, String().LengthMax(32)),
			T(u.age, Number[int]().Min(18)),
		}
	})

	invalidUserSchema := Struct[User](func(u *User) []Tuple[any] {
		return []Tuple[any]{
			T(u.name, String().LengthMin(32)),
			T(u.age, Number[int]().Max(21)),
		}
	})

	validSchema := Struct[Group](func(g *Group) []Tuple[any] {
		return []Tuple[any]{
			T(g.name, String().LengthMax(12)),
			T(g.member, validUserSchema),
		}
	})

	invalidSchema := Struct[Group](func(g *Group) []Tuple[any] {
		return []Tuple[any]{
			T(g.name, String().LengthMax(4)),
			T(g.member, invalidUserSchema),
		}
	})

	if err := validSchema.Validate(g); err != nil {
		t.Error(err)
	}
	if err := invalidSchema.Validate(g); err != nil {
		t.Error(err)
	}
}
