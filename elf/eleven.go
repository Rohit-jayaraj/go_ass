package main
import ("fmt")
func main(){
    var n int
    fmt.Print("Enter the number of Fibonacci numbers to generate: ")
    fmt.Scanf("%d", &n)
    if n <= 0 {
        fmt.Println("Please enter a positive number.")
        return
    }
    fmt.Println("Fibonacci Sequence:")
    for i := 0; i < n; i++ {
        fmt.Printf("%d\n", fibonacci(i))
    }
}

func fibonacci(n int) int{
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
