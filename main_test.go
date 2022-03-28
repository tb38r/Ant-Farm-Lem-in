package main

import (
	"testing"
	"fmt"
)

func BenchmarkAdd(b *testing.B){
    for i :=0; i < b.N ; i++{
        Add(4, 6)
    }
}

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func ExampleAdd() {
    fmt.Println(Add(4, 6))
    // Output: 10
}
// func Test_Add(t *testing.T){

//     got := Add(4, 6)
//     want := 10

//     if got != want {
//         t.Errorf("got %v, wanted %v", got, want)
//     }
// }
