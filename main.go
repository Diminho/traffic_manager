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
	for {
		cars := random.Intn(5)
		for i := 0; i <= cars; i++ {
			inputRoad1Pool = append(inputRoad1Pool, randomRange(1, 4))
		}

		//We send a car to circle if possible, otherwise skip and generate another set of cars for the next second
		select {
		case circle <- inputRoad1Pool[0]:
			fmt.Println(fmt.Sprintf("Car drove into circle. To drive out on %d exit", inputRoad1Pool[0]))
			inputRoad1Pool = inputRoad1Pool[1:]
		default:
		}

		time.Sleep(1 * time.Second)
	}
}

func inputRoad2(circle chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		inputRoad2Pool = append(inputRoad2Pool, randomRange(1, 4))
		//We send a car to circle if possible, otherwise skip and generate another set of cars for the next second
		select {
		case circle <- inputRoad2Pool[0]:
			fmt.Println(fmt.Sprintf("Car drove into circle. To drive out on %d exit", inputRoad2Pool[0]))
			inputRoad2Pool = inputRoad2Pool[1:]
		default:
		}
		time.Sleep(1 * time.Second)
	}
}

func inputRoad3(circle chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		inputRoad3Pool = append(inputRoad3Pool, randomRange(1, 4))
		//We send a car to circle if possible, otherwise skip and generate another set of cars for the next second
		select {
		case circle <- inputRoad3Pool[0]:
			fmt.Println(fmt.Sprintf("Car drove into circle. To drive out on %d exit", inputRoad3Pool[0]))
			inputRoad3Pool = inputRoad3Pool[1:]
		default:
		}
		time.Sleep(2 * time.Second)
	}
}

func inputRoad4(circle chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		inputRoad4Pool = append(inputRoad4Pool, randomRange(1, 4))
		//We send a car to circle if possible, otherwise skip and generate another set of cars for the next second
		select {
		case circle <- inputRoad4Pool[0]:
			fmt.Println(fmt.Sprintf("Car drove into circle. To drive out on %d exit", inputRoad4Pool[0]))
			inputRoad4Pool = inputRoad4Pool[1:]
		default:
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func outputRoad1(wg *sync.WaitGroup) {
	defer wg.Done()
	for car := range ouputRoad1ch {
		fmt.Println(fmt.Sprintf("Car [%d] drove ouf first exit", car))
		time.Sleep(time.Duration(random.Intn(5)) * time.Second)
	}
}

func outputRoad2(wg *sync.WaitGroup) {
	defer wg.Done()
	for car := range ouputRoad2ch {
		fmt.Println(fmt.Sprintf("Car [%d] drove ouf second exit", car))
		time.Sleep(1 * time.Second)
	}
}

func outputRoad3(wg *sync.WaitGroup) {
	defer wg.Done()
	for car := range ouputRoad3ch {
		fmt.Println(fmt.Sprintf("Car [%d] drove ouf third exit", car))
		time.Sleep(1 * time.Hour)
	}
}

func outputRoad4(wg *sync.WaitGroup) {
	defer wg.Done()
	for car := range ouputRoad4ch {
		fmt.Println(fmt.Sprintf("Car [%d] drove ouf fourth exit", car))
		time.Sleep(100 * time.Millisecond)
	}
}
