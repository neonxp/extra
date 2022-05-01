package levenshtein

import (
	"reflect"
	"testing"
)

func TestLevenshtein2(t *testing.T) {
	type args struct {
		s1   []myType
		s2   []myType
		cost EditCost[myType]
	}
	tests := []struct {
		name string
		args args
		want []Edit
	}{
		{
			name: "test1",
			args: args{
				s1: []myType{
					{
						a: "a",
					}, {
						a: "b",
					}, {
						a: "c",
					},
				},
				s2: []myType{
					{
						a: "a",
					}, {
						a: "d",
					}, {
						a: "c",
					},
				},
				cost: EditCost[myType](func(t EditType, from, to *myType) int {
					return 1
				}),
			},
			want: []Edit{
				{
					Type: Replace,
					Idx1: 1,
					Idx2: 1,
				},
			},
		},
		{
			name: "test2 - costly replace",
			args: args{
				s1: []myType{
					{
						a: "a",
					}, {
						a: "b",
					}, {
						a: "c",
					},
				},
				s2: []myType{
					{
						a: "a",
					}, {
						a: "d",
					}, {
						a: "c",
					},
				},
				cost: EditCost[myType](func(t EditType, from, to *myType) int {
					if t == Replace {
						return 100
					}
					return 1
				}),
			},
			want: []Edit{
				{
					Type: Delete,
					Idx1: 2,
					Idx2: 0,
				},
				{
					Type: Insert,
					Idx1: 1,
					Idx2: 2,
				},
				{
					Type: Delete,
					Idx1: 1,
					Idx2: 0,
				},
				{
					Type: Insert,
					Idx1: 0,
					Idx2: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Levenshtein(tt.args.s1, tt.args.s2, tt.args.cost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Levenshtein() = %v, want %v", got, tt.want)
			}
		})
	}
}

type myType struct {
	a string
}
