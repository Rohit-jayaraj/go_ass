package main
import "fmt"
func main() {
	fmt.Println("Enter the Unit consumed")
	var UC int
	fmt.Scan(&UC)
	var total int = 0
	fmt.Println("Consumer Number CDV0025")
	fmt.Println("Units Consumed", UC)
	if UC <= 200 {
		fmt.Printf("Split bill: %v x 0 \n", UC)
		fmt.Println("Total bill :", total)
	} else if UC >= 201 && UC <= 300 {
		total = (200 * 0) + (UC-200)*10
		fmt.Printf("Split bill: 200 x 0 + %v x 10 \n", (UC - 100))
		fmt.Println("Total bill:", total)
	} else if UC >= 301 && UC <= 500 {
		total = (200 * 0) + (100)*10 + (UC-300)*15
		fmt.Printf("Split bill : 200 x 0 + 200 x 10 + %v x 15", (UC - 300))
		fmt.Println("\nTotal bill:", total)
	} else {
		total = (200 * 0) + (100)*10 + (200)*15 + (UC-500)*25
		fmt.Printf("Split bill: 200 x 0 + 200 x 10 + 200 x 15 + %v x 25\n", (UC - 500))
		fmt.Println("Total bill: ", total)
	}
}
