package main
import "fmt"
func main(){
	var hobbies [3]string = [3]string{"Playing Cricket","Playing Piano", "Reading Book"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	newList := hobbies[1:3]
	fmt.Println(newList)

	nList := hobbies[:2]
	fmt.Println(nList)
}