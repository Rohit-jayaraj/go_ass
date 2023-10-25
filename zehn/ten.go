package main
import ("fmt")
func main(){
	greatest := findGreatestNumber(10,6,23,55,6,12)
	fmt.Printf("The greatest number is: %d\n", greatest)
}

func findGreatestNumber(numbers ...int) int{
	if len(numbers) == 0 {
		return 0
	}
	greatest := numbers[0]
	for _, num := range numbers {
		if num > greatest {
			greatest = num
		}
	}

	return greatest
}