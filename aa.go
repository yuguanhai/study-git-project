package study_git_project

import "fmt"

type Alpha struct {
	Name string
	Age  int
}

func (*Alpha) Add(a, b int) int {
	return a + b
}

type Beta struct {
	Alpha
	c string
}

func main() {
	b := Beta{
		Alpha: Alpha{
			Name: "test",
			Age:  1,
		},
		c: "test",
	}
	fmt.Println(b.Add(1, 2))
}
