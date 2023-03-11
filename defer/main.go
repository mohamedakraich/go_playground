package main

import "fmt"

// The behavior of defer statements is straightforward and predictable. There are three simple rules:

// 1. A deferred function’s arguments are evaluated when the defer statement is evaluated.
func a() {
	i := 0
	defer fmt.Println(i)
	i++
}

// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func b() {
	defer fmt.Println()
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// 3. Deferred functions may read and assign to the returning function’s named return values.
func c() (i int) {
	defer func() {
		i++
		fmt.Println(i)
	}()
	return 9
}

func main() {
	a()
	b()
	c()
	fmt.Println("https://go.dev/blog/defer-panic-and-recover")
	fmt.Println("mechanics of panic and defer")
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// https://pkg.go.dev/encoding/json
