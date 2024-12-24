package main

import "fmt" 

func add(x, y int) int {
	return x + y
}

func main() {	 

	fmt.Println("Start")
	data := add(5, 10)
	fmt.Println(data)
	defer fmt.Println("Middle")
	fmt.Println("End")
}

// Defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// The arguments are evaluated when the defer statement is evaluated, not when the function is called.
// Defer is commonly used to simplify functions that perform various clean-up actions.
// For example, if you open a file, you can defer a close on that file, so that it is closed as soon as the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// The arguments are evaluated when the defer statement is evaluated, not when the function is called.
// Defer is commonly used to simplify functions that perform various clean-up actions.
// For example, if you open a file, you can defer a close on that file, so that it is closed as soon as the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// The arguments are evaluated when the defer statement is evaluated, not when the function is called.
// Defer is commonly used to simplify functions that perform various clean-up actions.
// For example, if you open a file, you can defer a close on that file, so that it is closed as soon as the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// The arguments are evaluated when the defer statement is evaluated, not when the function is called.
// Defer is commonly used to simplify functions that perform various clean-up actions.
// For example, if you open a file, you can defer a close on that file, so that it is closed as soon as the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// The arguments are evaluated when the defer statement is evaluated, not when the function is called.
// Defer is commonly used to simplify functions that perform various clean-up actions.
// For example, if you open a file, you can defer a close on that file, so that it is closed as soon as the surrounding function returns.
//