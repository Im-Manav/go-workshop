package fuzzing

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		input        string
		entity, term string
	}{
		{
			input:  "/search/abc/def",
			entity: "abc",
			term:   "def",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actualEntity, actualTerm, err := Parse(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
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
	f.Fuzz(func(t *testing.T, path string) {
		_, _, err := Parse(path)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
