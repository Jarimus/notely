package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		inputHeaders http.Header
		want         string
		hasError     bool
	}

	tests := []test{
		{
			inputHeaders: http.Header{"Authorization": []string{"ApiKey thisIsTheApiKey"}},
			want:         "thisIsTheApiKey",
			hasError:     false,
		},
		{
			inputHeaders: http.Header{"Authorization": []string{"notApiKey thisIsTheApiKey"}},
			want:         "",
			hasError:     true,
		},
		{
			inputHeaders: http.Header{"wrong header": []string{"ApiKey thisIsTheApiKey"}},
			hasError:     true,
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.inputHeaders)

		// Compare wanted errors to actual errors
		if err != nil && !tc.hasError {
			t.Fatalf("did not expect an error\nerror: %s\n", err)
		} else if tc.hasError && err == nil {
			t.Fatalf("expected an error, did not get one.\ninput: %s\n", tc.inputHeaders)
		}

		// Compare got value to want value
		if got != tc.want {
			t.Fatalf("wrong output\nwant: %s\ngot: %s", tc.want, got)
		}
	}
}
