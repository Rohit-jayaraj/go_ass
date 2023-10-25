package main
import "fmt"
func main() {
	data := map[int]int{1001: 21, 1002: 35, 1003: 12, 1004: 64, 1005: 17, 1006: 59, 1007: 43}
	young := []int{}
	senior:= []int{}
	for id, age := range data {
		if age < 18 {
			young = append(young, id)
		} else if age > 60 {
			senior = append(senior, id)
		}
	}
	fmt.Println("Young:", young)
	fmt.Println("\nSenior Citizen:", senior)
}
