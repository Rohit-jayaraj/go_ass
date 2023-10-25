package main
import "fmt"
func main(){
	x,y := check()
	fmt.Print("Is even:",y)
	fmt.Print("\nHalf:",x)
}

func check() (int,bool){
	var num int
	fmt.Scan(&num)
	var test bool
	var half int
	if num%2 == 0{
		test = true
	}else{
		test = false
	}
	half = num/2
	return half, test
}
	
	