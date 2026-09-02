package pstring

import (
	"fmt"
	"strings"
	"unicode"
)

// StringParser is a function type for parsing/transforming strings
type StringParser func(string) string

// String groups multiple parsers into a single parser, applying them sequentially
func String(name string, parsers ...StringParser) StringParser {
	return func(s string) string {
		for _, parse := range parsers {
			s = parse(s)
		}

		return s
	}
}

// TrimSpace removes leading and trailing whitespace
func TrimSpace() StringParser {
	return func(s string) string {
		return strings.TrimSpace(s)
	}
}

// Trim removes the specified characters from both ends of the string
func Trim(cut string) StringParser {
	return func(s string) string {
		return strings.Trim(s, cut)
	}
}

// TrimPrefix removes the specified prefix from the string
func TrimPrefix(prefix string) StringParser {
	return func(s string) string {
		return strings.TrimPrefix(s, prefix)
	}
}

// TrimSuffix removes the specified suffix from the string
func TrimSuffix(suffix string) StringParser {
	return func(s string) string {
		return strings.TrimSuffix(s, suffix)
	}
}

// NoSpace removes all spaces from the string
func NoSpace() StringParser {
	return func(s string) string {
		return strings.ReplaceAll(s, " ", "")
	}
}

// NoChar removes all occurrences of the specified character
func NoChar(char string) StringParser {
	return func(s string) string {
		return strings.ReplaceAll(s, char, "")
	}
}

// Replace replaces all occurrences of old with new
func Replace(old, new string) StringParser {
	return func(s string) string {
		return strings.ReplaceAll(s, old, new)
	}
}

// Lower converts the string to lowercase
func Lower() StringParser {
	return func(s string) string {
		return strings.ToLower(s)
	}
}

// Upper converts the string to uppercase
func Upper() StringParser {
	return func(s string) string {
		return strings.ToUpper(s)
	}
}

// Capitalize capitalizes the first character of the string
func Capitalize() StringParser {
	return func(s string) string {
		if len(s) == 0 {
			return s
		}
		return strings.ToUpper(string(s[0])) + s[1:]
	}
}

// Reverse reverses the string
func Reverse() StringParser {
	return func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
}

// Repeat repeats the string n times
func Repeat(n int) StringParser {
	if n < 0 {
		panic("repeat count cannot be negative")
	}
	return func(s string) string {
		return strings.Repeat(s, n)
	}
}

// PadLeft pads the string on the left with the specified character to reach the desired length
func PadLeft(length int, char string) StringParser {
	if length < 0 {
		panic("padding length cannot be negative")
	}
	return func(s string) string {
		if len(s) >= length {
			return s
		}
		padding := strings.Repeat(char, length-len(s))
		return padding + s
	}
}

// PadRight pads the string on the right with the specified character to reach the desired length
func PadRight(length int, char string) StringParser {
	if length < 0 {
		panic("padding length cannot be negative")
	}
	return func(s string) string {
		if len(s) >= length {
			return s
		}
		padding := strings.Repeat(char, length-len(s))
		return s + padding
	}
}

// Truncate truncates the string to the specified length
func Truncate(length int) StringParser {
	if length < 0 {
		panic("truncate length cannot be negative")
	}
	return func(s string) string {
		if len(s) <= length {
			return s
		}
		return s[:length]
	}
}

// RemoveNewlines removes all newline characters from the string
func RemoveNewlines() StringParser {
	return func(s string) string {
		return strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), "\r", "")
	}
}

// RemoveWhitespace removes all whitespace characters from the string
func RemoveWhitespace() StringParser {
	return func(s string) string {
		return strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}
			return r
		}, s)
	}
}

// CollapseSpace collapses multiple consecutive spaces into a single space
func CollapseSpace() StringParser {
	return func(s string) string {
		return strings.Join(strings.Fields(s), " ")
	}
}

// TrimNewlines removes leading and trailing newlines
func TrimNewlines() StringParser {
	return func(s string) string {
		return strings.Trim(s, "\n\r")
	}
}

// RemoveNumbers removes all numeric characters from the string
func RemoveNumbers() StringParser {
	return func(s string) string {
		return strings.Map(func(r rune) rune {
			if unicode.IsDigit(r) {
				return -1
			}
			return r
		}, s)
	}
}

// RemoveLetters removes all alphabetic characters from the string
func RemoveLetters() StringParser {
	return func(s string) string {
		return strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) {
				return -1
			}
			return r
		}, s)
	}
}

// RemoveSpecial removes all non-alphanumeric characters from the string
func RemoveSpecial() StringParser {
	return func(s string) string {
		return strings.Map(func(r rune) rune {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				return -1
			}
			return r
		}, s)
	}
}

// Surround surrounds the string with the specified prefix and suffix
func Surround(prefix, suffix string) StringParser {
	return func(s string) string {
		return prefix + s + suffix
	}
}

// Quote surrounds the string with double quotes
func Quote() StringParser {
	return func(s string) string {
		return fmt.Sprintf(`"%s"`, s)
	}
}

// SingleQuote surrounds the string with single quotes
func SingleQuote() StringParser {
	return func(s string) string {
		return fmt.Sprintf("'%s'", s)
	}
}
