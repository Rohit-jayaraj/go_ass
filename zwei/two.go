package main
import "fmt"
func main(){
	var n1 int
	fmt.Print("Enter number of elements:")
	fmt.Scan(&n1)
	
	lst1 := make([]int,n1)
	sum1 := 0
	for i:=0;i<n1;i++{
		var inp int
		fmt.Scan(&inp)
		lst1[i] = inp
		sum1 = sum1 + inp
	}
	
	var n2 int
	fmt.Print("Enter number of elements in list2:")
	fmt.Scan(&n2)
	
	lst2 := make([]int,n2)
	sum2 := 0
	for i:=0;i<n2;i++{
		var x int
		fmt.Scan(&x)
		lst2[i] = x
		sum2 = sum2 + x
	}

	var length string
	if n1 == n2{
		length = "Yes"
	}else{
		length = "No"
	}
	
	var sum string
	if sum1 == sum2{
		sum = "Yes"
	}else{
		sum = "No"
	}
	
	fmt.Print("Are the lengths same? : ", length)
	fmt.Print("\nAre the sums same?: ", sum)
	
	fmt.Print("\nCommon elements:\n")
	for i:=0;i<n1;i++{
		for j:=0;j<n2;j++{
			if lst1[i] == lst2[j]{
				fmt.Print("\n")
				fmt.Print(lst1[i])
				}	
			}
	}
}