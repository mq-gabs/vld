# vld

`vld` is a small, generic Go library for validating and transforming values with ordinary functions.

Validators have the shape:

```go
func(T) error
```

Parsers have the shape:

```go
func(string) string
```

This keeps the library composable: use the built-in validators, pass your own functions, or combine both into a reusable validator.

## Requirements

- Go 1.26.4 or newer

## Installation

```bash
go get github.com/mq-gabs/vld
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/mq-gabs/vld/validator/vstring"
)

func main() {
	validateUsername := vstring.String(
		"username",
		vstring.Required(),
		vstring.MinLen(3),
		vstring.MaxLen(20),
		vstring.Alphanumeric(),
	)

	if err := validateUsername("ab!"); err != nil {
		fmt.Println(err)
		// [name=username]:required min length,value must be alphanumeric
	}
}
```

A grouped validator runs every supplied validator and returns all failures. This makes it possible to show a user more than one correction at a time.

## How It Works

Each validator is a function that returns `nil` for a valid value or an error for an invalid value. The typed validator aliases make APIs easier to read:

```go
type StringValidator func(string) error
type NumberValidator[T Numeric] func(T) error
type SliceValidator[T any] func([]T) error
type MapValidator[K comparable, V any] func(map[K]V) error
type TimeValidator func(time.Time) error
```

The package-specific grouping functions execute validators in order and collect their errors:

- `vstring.String(name, validators...)`
- `vnumber.Number(name, validators...)`
- `vslices.Slice(name, validators...)`
- `vmap.Map(name, validators...)`
- `vtime.Time(name, validators...)`

The `name` is included in the formatted error. Nested slice and map validation errors include their index or key. Map error ordering is not deterministic because Go map iteration order is not deterministic.

Errors remain compatible with the standard library's `errors.Is` and `errors.As`, including errors collected by grouped validators.

## Generic Validator Utilities

The `validator` package provides utilities that work with any `func(T) error`.

### Conditional validation

Use `When` when the condition is already known while building the validator:

```go
validatePassword := validator.When(
	isPasswordRequired,
	vstring.MinLen(12),
)
```

Use `WhenFunc` when the condition depends on the value being validated:

```go
validateCode := validator.WhenFunc(
	func(value string) bool { return value != "" },
	vstring.Pattern(`^[A-Z0-9]+$`),
)
```

When the condition is false, the returned validator succeeds without calling the wrapped validator.

### Alternatives with `OneOf`

`OneOf` succeeds as soon as one validator succeeds. If every validator fails, it returns the collected errors:

```go
validateContact := validator.OneOf(
	vstring.Email(),
	vstring.Pattern(`^\+?[0-9]{10,15}$`),
)
```

This is useful for values that may be represented in more than one format.

## String Validation

Package: `github.com/mq-gabs/vld/validator/vstring`

Available validators:

- `Required`
- `MinLen`, `MaxLen`
- `Email`, `UUIDv4`, `Pattern`
- `Alphanumeric`
- `StartsWith`, `EndsWith`, `NotContains`
- `OneOf`
- `NoWhitespace`
- `LowerCase`, `UpperCase`

Example:

```go
validateEmail := vstring.String(
	"email",
	vstring.Required(),
	vstring.Email(),
)

err := validateEmail("person@example.com")
```

`MinLen` and `MaxLen` measure byte length, not rune length. `Email` uses a deliberately simple regular expression rather than complete RFC email validation. `UUIDv4` accepts the UUID format case-insensitively.

## Number Validation

Package: `github.com/mq-gabs/vld/validator/vnumber`

`Numeric` supports `int`, `int32`, `int64`, `float32`, and `float64`, including named types with those underlying types.

Available validators:

- `MinValue`, `MaxValue`, `InRange`
- `IsPositive`, `IsNegative`
- `IsZero`, `IsNonZero`
- `IsEven`, `IsOdd`
- `OneOf`

Example:

```go
validateAge := vnumber.Number(
	"age",
	vnumber.MinValue(18),
	vnumber.MaxValue(120),
)

err := validateAge(42)
```

`IsEven` and `IsOdd` are available for all types accepted by `Numeric`, but their parity check converts values to `int64`. They should therefore be used with integer values.

## Slice Validation

Package: `github.com/mq-gabs/vld/validator/vslices`

Available validators:

- `Required`
- `MinLen`, `MaxLen`, `ExactLen`, `InRange`
- `Each`
- `Contains`, `NotContains`
- `NoDuplicates`

`Each` applies an ordinary element validator and reports failing indexes:

```go
validateTags := vslices.Slice(
	"tags",
	vslices.Required[string](),
	vslices.MinLen[string](1),
	vslices.Each[string](vstring.MinLen(2)),
)

err := validateTags([]string{"go", "x"})
// The second item is reported with its index.
```

Slice length bounds are checked when the validator is created. Negative bounds and an inverted `InRange` panic because they are invalid configuration.

## Map Validation

Package: `github.com/mq-gabs/vld/validator/vmap`

Available validators:

- `Required`
- `MinLen`, `MaxLen`, `ExactLen`, `InRange`
- `ContainsKey`, `NotContainsKey`
- `HasKeys`, `NotHasKeys`
- `EachValue`, `EachKey`, `Each`

Example:

```go
validateConfig := vmap.Map[string, string](
	"config",
	vmap.HasKeys[string, string]("host", "port"),
	vmap.EachValue[string, string](vstring.Required()),
)

err := validateConfig(map[string]string{
	"host": "localhost",
	"port": "5432",
})
```

Map validators accept any comparable key type and any value type. `EachValue` and `EachKey` annotate failures with the relevant map key.

## Time Validation

Package: `github.com/mq-gabs/vld/validator/vtime`

Available validators:

- `NotZero`, `Midnight`
- `After`, `NotAfter`, `Before`, `NotBefore`, `Between`
- `Past`, `Future`, `Today`, `SameDate`
- `Weekday`, `Weekend`, `MatchWeekday`

Example:

```go
validateRelease := vtime.Time(
	"release_at",
	vtime.NotZero(),
	vtime.Future(),
)

err := validateRelease(time.Now().Add(24 * time.Hour))
```

`Between` is exclusive at both endpoints. `Past` and `Future` compare against the time at which validation runs. `Today` and `SameDate` compare calendar fields without converting time zones.

## String Parsing

Package: `github.com/mq-gabs/vld/parser/pstring`

Parsers transform strings and never return validation errors. `pstring.String` applies the supplied parsers sequentially, passing each result to the next parser:

```go
cleanName := pstring.String(
	"name",
	pstring.TrimSpace(),
	pstring.CollapseSpace(),
	pstring.Capitalize(),
)

name := cleanName("  ada   lovelace  ")
// "Ada lovelace"
```

Available parser operations include:

- Trimming: `TrimSpace`, `Trim`, `TrimPrefix`, `TrimSuffix`, `TrimNewlines`
- Removing or replacing content: `NoSpace`, `NoChar`, `Replace`, `RemoveNewlines`, `RemoveWhitespace`, `RemoveNumbers`, `RemoveLetters`, `RemoveSpecial`
- Case and ordering: `Lower`, `Upper`, `Capitalize`, `Reverse`
- Size and repetition: `Repeat`, `PadLeft`, `PadRight`, `Truncate`
- Formatting: `Surround`, `Quote`, `SingleQuote`

Parser configuration functions such as `Repeat`, `PadLeft`, `PadRight`, and `Truncate` panic for negative sizes. `Reverse` is rune-aware; `Capitalize`, padding, and truncation use byte lengths.

## Custom Validators and Parsers

Because the public API is function-based, custom rules need no special interface:

```go
validateSlug := func(value string) error {
	if value == "" {
		return errors.New("slug is required")
	}
	return nil
}

normalizeSlug := func(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}
```

Custom functions can be passed directly to the grouping functions or composed with `validator.When`, `validator.WhenFunc`, and `validator.OneOf`.

## Packages

```text
validator/                 Generic validator utilities and common errors
validator/vstring/         String validators
validator/vnumber/         Numeric validators
validator/vslices/         Slice validators
validator/vmap/            Map validators
validator/vtime/           time.Time validators
parser/pstring/            String parsers
internal/utils/            Internal error and time helpers
```

## Testing

Run the complete test suite with:

```bash
go test ./...
```
