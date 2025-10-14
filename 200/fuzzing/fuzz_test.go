package fuzzing

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		input  string
		entity string
		term   string
		ok     bool
	}{
		{
			input:  "/search/abc/def",
			entity: "abc",
			term:   "def",
			ok:     true,
		},
		{
			input:  "/rch/xyz/123",
			entity: "",
			term:   "",
			ok:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actualEntity, actualTerm, ok := Parse(test.input)
			if ok != test.ok {
				t.Errorf("expected ok %v, got %v", test.ok, ok)
			}
			if actualEntity != test.entity {
				t.Errorf("expected entity %q, got %q", test.entity, actualEntity)
			}
			if actualTerm != test.term {
				t.Errorf("expected term %q, got %q", test.term, actualTerm)
			}
		})
	}
}

func FuzzParse(f *testing.F) {
	f.Add("/search/abc/def")
	f.Add("/search/abc/def/ghi")
	f.Add("/rch/xyz/123")
	f.Fuzz(func(t *testing.T, path string) {
		_, _, _ = Parse(path)
	})
}
