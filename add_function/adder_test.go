package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
 sum := Add(2, 2)
 expected := 4

 if sum != expected {
	 t.Errorf("sum %d but expected %d", sum, expected)
 }
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}