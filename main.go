package main

import (
    "fmt"
    "os"
    "bufio"
)

 func main() {
    fmt.Print("Enter a grade: ")
    reader := bufio.NewReader(os.Stdin)
    input := reader.ReadString('\n')
    fmt.Println(input)
 }
