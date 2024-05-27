package main

import "fmt"

func main(){
	userNames := make([]string,2, 5) // 5 is maximum capacity

	userNames = append(userNames,"Aditya")
	userNames = append(userNames,"Jimit")

	fmt.Println(userNames)
}