package collections

import (
	"reflect"
	"testing"
)

func TestMergeScalar(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				s1: []int{1, 2, 4, 5, 6},
				s2: []int{3, 3},
			},
			want: []int{1, 2, 3, 3, 4, 5, 6},
		},
		{
			name: "test 2",
			args: args{
				s1: []int{1, 3, 5, 7},
				s2: []int{0, 2, 4, 6, 8},
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "test 3",
			args: args{
				s1: []int{8, 6, 4, 2, 0},
				s2: []int{1, 3, 5, 7},
			},
			want: []int{1, 3, 5, 7, 8, 6, 4, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeScalar(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		s1 []*myType
		s2 []*myType
	}
	tests := []struct {
		name string
		args args
		want []*myType
	}{
		{
			name: "test1",
			args: args{
				s1: []*myType{
					{400}, {200}, {100},
				},
				s2: []*myType{
					{300},
				},
			},
			want: []*myType{
				{400}, {300}, {200}, {100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

type myType struct {
	Weight int
}

func (t *myType) Less(t2 *myType) bool {
	return t2.Weight < t.Weight
}
