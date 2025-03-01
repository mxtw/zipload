package upload

import (
	"testing"
)

func TestOptions_toHeaders(t *testing.T) {
	tests := []struct {
		name     string
		options  Options
		expected map[string]string
	}{
		{
			"All fields set",
			Options{
				DeletesAt:               "2025-01-01T00:00:00Z",
				Domain:                  "example.com",
				FileExtension:           "png",
				Filename:                "file",
				Folder:                  3,
				Format:                  FormatUUID,
				ImageCompressionPercent: 80,
				MaxViews:                10,
				OriginalName:            true,
				Password:                "securepassword",
			},
			map[string]string{
				"x-zipline-deletes-at":                "2025-01-01T00:00:00Z",
				"x-zipline-domain":                    "example.com",
				"x-zipline-file-extension":            "png",
				"x-zipline-filename":                  "file",
				"x-zipline-folder":                    "3",
				"x-zipline-format":                    "uuid",
				"x-zipline-image-compression-percent": "80",
				"x-zipline-max-views":                 "10",
				"x-zipline-original-name":             "true",
				"x-zipline-password":                  "securepassword",
			},
		},
		{
			"Only required fields set",
			Options{
				Format: FormatRandom,
			},
			map[string]string{
				"x-zipline-format": "random",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := tt.options.toHeaders()
			for key, expectedValue := range tt.expected {
				if headers.Get(key) != expectedValue {
					t.Errorf("Expected %q for header %q, got %q", expectedValue, key, headers.Get(key))
				}
			}
		})
	}
}
