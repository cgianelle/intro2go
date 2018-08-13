package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const maxNumGuesses int = 5

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var (
		guess string
		n     int
		err   error
	)
	found := false
	numberToGuess := r1.Intn(100)

	fmt.Println("Pick a number between 0 and 100")
	fmt.Println("You get a maximum of", maxNumGuesses, "guesses")

	for i := 1; i <= maxNumGuesses; i++ {
		fmt.Print("Guess #", i, " : ")
		n, err = fmt.Scan(&guess)
		if err == nil && n > 0 {
			g, e := strconv.Atoi(guess)
			if e != nil {
				fmt.Println("Huh, I do not recognize what you typed, is it a number? Goodbye")
				panic(e)
			}
			if g > numberToGuess {
				fmt.Println("You're guess was too high!")
			} else if g < numberToGuess {
				fmt.Println("You're guess was too low!")
			} else {
				if i == maxNumGuesses {
					fmt.Println("Well, somebody got lucky did't they!!! Guessed the correct number with no chances left! ")
				} else if i == (maxNumGuesses - 1) {
					fmt.Println("Cutting it close there aren't you! Only one chance remained! But you got it!")
				} else {
					fmt.Println("TIME TO CELEBRATE!!!! AWESOME JOB!!!! You correctly guessed the number.")
				}
				fmt.Println("Total Number of Guesses:", i)
				found = true
				break
			}
		} else {
			fmt.Println("Huh, I do not recognize what you typed, goodbye")
			break
		}
	}

	if !found {
		fmt.Println("Oh too bad! You're out of guesses.")
		fmt.Println("Better luck next time!!!")
		fmt.Println("The correct number was ", numberToGuess)
	}

}
