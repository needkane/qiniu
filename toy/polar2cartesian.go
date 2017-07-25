package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

type polar struct {
	radius float64
	degree float64
}
type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and an angle(in degree),e.g ., 12.5 90 , " + "or %s to quit."

//before main auto execute but it can't showing call
func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "ctrl+Z,enter")
	} else {
		prompt = fmt.Sprintf(prompt, "ctrl+D")
	}
}
func main() {
	//channel simple demo
	/*message := make(chan string, 10)
	//receive
	message <- "kane"
	message <- "crystal"
	//transmission,because block to obtain one data
	message1 := <-message
	message2 := <-message
	fmt.Println(message1)
	fmt.Println(message2)*/

	questions := make(chan polar)
	//when the channel created ,it will close
	defer close(questions)

	//receive message
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}
func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			//block until receive one data
			polarCoord := <-questions
			degree := polarCoord.degree * math.Pi / 180.0 //度变弧度
			x := polarCoord.radius * math.Cos(degree)
			y := polarCoord.radius * math.Sin(degree)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

const result = "Polar radius=%.02f degree=%.02f -> cartesian x=%.02f y=%.02f\n"

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Println("radius and angle:")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, degree float64

		if _, err := fmt.Sscanf(line, "%f %f", &radius, &degree); err != nil {
			fmt.Println(os.Stderr, "invalid input")
			continue
		}

		questions <- polar{radius, degree}
		coord := <-answers
		fmt.Printf(result, radius, degree, coord.x, coord.y)
	}
	fmt.Println()
}
