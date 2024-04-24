package cmd

import (
	"bufio"
	"fmt"
	"martian-robots/internal/mapMars"
	"martian-robots/internal/parsing"
	"martian-robots/internal/robot"
	"os"
)

func Execute() error {
	var mars mapMars.MarsField
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Size of Mars: ")
	marsSize, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input(incorrect input):", err)
		return err
	}

	mars.XMarsSize, mars.YMarsSize, err = parsing.ParseMap(marsSize)
	if err != nil {
		fmt.Println("Error reading input(incorrect input):", err)
		return err
	}

	robotCount := 1
	for {
		var robotDetails, path string
		newRobot := robot.NewRobot(&mars)
		fmt.Printf("Robot initial coordinates and direction #%d: ", robotCount)
		robotDetails, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input(incorrect input):", err)
			return err
		}

		newRobot.X, newRobot.Y, newRobot.Direction, err = parsing.ParseRobot(robotDetails)
		if err != nil {
			fmt.Println("Error reading input(incorrect input):", err)
			return err
		}

		fmt.Printf("Robot route №%d: ", robotCount)
		path, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input(incorrect input):", err)
			return err
		}
		newRobot.WalkPath = []rune(path)

		newRobot.RobotWalk(0)

		if newRobot.Lost {
			fmt.Printf("The result coordinates are: %d %d %c %s\n", newRobot.X, newRobot.Y, parsing.AngleMap[newRobot.Direction].(rune), "LOST")
		} else {
			fmt.Printf("The result coordinates are: %d %d %c\n", newRobot.X, newRobot.Y, parsing.AngleMap[newRobot.Direction].(rune))
		}

		mars = newRobot.Mars
		robotCount++

		var continueInput string
		fmt.Print("Do you want to continue adding robots? (yes/no): ")
		_, err = fmt.Scan(&continueInput)
		if err != nil {
			return err
		}
		if continueInput != "yes" {
			break
		}
	}
	return nil
}
