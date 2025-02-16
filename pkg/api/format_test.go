package api

import (
	"testing"
)

func TestFormatFlag_Set(t *testing.T) {
	testCases := []struct {
		input    string
		expected Format
		wantErr  bool
	}{
		{"uuid", FormatUUID, false},
		{"date", FormatDate, false},
		{"random", FormatRandom, false},
		{"name", FormatName, false},
		{"invalid", 0, true},
	}

	for _, tc := range testCases {
		flag := &FormatFlag{}
		err := flag.Set(tc.input)

		if (err != nil) != tc.wantErr {
			t.Errorf("Set(%q) unexpected error: %v", tc.input, err)
		}
		if (err == nil) && tc.expected != flag.Value {
			t.Errorf("Set(%q) = %v, want %v", tc.input, flag.Value, tc.expected)
		}
	}
}

func TestFormatFlag_String(t *testing.T) {
	testCases := []struct {
		value    Format
		expected string
	}{
		{FormatUUID, "uuid"},
		{FormatDate, "date"},
		{FormatName, "name"},
		{FormatRandom, "random"},
	}

	for _, tc := range testCases {
		flag := &FormatFlag{tc.value}
		result := flag.String()
		if result != tc.expected {
			t.Errorf("String() = %s, want %s", result, tc.expected)
		}
	}
}
