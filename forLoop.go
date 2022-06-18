package main

import "fmt"

func main() {
	//for i := 1; i <= 5; i++ {
	//	fmt.Println(i * i)
	//}

	i := 1
	for i <= 5 {
		fmt.Println(i * i)
		i += 1
	}

	sum := 0
	//for {
	//	sum++ //repeated forever
	//	fmt.Println(sum)
	//}
	fmt.Println(sum) //never reached

	for i := 1; i <= 5; i++ {
		if i == 3 {
			//break
			continue
		}
		fmt.Println(i)
	}
}

//for initialization; condition; post {

//}
