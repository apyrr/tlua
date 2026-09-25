package stringutil

import "testing"

func TestLuaCasing(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		got  string
		want string
	}{
		{name: "ascii uppercase", got: ToUpperLua("hello, World"), want: "HELLO, WORLD"},
		{name: "ascii lowercase", got: ToLowerLua("HELLO, World"), want: "hello, world"},
		{name: "utf-8 letters are kept", got: ToUpperLua("éß"), want: "éß"},
		{name: "high bytes are kept", got: ToUpperLua("\xffa\xfe"), want: "\xffA\xfe"},
		{name: "unchanged string", got: ToLowerLua("abc"), want: "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.got != tt.want {
				t.Fatalf("got %q, want %q", tt.got, tt.want)
			}
		})
	}
}
