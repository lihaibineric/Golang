package main

import (
	"fmt"
)

// func  main(){
// 	c := make(chan int)
// 	i:=1

// 	go func(){
// 		c <- i
// 		fmt.Println(i)
// 	}()

// 	for{
// 		num := <- c 
// 		fmt.Println(string(96+num))
// 		i++
// 		if i>26{
// 			break
// 		}
// 	}
// }


// func main() {
// 	c := make(chan int)
// 	i := 1

// 	go func() {
// 		for {
// 			c <- i
// 			fmt.Println(i)
// 			i++
// 			if i > 26 {
// 				close(c)
// 				break
// 			}
// 		}
// 	}()

// 	go func(){
// 		for num := range c {
// 			fmt.Println(string(96 + num))
// 		}
// 	}()

// }


// func main() {
//    c := make(chan int)
//    go func() {
//       for i := 1; i <= 26; i++ {
// 		c <- 0
//         fmt.Println(i)
// 		<-c
//       }
//    }()
// 	for i := 1; i <= 26; i++ {
//         <-c
//         fmt.Println(string(96+i))
// 		c<-0
//     }
// }


// func main() {
// 	exit := make(chan bool)
// 	c := make(chan int)

// 	go func() {
// 		for i := 1; i <= 26; i++ {
// 			<-c
// 			fmt.Println(string(96 + i))
// 			c <- 0
// 		}
// 		exit <- true
// 	}()

// 	go func() {
// 		for i := 1; i <= 26; i++ {
// 			fmt.Println(i)
// 			c <- 0
// 			<-c
// 		}
// 	}()

// 	<-exit
// }


func main() {
	c := make(chan int)

	go func() {
		for i := 1; i <= 26; i++ {
			<-c
			fmt.Println(string(96 + i))
			c <- 0
		}
	}()

	for i := 1; i <= 26; i++ {
			fmt.Println(i)
			c <- 0
			<-c
	}
}

