package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var random *rand.Rand

var inputRoad1Pool []int
var inputRoad2Pool []int
var inputRoad3Pool []int
var inputRoad4Pool []int

var ouputRoad1ch chan int
var ouputRoad2ch chan int
var ouputRoad3ch chan int
var ouputRoad4ch chan int

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s)
}

func randomRange(min, max int) int {
	return random.Intn((max-min)+1) + min
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Starting a party...")

	var circle = make(chan int, 8)
	ouputRoad1ch = make(chan int)
	ouputRoad2ch = make(chan int)
	ouputRoad3ch = make(chan int)
	ouputRoad4ch = make(chan int)

	wg.Add(9)
	go circlePool(circle, &wg)
	go inputRoad1(circle, &wg)
	go inputRoad2(circle, &wg)
	go inputRoad3(circle, &wg)
	go inputRoad4(circle, &wg)
	go outputRoad1(&wg)
	go outputRoad2(&wg)
	go outputRoad3(&wg)
	go outputRoad4(&wg)

	wg.Wait()
	fmt.Println("Main goroutine exits")

}

func circlePool(circle chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for car := range circle {
		carDistribution(car)
	}
}

func DrivingForNSecs(nSecs int) {
	time.Sleep(time.Duration(nSecs) * time.Second)
}

func carDistribution(car int) {
	switch car {
	// cases represents the exit number
	case 1:
		{
			DrivingForNSecs(1)
			ouputRoad1ch <- car
		}
	case 2:
		{
			DrivingForNSecs(2)
			ouputRoad2ch <- car
		}
	case 3:
		{
			DrivingForNSecs(3)
			ouputRoad3ch <- car
		}
	case 4:
		{
			DrivingForNSecs(4)
			ouputRoad4ch <- car
		}
	}
}

func inputRoad1(circle chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	processInputRoad(circle, inputRoad1Pool, (time.Duration(random.Intn(5)) * time.Second))
}

func inputRoad2(circle chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	processInputRoad(circle, inputRoad2Pool, 1*time.Second)
}

func inputRoad3(circle chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	processInputRoad(circle, inputRoad3Pool, 2*time.Second)
}

func inputRoad4(circle chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	processInputRoad(circle, inputRoad4Pool, 100*time.Millisecond)
}

func outputRoad1(wg *sync.WaitGroup) {
	defer wg.Done()
	processOutputRoad(ouputRoad1ch, time.Duration(random.Intn(5))*time.Second)
}

func outputRoad2(wg *sync.WaitGroup) {
	defer wg.Done()
	processOutputRoad(ouputRoad2ch, 1*time.Second)
}

func outputRoad3(wg *sync.WaitGroup) {
	defer wg.Done()
	processOutputRoad(ouputRoad3ch, 1*time.Hour)
}

func outputRoad4(wg *sync.WaitGroup) {
	defer wg.Done()
	processOutputRoad(ouputRoad4ch, 100*time.Millisecond)
}

// =======PROCESSING FUNCTIONS=============

func processOutputRoad(outputRoad chan int, d time.Duration) {
	for car := range outputRoad {
		fmt.Println(fmt.Sprintf("Car drove out [%d] exit", car))
		time.Sleep(d)
	}
}

func processInputRoad(circle chan int, inputRoadPool []int, d time.Duration) {
	for {
		//random in range to determine what exit to drive
		inputRoadPool = append(inputRoadPool, randomRange(1, 4))
		//We send a car to circle if possible, otherwise skip and generate another set of cars for the next second
		select {
		case circle <- inputRoadPool[0]:
			fmt.Println(fmt.Sprintf("Car drove into circle. To drive out on %d exit", inputRoadPool[0]))
			inputRoadPool = inputRoadPool[1:]
		default:
		}
		time.Sleep(d)
	}
}
