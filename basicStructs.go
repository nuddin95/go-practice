package main

import "fmt"

type car struct {
	gasPedal uint16
	brakePedal uint16
	steeringWheel int16
	topSpeeKmh float64
}

func main(){
	var sampleCar car = car{gasPedal: 1000,
						brakePedal:200,
						steeringWheel: 12,
						topSpeeKmh: 12345.43}

	fmt.Println("SHOULD BE 1000: ", sampleCar.gasPedal)


}