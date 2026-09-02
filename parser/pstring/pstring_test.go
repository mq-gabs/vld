package pstring

import "testing"

func Test_String_Grouped(t *testing.T) {
	parser := String("process", TrimSpace(), Lower(), NoSpace())

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with spaces and mixed case",
			input:    "  Hello World  ",
			expected: "helloworld",
		},
		{
			name:     "single word",
			input:    "  HELLO  ",
			expected: "hello",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_TrimSpace(t *testing.T) {
	parser := TrimSpace()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with leading and trailing spaces",
			input:    "  hello  ",
			expected: "hello",
		},
		{
			name:     "with tabs and newlines",
			input:    "\t\nhello\n\t",
			expected: "hello",
		},
		{
			name:     "no spaces",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Trim(t *testing.T) {
	parser := Trim("-")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "trim dashes",
			input:    "---hello---",
			expected: "hello",
		},
		{
			name:     "mixed dashes and letters",
			input:    "--he-llo--",
			expected: "he-llo",
		},
		{
			name:     "no trim characters",
			input:    "hello",
			expected: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_TrimPrefix(t *testing.T) {
	parser := TrimPrefix("http://")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with prefix",
			input:    "http://example.com",
			expected: "example.com",
		},
		{
			name:     "without prefix",
			input:    "example.com",
			expected: "example.com",
		},
		{
			name:     "prefix multiple times",
			input:    "http://http://example.com",
			expected: "http://example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_TrimSuffix(t *testing.T) {
	parser := TrimSuffix(".txt")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with suffix",
			input:    "document.txt",
			expected: "document",
		},
		{
			name:     "without suffix",
			input:    "document.pdf",
			expected: "document.pdf",
		},
		{
			name:     "suffix multiple times",
			input:    "document.txt.txt",
			expected: "document.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_NoSpace(t *testing.T) {
	parser := NoSpace()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with spaces",
			input:    "hello world test",
			expected: "helloworldtest",
		},
		{
			name:     "no spaces",
			input:    "helloworld",
			expected: "helloworld",
		},
		{
			name:     "only spaces",
			input:    "   ",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_NoChar(t *testing.T) {
	parser := NoChar("e")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with target char",
			input:    "hello",
			expected: "hllo",
		},
		{
			name:     "no target char",
			input:    "xyzabc",
			expected: "xyzabc",
		},
		{
			name:     "all target char",
			input:    "eeee",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Replace(t *testing.T) {
	parser := Replace("old", "new")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single replacement",
			input:    "old value",
			expected: "new value",
		},
		{
			name:     "multiple replacements",
			input:    "old old old",
			expected: "new new new",
		},
		{
			name:     "no match",
			input:    "something else",
			expected: "something else",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Lower(t *testing.T) {
	parser := Lower()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "uppercase",
			input:    "HELLO",
			expected: "hello",
		},
		{
			name:     "mixed case",
			input:    "HeLLo",
			expected: "hello",
		},
		{
			name:     "lowercase",
			input:    "hello",
			expected: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Upper(t *testing.T) {
	parser := Upper()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "lowercase",
			input:    "hello",
			expected: "HELLO",
		},
		{
			name:     "mixed case",
			input:    "HeLLo",
			expected: "HELLO",
		},
		{
			name:     "uppercase",
			input:    "HELLO",
			expected: "HELLO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Capitalize(t *testing.T) {
	parser := Capitalize()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "lowercase",
			input:    "hello",
			expected: "Hello",
		},
		{
			name:     "uppercase",
			input:    "HELLO",
			expected: "HELLO",
		},
		{
			name:     "mixed case",
			input:    "hELLO",
			expected: "HELLO",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Reverse(t *testing.T) {
	parser := Reverse()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple word",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "unicode characters",
			input:    "café",
			expected: "éfac",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Repeat(t *testing.T) {
	parser := Repeat(3)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "repeat word",
			input:    "ab",
			expected: "ababab",
		},
		{
			name:     "repeat single char",
			input:    "x",
			expected: "xxx",
		},
		{
			name:     "repeat empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_PadLeft(t *testing.T) {
	parser := PadLeft(10, "-")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "pad short string",
			input:    "hello",
			expected: "-----hello",
		},
		{
			name:     "string at length",
			input:    "helloworld",
			expected: "helloworld",
		},
		{
			name:     "string longer than length",
			input:    "hello world",
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_PadRight(t *testing.T) {
	parser := PadRight(10, ".")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "pad short string",
			input:    "hello",
			expected: "hello.....",
		},
		{
			name:     "string at length",
			input:    "helloworld",
			expected: "helloworld",
		},
		{
			name:     "string longer than length",
			input:    "hello world",
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Truncate(t *testing.T) {
	parser := Truncate(5)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "truncate long string",
			input:    "hello world",
			expected: "hello",
		},
		{
			name:     "string shorter than length",
			input:    "hi",
			expected: "hi",
		},
		{
			name:     "string at length",
			input:    "hello",
			expected: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_RemoveNewlines(t *testing.T) {
	parser := RemoveNewlines()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with unix newlines",
			input:    "hello\nworld",
			expected: "helloworld",
		},
		{
			name:     "with windows newlines",
			input:    "hello\r\nworld",
			expected: "helloworld",
		},
		{
			name:     "multiple newlines",
			input:    "a\n\n\nb",
			expected: "ab",
		},
		{
			name:     "no newlines",
			input:    "helloworld",
			expected: "helloworld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_RemoveWhitespace(t *testing.T) {
	parser := RemoveWhitespace()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with various whitespace",
			input:    "hello \t world \n test",
			expected: "helloworldtest",
		},
		{
			name:     "no whitespace",
			input:    "helloworld",
			expected: "helloworld",
		},
		{
			name:     "only whitespace",
			input:    "   \t\n  ",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_CollapseSpace(t *testing.T) {
	parser := CollapseSpace()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "multiple spaces",
			input:    "hello    world",
			expected: "hello world",
		},
		{
			name:     "tabs and spaces",
			input:    "hello\t\tworld",
			expected: "hello world",
		},
		{
			name:     "single space",
			input:    "hello world",
			expected: "hello world",
		},
		{
			name:     "leading and trailing",
			input:    "  hello world  ",
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_TrimNewlines(t *testing.T) {
	parser := TrimNewlines()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with leading and trailing newlines",
			input:    "\n\nhello\n\n",
			expected: "hello",
		},
		{
			name:     "windows line endings",
			input:    "\r\nhello\r\n",
			expected: "hello",
		},
		{
			name:     "newlines in middle",
			input:    "hello\n\nworld",
			expected: "hello\n\nworld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_RemoveNumbers(t *testing.T) {
	parser := RemoveNumbers()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with numbers",
			input:    "hello123world456",
			expected: "helloworld",
		},
		{
			name:     "no numbers",
			input:    "helloworld",
			expected: "helloworld",
		},
		{
			name:     "only numbers",
			input:    "123456",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_RemoveLetters(t *testing.T) {
	parser := RemoveLetters()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with letters",
			input:    "hello123world456",
			expected: "123456",
		},
		{
			name:     "only letters",
			input:    "helloworld",
			expected: "",
		},
		{
			name:     "no letters",
			input:    "123456",
			expected: "123456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_RemoveSpecial(t *testing.T) {
	parser := RemoveSpecial()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with special chars",
			input:    "hello-world_123!@#",
			expected: "helloworld123",
		},
		{
			name:     "only alphanumeric",
			input:    "helloworld123",
			expected: "helloworld123",
		},
		{
			name:     "only special chars",
			input:    "!@#$%^&*()",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Surround(t *testing.T) {
	parser := Surround("[", "]")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "surround word",
			input:    "hello",
			expected: "[hello]",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_Quote(t *testing.T) {
	parser := Quote()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "quote word",
			input:    "hello",
			expected: `"hello"`,
		},
		{
			name:     "quote with spaces",
			input:    "hello world",
			expected: `"hello world"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_SingleQuote(t *testing.T) {
	parser := SingleQuote()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single quote word",
			input:    "hello",
			expected: "'hello'",
		},
		{
			name:     "single quote with spaces",
			input:    "hello world",
			expected: "'hello world'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}

func Test_ComplexParsing_Chain(t *testing.T) {
	// Complex chain: trim spaces, remove special chars, capitalize, pad, surround
	parser := String("process",
		TrimSpace(),
		RemoveSpecial(),
		Capitalize(),
		Surround("<", ">"),
	)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "complex chain",
			input:    "  hello-world  ",
			expected: "<Helloworld>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}
