package channels

import (
	"reflect"
	"sort"
	"testing"
)

func TestFanOut(t *testing.T) {
	type args struct {
		in      []int
		workers int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test fanout - fanin",
			args: args{
				in:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				workers: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan int)
			go func() {
				for _, v := range tt.args.in {
					in <- v
				}
				close(in)
			}()

			ch := FanOut(in, tt.args.workers)
			gotCh := FanIn(ch...)
			got := []int{}
			for v := range gotCh {
				got = append(got, v)
			}
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.args.in) {
				t.Errorf("FanIn() = %v, want %v", got, tt.args.in)
			}
		})
	}
}
