package main
import "fmt"
func main() {
	person1 := Bank_account{
		account_no: 0025,
		name: "Rohit",
		balance_amount: 0,
		amount_deposited: 1310,
	}
	person2 := Bank_account{
		account_no: 0501,
		name: "Vladimir Putin",
		balance_amount: 500,
		amount_deposited: 700,
	}
	printUpdated(&person1)
	println()
	printUpdated(&person2)
}


func printUpdated(user *Bank_account) {
	fmt.Println("Details")
	user.balance_amount += user.amount_deposited
	fmt.Println("Acc ID", user.account_no)
	fmt.Println("Name", user.name)
	fmt.Println("Current Balance", user.balance_amount)
}

type Bank_account struct {
	account_no int
	name string
	balance_amount int
	amount_deposited int
}
