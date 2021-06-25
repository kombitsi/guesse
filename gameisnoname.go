package main

import (
	"bufio"
	_ "bufio"
	"fmt"
	_ "fmt"
	"log"
	_ "log"
	"math/rand"
	_ "math/rand"
	"os"
	_ "os"
	"strconv"
	_ "strconv"
	"strings"
	_ "strings"
	"time"
	_ "time"
)


func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 5
	fmt.Println("Game is not name. Угадай число между 1 и 100.")
//	fmt.Println("Can you guess it?")
	fmt.Println(target)

		reader := bufio.NewReader (os.Stdin) // блок для ввода чисел и проверки
		success := false
		for guesses := 0; guesses < 10; guesses++ {
			fmt.Println("У тебя есть", 10-guesses , "попыток.")

			fmt.Print("Введите число:")
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input = strings.TrimSpace(input)
				guess, err := strconv.Atoi(input)
			if err != nil {
				log.Fatal(err)

			}
			if guess < target {
				fmt.Println("Oops. Ты ввел маленькое число.")
			} else if guess > target {
				fmt.Println("Oops. Ты ввел большое число.")
			} else {
				success = true
				fmt.Println("Ура! Ты угадал!")
				break
		}
	}
	if !success {
		fmt.Println("Сорян, ты лузер и ты не угадал. :", target)
	}
}
