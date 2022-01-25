package IfElseSwitchCaseAndSelect

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func PlayRockPaperScisors() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(3)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors -> ")
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose ROCK")
		break
	case PAPER:
		fmt.Println("Computer chose PAPER")
		break
	case SCISSORS:
		fmt.Println("Computer chose SCISSORS")
		break
	default:
	}

	fmt.Println()

	if playerValue == computerValue {
		fmt.Println("It's a draw")
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				fmt.Println("Computer wins!")
			} else {
				fmt.Println("Player wins!")
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				fmt.Println("Computer wins!")
			} else {
				fmt.Println("Player wins!")
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				fmt.Println("Computer wins!")
			} else {
				fmt.Println("Player wins!")
			}
			break
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}


var chan1 = make(chan string)
var chan2 = make(chan string)

func task1() {
	time.Sleep(1 * time.Second)
	chan1 <- "one"
}

func task2() {
	time.Sleep(2 * time.Second)
	chan2 <- "two"
}

func RunSelect() {
	go task1()
	go task2()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- chan1:
			fmt.Println("received", msg1)
		case msg2 := <- chan2:
			fmt.Println("received", msg2)
		}
	}
}