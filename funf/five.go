package main
import "fmt"
func main() {
	data := map[string]map[string]float32{"cse1": {"cse0001": 8.9, "cse003": 7.8},"cse2": {"cse0010": 8.0, "cse0011": 7.0}}
	fmt.Println("Enter roll number of student to search")
	var roll string
	fmt.Scan(&roll)
	section := ""
	for key, value := range data {
	_, rollexist := value[roll]
	if rollexist {
		section = key
		break
		}
	}
	if section == "" {
		cgpa := data[section][roll]
		fmt.Println(section)
		fmt.Println(cgpa)
	} else {
		fmt.Println("No such student exists")
	}
}