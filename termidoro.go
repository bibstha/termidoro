package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

func main() {
	minutes := userChoice()
	if minutes != -1 {
		showTimer(minutes)
		notify(minutes)
	}
}

func userChoice() int {
	a := `Select option:
1. 25 minutes break
2. 5 minutes break
3. 15 minutes break
m<number> for manual number of minutes
q to quit

Your choice? `
	fmt.Printf(a)

	var choice string
	fmt.Scanf("%s", &choice)
	minutes, err := parseInput(choice)

	for err != nil {
		fmt.Printf("Your choice? ")
		fmt.Scanf("%s", &choice)
		minutes, err = parseInput(choice)
	}

	return minutes
}

func parseInput(choice string) (int, error) {
	var validChoice = regexp.MustCompile(`^(1|2|3)$`)
	var validManualChoice = regexp.MustCompile(`m([0-9]+)`)

	var minutes string
	if validChoice.MatchString(choice) {
		switch choice {
		case "1":
			minutes = "25"
		case "2":
			minutes = "5"
		case "3":
			minutes = "15"
		}
	} else if validManualChoice.MatchString(choice) {
		minutes = validManualChoice.FindStringSubmatch(choice)[1]
	} else if choice == "q" {
		minutes = "-1"
	} else {
		return 0, errors.New("invalid argument")
	}
	i, _ := strconv.Atoi(minutes)
	return i, nil
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
	fmt.Println()
}

func notify(minutes int) {
	// OSX
	text := fmt.Sprintf("Completed timer for %d minute(s)", minutes)
	title := "Termidoro"
	cmd := "display notification \"" + text + "\" with title \"" + title + "\""
	fmt.Println("osascript " + " -e " + cmd)
	err := exec.Command("osascript", "-e", cmd).Run()
	if err != nil {
		log.Fatal(err)
	}
}
