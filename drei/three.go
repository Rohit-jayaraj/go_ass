package main
import "fmt"
func main() {
	fmt.Println("Enter elements of the list")
	var inp []int
	for{
		var data int
		fmt.Scan(&data)
		if data == 0 {
			break
		}
		inp = append(inp, data)
	}
	var square []int
	for _, i := range inp{
		square = append(square, i*i)
	}
	var positive []int
	for _, i := range inp{
		if i > 0 {
			positive = append(positive, i)
		}
	}
	fmt.Print("Input list", inp)
	fmt.Println()
	fmt.Print("Square list", square)
	fmt.Println()
	fmt.Print("positive list", positive)
}
