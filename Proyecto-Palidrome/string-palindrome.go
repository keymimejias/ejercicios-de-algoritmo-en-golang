package main

import "fmt"

func main() {

	pali := "otto"

	for _, i := range pali {

		//fmt.Printf("%c\n", i)

		if i < 3 {

			if pali[i] == pali[i+3] {

				fmt.Println(true)

				fmt.Println("La palabra", pali, "si es palindrome")
				//break
			} else {
				fmt.Println(false)

				fmt.Println("La palabra", pali, "no es palindrome")
				//break
			}
		}
	}
}
