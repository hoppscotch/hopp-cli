package methods

import "testing"

var checkURLCases = []struct {
	name      string
	url       string
	expectErr bool
}{
	{
		name:      "valid http://",
		url:       "http://example.com",
		expectErr: false,
	},
	{
		name:      "valid https://",
		url:       "https://example.com",
		expectErr: false,
	},
	{
		name:      "empty url",
		url:       "",
		expectErr: true,
	},
	{
		name:      "invalid protocol",
		url:       "htp://example.com",
		expectErr: true,
	},
	{
		name:      "disallowed protocol",
		url:       "irc://example.com",
		expectErr: true,
	},
}

func Test_checkURL(t *testing.T) {
	for _, tt := range checkURLCases {
		out, err := checkURL(tt.url)
		if err != nil && !tt.expectErr {
			t.Errorf("%s :: %s", tt.name, err.Error())
		}
		if out != tt.url && !tt.expectErr {
			t.Errorf("URL mangled. Got %s - expected %s", out, tt.url)
		}
		if out != "" && err != nil && tt.expectErr {
			t.Errorf("Didn't fail when expected")
		}
	}
}
