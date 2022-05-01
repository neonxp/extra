package channels

import (
	"reflect"
	"sort"
	"testing"
)

func TestFanIn(t *testing.T) {
	type args struct {
		chans [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test fanin",
			args: args{
				chans: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			chans := make([]chan int, len(tt.args.chans))
			for i, v := range tt.args.chans {
				func(i int, v []int) {
					chans[i] = make(chan int, len(v))
					go func(i int, v []int) {
						for _, v2 := range v {
							chans[i] <- v2
						}
						close(chans[i])
					}(i, v)
				}(i, v)
			}

			ch := FanIn(chans...)

			for o := range ch {
				got = append(got, o)
			}
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FanIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
