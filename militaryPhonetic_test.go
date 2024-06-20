package MilitaryPhonetic

import (
	"testing"
)

// TestConvertTextToPhonetic tests the ConvertTextToPhonetic function
func TestConvertTextToPhonetic(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"A", "Alpha"},
		{"B", "Bravo"},
		{"C", "Charlie"},
		{"Hello World", "Hotel Echo Lima Lima Oscar Space Whiskey Oscar Romeo Lima Delta"},
		{"GoLang 123", "Golf Oscar Lima Alpha November Golf Space One Two Three"},
		{"", ""},
		{"123", "One Two Three"},
		{"!@#", "! @ #"},
	}

	for _, test := range tests {
		result := ConvertTextToPhonetic(test.input)
		if result != test.expected {
			t.Errorf("ConvertTextToPhonetic(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
