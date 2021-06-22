package main

import "fmt"

type people struct {
    name string
    age int
}

type MyInt1 int
type MyInt2 = int 	// 定义别名

func main() {
	var xiaoming *people;
	xiaoming = &people{"xiaoming",22};
	fmt.Println((*xiaoming).name);

	var i int = 0
	// var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i2)

	a := []int{7, 8, 9}
		fmt.Printf("%+v\n", a)
		ap(a)
		fmt.Printf("%+v\n", a)
		app(a)
		fmt.Printf("%+v\n", a)
}

func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}