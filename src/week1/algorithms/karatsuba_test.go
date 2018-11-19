package algorithms_test

import (
	"testing"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week1/algorithms"
)

func TestKaratsuba(t *testing.T) {
	// x, y, expected value
	cases := []struct {
		x string
		y string
		z string
	}{
		{"5", "6", "30"},
		{"2", "8", "16"},
		{"3", "-2", "-006"},
		{"-2", "3", "-006"},
		{"3141592653589793238462643383279502884197169399375105820974944592",
			"2718281828459045235360287471352662497757247093699959574966967627",
			"8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184"},
	}

	for i := 0; i < len(cases); i++ {
		testCase := cases[i]
		actual := algorithms.KaratsubaMult(testCase.x, testCase.y)
		if testCase.z != actual {
			t.Logf("KaratsubaMult failed to add %s * %s = %s, expected %s", testCase.x, testCase.y, actual, testCase.z)
			t.Fail()
		}
	}
}
