package C298

import "testing"

//func test_getlastnum(t *testing.t) {
//	type args struct {
//		num int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// todo: add test cases.
//	}
//	for _, tt := range tests {
//		t.run(tt.name, func(t *testing.t) {
//			if got := getlastnum(tt.args.num); got != tt.want {
//				t.errorf("getlastnum() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func test_minimumnumbers(t *testing.t) {
//	type args struct {
//		num int
//		k   int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// todo: add test cases.
//	}
//	for _, tt := range tests {
//		t.run(tt.name, func(t *testing.t) {
//			if got := minimumnumbers(tt.args.num, tt.args.k); got != tt.want {
//				t.errorf("minimumnumbers() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestMinimumNumbers(t *testing.T) {
	t.Log(minimumNumbers(5, 1))
}
