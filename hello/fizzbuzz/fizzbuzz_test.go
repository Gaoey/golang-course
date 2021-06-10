package fizzbuzz

import "testing"

func Test_GivenOne_WantOne(t *testing.T) {
	given := 1
	want := "1"

	get := Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

type Testcase struct {
	given int
	want  string
}

func Test_Fizzbuzz(t *testing.T) {
	testcases := []Testcase{
		{
			given: 1,
			want:  "1",
		},
		{
			given: 3,
			want:  "Fizz",
		},
		{
			given: 5,
			want:  "Buzz",
		},
	}

	for _, v := range testcases {
		result := FizzBuzz(v.given)
		if v.want != result {
			t.Errorf("given %d want %q but get %q\n", v.given, v.want, result)
		}
	}
}
