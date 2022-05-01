package subseq

import (
	"testing"
)

func TestMaxSubstring(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s1: "какая то строка текста",
				s2: "другая строка другого",
			},
			want: " строка ",
		},
		{
			name: "test2",
			args: args{
				s1: "qwertyuiop",
				s2: "asdfghertyjklzxcvzxqertyuvzvb",
			},
			want: "ertyu",
		},
		{
			name: "empty",
			args: args{
				s1: "abc",
				s2: "def",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubstring(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("MaxSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubseq(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s1: "nematode knowledge",
				s2: "empty bottle",
			},
			want: "emt ole",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubsequence(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("MaxSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
