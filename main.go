package main

import (
	"log"
	"time"

	"github.com/alessiosavi/ProjectEulerGo/problem1"
	"github.com/alessiosavi/ProjectEulerGo/problem2"
	"github.com/alessiosavi/ProjectEulerGo/problem3"
	"github.com/alessiosavi/ProjectEulerGo/problem4"
	"github.com/alessiosavi/ProjectEulerGo/problem5"
	"github.com/alessiosavi/ProjectEulerGo/problem6"
	"github.com/alessiosavi/ProjectEulerGo/problem7"
	"github.com/alessiosavi/ProjectEulerGo/problem8"
	"github.com/alessiosavi/ProjectEulerGo/problem9"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	var fns []func()

	fns = append(fns, problem1.Win)
	fns = append(fns, problem2.Win)
	fns = append(fns, problem3.Win)
	fns = append(fns, problem4.Win)
	fns = append(fns, problem5.Win)
	fns = append(fns, problem6.Win)
	fns = append(fns, problem7.Win)
	fns = append(fns, problem8.Win)
	fns = append(fns, problem9.Win)

	for i, fn := range fns {
		start := time.Now()
		fn()
		duration := time.Since(start)
		log.Printf("Function [%d] duration: %v\n", i+1, duration)
	}

}
