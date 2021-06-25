package main

import (
	_ "bufio"
	"fmt"
	_ "fmt"
	_ "log"
	"math/rand"
	_ "math/rand"
	_ "os"
	_ "strconv"
	_ "strings"
	"time"
	_ "time"
)


func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 5
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target) }

	reader := bufio.NewReader (os.Stdin) // блок для ввода чисел и проверки

		fmt.Print("Введите число:")
		input, err := reader.ReadString('/n')
		if err != nul {
			log.Fatal(err)
		}
		input = strings.TrimSpase(input) {
			guess, err := strconv.Atoi(input)
			if err != nil {
			log.Fatal(err)
			}
		}
		for guesses := 0; guesses < 10; guesses++ {
			if guess < target {
				fmt.Println("Oops. Ты ввел маленькое число.")
			} else if guess > target {
				fmt.Println("Oops. Ты ввел большое число.")
			} else {
				fmt.Println("Good job! Ты угадал!")
				break
		}
	}
}
