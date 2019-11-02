package scanner

import (
	"reflect"
	"testing"
)

func TestScanner_Tokenise(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{
			name:  "integer",
			input: "123",
			want:  []string{"123"},
		},
		{
			name:  "decimal",
			input: "12.90",
			want:  []string{"12.90"},
		},
		{
			name:  "negative number",
			input: "-392",
			want:  []string{"-", "392"},
		},
		{
			name:  "simple expression",
			input: "2xx3+4",
			want:  []string{"2", "xx", "3", "+", "4"},
		},
		{
			name:  "simple algebraic expression",
			input: "x+2",
			want:  []string{"x", "+", "2"},
		},
		{
			name:  "vars without space",
			input: "xyz",
			want:  []string{"x", "y", "z"},
		},
		{
			name:  "vars with space",
			input: "x y z  ",
			want:  []string{"x", "y", "z"},
		},
		{
			name:  "trig",
			input: "sincostheta",
			want:  []string{"sin", "cos", "theta"},
		},
		{
			name:  "sqrt",
			input: "sqrt(2x)",
			want:  []string{"sqrt", "(", "2", "x", ")"},
		},
		{
			name:  "algebraic fraction",
			input: "(2x^2-1)/(xy)",
			want:  []string{"(", "2", "x", "^", "2", "-", "1", ")", "/", "(", "x", "y", ")"},
		},
	}
	s := newScanner()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Tokenise(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scanner.Tokenise() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.Tokenise() = %v, want %v", got, tt.want)
			}
		})
	}
}
