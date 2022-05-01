package levenshtein

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []Edit
	}{
		{
			name: "test1",
			args: args{
				s1: "абвзгде",
				s2: "абвжгде",
			},
			want: []Edit{
				{
					Type: Replace,
					Idx1: 3,
					Idx2: 3,
				},
			},
		},
		{
			name: "test2",
			args: args{
				s1: "абавзгд",
				s2: "абввжгде",
			},
			want: []Edit{
				{
					Type: Insert, // абавзгд -> абавзгд[е]
					Idx1: 6,
					Idx2: 7,
				},
				{
					Type: Replace, // абав[з]где -> абав[ж]где
					Idx1: 4,
					Idx2: 4,
				},
				{
					Type: Replace, // аб[а]вжгде -> аб[в]вжгде
					Idx1: 2,
					Idx2: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevenshteinString() = %v, want %v", got, tt.want)
			}
		})
	}
}
