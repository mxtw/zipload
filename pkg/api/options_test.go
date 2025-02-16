package api

import (
	"testing"
)

func TestOptions_ToHeaders(t *testing.T) {
	tests := []struct {
		name     string
		options  Options
		expected map[string]string
	}{
		{
			"All fields set",
			Options{
				Format:                  FormatUUID,
				ImageCompressionPercent: 80,
				ExpiresAt:               "2025-01-01T00:00:00Z",
				Password:                "securepassword",
				ZeroWidthSpace:          true,
				Embed:                   true,
				MaxViews:                10,
				UploadText:              true,
				XZiplineFilename:        "file.txt",
				OverrideDomain:          "example.com",
				XZiplineFolder:          5,
			},
			map[string]string{
				"Format":                    "uuid",
				"Image-Compression-Percent": "80",
				"Expires-At":                "2025-01-01T00:00:00Z",
				"Password":                  "securepassword",
				"Zws":                       "true",
				"Embed":                     "true",
				"Max-Views":                 "10",
				"UploadText":                "true",
				"X-Zipline-Filename":        "file.txt",
				"Override-Domain":           "example.com",
				"X-Zipline-Folder":          "5",
			},
		},
		{
			"Only required fields set",
			Options{
				Format: FormatRandom,
			},
			map[string]string{
				"Format": "random",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := tt.options.ToHeaders()
			for key, expectedValue := range tt.expected {
				if headers.Get(key) != expectedValue {
					t.Errorf("Expected %q for header %q, got %q", expectedValue, key, headers.Get(key))
				}
			}
		})
	}
}
