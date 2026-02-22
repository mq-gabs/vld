package main

import (
	"errors"
	"fmt"

	"github.com/mq-gabs/dilav/schema"
)

type User struct {
	ID           int
	Name         string
	Email        string
	Password     string
	PersonalPage string
	Products     []string
}

func (u User) SchemaJSON() map[string]any {
	return map[string]any{
		"id":           u.ID,
		"name":         u.Name,
		"email":        u.Email,
		"password":     u.Password,
		"personalPage": u.PersonalPage,
		"products":     u.Products,
	}
}

func main() {
	idSchema := schema.Number[int]().Positive()
	nameSchema := schema.String().LengthMin(4).LengthMax(64)
	emailSchema := schema.String().Email()
	passwordSchema := schema.String().LengthMin(8).LengthMax(64)
	personalPageSchema := schema.String().URL()
	productSchema := schema.String().Enum([]string{"A", "B", "C", "D", "E"})
	userProductsSchema := schema.Slice[string]().LengthMin(1).LengthMax(5).Custom(func(s []string) error {
		var err error
		for _, v := range s {
			if e := productSchema.Validate(v); e != nil {
				return errors.Join(err, e)
			}
		}

		return err
	})

	userSchema := schema.Struct(func(b schema.Builder, u *User) {
		b.F(u.ID, idSchema)
		b.F(u.Name, nameSchema)
		b.F(u.Email, emailSchema)
		b.F(u.Password, passwordSchema)
		b.F(u.PersonalPage, personalPageSchema)
		b.F(u.Products, userProductsSchema)
	})

	validUser := User{
		ID:           1,
		Name:         "João Silva",
		Email:        "joao.silva@email.com",
		Password:     "strongPass123",
		PersonalPage: "https://example.com/profile/joao",
		Products:     []string{"A", "C", "E"},
	}
	invalidUser := User{
		ID:           -10,                     // inválido (não positivo)
		Name:         "Ana",                   // inválido (< 4 chars)
		Email:        "email-invalido",        // inválido (não é email)
		Password:     "123",                   // inválido (< 8 chars)
		PersonalPage: "not-a-url",             // inválido (não é URL)
		Products:     []string{"A", "X", "Z"}, // inválido (X e Z não estão no enum)
	}

	if err := userSchema.Validate(invalidUser); err != nil {
		fmt.Printf("invalid schema err: %v\n", err)
	}
	if err := userSchema.Validate(validUser); err != nil {
		fmt.Printf("valid schema err: %v\n", err)
	}
}
