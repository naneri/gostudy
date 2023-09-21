package search

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		needle int
		nums   []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "plain",
			args: struct {
				needle int
				nums   []int
			}{needle: 12, nums: []int{1, 2, 3, 4, 5, 12, 16}},
			want:  5,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := binarySearch(tt.args.needle, tt.args.nums)
			if got != tt.want {
				t.Errorf("binarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("binarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
