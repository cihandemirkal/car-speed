package main

import "fmt"

// Define a Car struct with the following int type fields:
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// Define another struct type called Track with the field distance of type integer
type Track struct {
	distance int
}

/*
Allow creating a remote controlled car by defining a function NewCar
that takes the speed of the car in meters, and the battery drain percentage
as its two parameters (both of type int) and returns a Car instance:
*/

func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

/*
Allow creating a race track by defining a function NewTrack that takes
the track's distance in meters as its sole parameter (which is of type int):
*/

func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

/*
3. Drive the car
Implement the Drive function that updates the number of meters driven based
on the car's speed, and reduces the battery according to the battery drainage.
If there is not enough battery to drive one more time the car will not move:
*/

func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
		return car
	}
	return car
}

/*
To finish a race, a car has to be able to drive the race's distance.
This means not draining its battery before having crossed the finish line.
Implement the CanFinish function that takes a Car and a Track instance as
its parameter and returns true if the car can finish the race; otherwise,
return false.
*/

func CanFinish(car Car, track Track) bool {
	return (car.battery/car.batteryDrain)*car.speed >= track.distance
}

func main() {
	car := NewCar(10, 2)
	track := NewTrack(100)

	fmt.Println(CanFinish(car, track))
}
