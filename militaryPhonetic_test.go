package MilitaryPhonetic

import (
	"testing"
)

// TestConvertTextToPhonetic tests the ConvertTextToPhonetic function
func TestConvertTextToPhonetic(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Alpha":       {"A", "Alpha"},
		"Bravo":       {"B", "Bravo"},
		"Charlie":     {"C", "Charlie"},
		"Hello World": {"Hello World", "Hotel Echo Lima Lima Oscar Space Whiskey Oscar Romeo Lima Delta"},
		"GoLang 123":  {"GoLang 123", "Golf Oscar Lima Alpha November Golf Space One Two Three"},
		"no text":     {"", ""},
		"123":         {"123", "One Two Three"},
		"Symbols !@#": {"!@#", "! @ #"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ConvertTextToPhonetic(test.input)
			if result != test.expected {
				t.Errorf("ConvertTextToPhonetic(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}

}
