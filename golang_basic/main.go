/*
	配列
*/

package golang_basic

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.X)
}
