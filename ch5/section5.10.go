package main

//***************************
// import (
// 	"fmt"
// )

// func main () {
// 	defer func () {
// 		fmt.Println("Before call parse")
// 	}()

// 	fmt.Println(parse())
// 	panic(2)
// 	p := recover();
// 	if p != nil {
// 		fmt.Printf("internal error: %v", p)
// 	}

// 	defer func () {
// 		fmt.Println("After call parse")
// 	}()
// }

// func parse() (err error) {
// 	defer func () {
// 		p := recover();
// 		if p != nil {
// 			err = fmt.Errorf("internal error: %v", p)
// 		}

// 		fmt.Printf("p:%v\n", p)
// 	}()

// 	panic(8)
// 	return
// }

//***************************
// import (
// 	"time"
// 	"runtime"
// )

// func main () {
// 	defer func () {
// 		fmt.Println("Before call parse")
// 	}()

// 	go test()

// 	time.Sleep(2 * time.Second)
// }

// func test() {
// 	defer func () {
// 		fmt.Println("call test")
// 	}()

// 	runtime.Goexit()
// 	panic(2)
// }

//***************************
// import (
// 	"fmt"
// 	"time"
// )

// func main () {
// 	defer fmt.Println("in main: ", time.Now())

// 	defer func () {
// 		fmt.Println("defer outer ")
		
// 		defer func () {
// 			fmt.Println("defer inner ")
// 		}()
// 	}()
// 	if err := recover(); err != nil {
// 		fmt.Println(err)
// 	}

// 	panic("unknown error")
// }

//***************************
// import (
// 	"fmt"
// 	"time"
// )

// func main () {
// 	defer fmt.Println("in main: ", time.Now())

// 	defer func () {
// 		defer func () {
// 			panic("panic again and again")
// 		}()

// 		panic("panic again")
// 	}()

// 	panic("panic once")
// }

//***************************
// func main () {
// 	// test1()
// 	// test2()
// 	// test3()
// 	test4()
// }

// func test1() {
// 	defer recover()
// 	panic("1")
// }

// func test2() {
// 	defer func () {
// 		defer func () {
// 			recover()
// 		}()
// 	}() 

// 	panic("2")
// }

// func test3() {
// 	defer func () {
// 		recover()
// 	}()

// 	panic("3")
// }

// func test4() {
// 	defer func () {
// 		recover()
// 	}()

// 	defer func () {
// 		panic("4")
// 	}()
// }


//***************************
import (
	"fmt"
)

func main() {
	defer func() {
		e := recover()
		fmt.Println("f1", e)
	}()

	defer func() {
		fmt.Println("f2")
		panic("f2")
	}()

	panic("main")
}