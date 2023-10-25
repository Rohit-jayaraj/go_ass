package main
import ("fmt";"math")


func main() {
	var start, end int
	fmt.Print("Enter the starting number: ")
	_, err := fmt.Scanf("%d", &start)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	fmt.Print("Enter the ending number: ")
	_, err = fmt.Scanf("%d", &end)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	if start < 2 {
		start = 2
	}
	fmt.Printf("Prime numbers between %d and %d:\n", start, end)
	for num := start; num <= end; num++ {
		if isPrime(num) {
			fmt.Println(num)
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n)); i += 2) {
		if n%i == 0 {
			return false
		}
	}
	return true
}