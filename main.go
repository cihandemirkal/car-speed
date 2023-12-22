package main

import (
	"fmt"
	"myapp/speed-package"
)

func main() {
	car := speed.NewCar(10, 2)
	track := speed.NewTrack(100)

	fmt.Println(speed.CanFinish(car, track))
}
