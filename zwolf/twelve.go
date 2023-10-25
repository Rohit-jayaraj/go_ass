package main
import ("fmt";"strings")
func main() {
    var inp string
    fmt.Scan(&inp)
    vowelCounts := make(map[rune]int)
    vowels := "aeiouAEIOU"
    inp= strings.ToLower(inp)
    for _, char := range inp{
        if strings.ContainsRune(vowels, char) {
            vowelCounts[char]++
        }
    }
    fmt.Println("Vowel Counts:")
    for vowel, count := range vowelCounts {
        fmt.Printf("%c: %d\n", vowel, count)
    }
}
