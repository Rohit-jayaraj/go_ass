package main
import "fmt"
func main() {
	temp := map[string]int{"sun": 32, "mon": 30, "tue": 29, "wed": 25, "thur": 25, "fri": 27, "sat": 24}
	len := len(temp)
	total := 0
	max := 0
	for _, i := range temp {
		total += i
	}
	maxday := ""
	for day, i := range temp {
		if max < i {
			max = i
			maxday = day
		}
	}
	fmt.Printf("day recorded maximum temperature %v : %v \n", maxday, max)
	fmt.Println("recorded average temperature of week ", total/len)
}
