package main
import ("bufio";"fmt";"os";"strings")
func main() {
    wordCounts := make(map[string]int)
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a sentence: ")
    inputSentence, _ := reader.ReadString('\n')
    words := strings.Fields(inputSentence)
    for _, word := range words {
        word = strings.ToLower(word)
        wordCounts[word]++
    }
    fmt.Println("Word Counts:")
    for word, count := range wordCounts {
        fmt.Printf("%s: %d\n", word, count)
    }
}
