package C298

import "testing"

//func Test_greatestLetter(t *testing.T) {
//	type args struct {
//		s string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := greatestLetter(tt.args.s); got != tt.want {
//				t.Errorf("greatestLetter() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestGreatestLetter(t *testing.T) {
	t.Log(greatestLetter("AbCdEfGhIjK"))
}
