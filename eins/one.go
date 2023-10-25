package main 
import("bufio";"fmt";"os";"strings")
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s:= scanner.Text()
	fmt.Println(strings.Title(s))
}
