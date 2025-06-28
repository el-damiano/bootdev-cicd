package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	cases := map[string]struct {
		header  http.Header
		want    string
		wantErr bool
	}{
		"valid header": {
			header: http.Header{
				"Authorization": []string{"ApiKey token_here"},
			},
			want:    "token_here",
			wantErr: false,
		},
		"malformed Authorization header": {
			header: http.Header{
				"Authorization": []string{"yup"},
			},
			want:    "",
			wantErr: true,
		},
		"missing Authorization header": {
			header:  http.Header{},
			want:    "",
			wantErr: true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			got, err := GetAPIKey(c.header)
			if (err != nil) != c.wantErr {
				t.Errorf("wantErr %v. Token error = %v", c.wantErr, err)
				t.Fail()
			}

			if got != c.want {
				t.Errorf("Token got = %v, want %v", got, c.want)
				t.Fail()
			}
		})
	}

}
