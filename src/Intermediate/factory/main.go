package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

type Plane struct {
	val string
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire_C, "
			outCarChan <- car

		case plane := <-planeChan:
			plane.val += "Tire_P, "
			outPlaneChan <- plane
		}

	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += " Engine_P, "
			outPlaneChan <- plane
		}
	}
}

func StartCarWork(carChan chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		carChan <- Car{val: "Car" + strconv.Itoa(i) + " :"}
		i++
	}

}

func StartPlaneWork(planeChan chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		planeChan <- Plane{val: "Plane" + strconv.Itoa(i) + " :"}
		i++
	}

}

func main() {
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)
	go StartCarWork(carChan1)
	go StartPlaneWork(planeChan1)
	go MakeTire(carChan1, planeChan1, carChan2, planeChan2)
	go MakeEngine(carChan2, planeChan2, carChan3, planeChan3)

	for {
		select {
		case resultCar := <-carChan3:
			fmt.Println(resultCar)
		case resultPlane := <-planeChan3:
			fmt.Println(resultPlane)
		}
	}
}
