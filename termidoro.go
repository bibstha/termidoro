package main

import (
	"fmt"
	"time"
)

func main() {
	choice := user_choice()
	fmt.Println("You chose: " + choice)
	showTimer(1)
}

func user_choice() string {
	a := `Select option:
1. 25 minutes break
2. 5 minutes break
3. 15 minutes break
4. m<number> for manual number of minutes
5. q to quit`
	fmt.Println(a)

	var choice string
	fmt.Scanf("%s", &choice)
	return choice
}

func showTimer(minutes int) {
	totalSeconds := minutes * 60
	elapsedSeconds := 0

	fmt.Printf("Running for %d minute(s)\n", minutes)

	ticker := time.Tick(time.Second)
	for range ticker {
		minutes := elapsedSeconds / 60
		seconds := elapsedSeconds % 60

		fmt.Printf("\r %0#2d:%0#2d", minutes, seconds)
		if elapsedSeconds == totalSeconds {
			break
		}

		elapsedSeconds++
	}
}
