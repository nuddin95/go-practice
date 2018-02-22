package main

import "fmt"

type car struct {
	gasPedal uint16
	brakePedal uint16
	steeringWheel int16
	topSpeeKmh float64
}

func (c car) totalPedal() uint16 {
	return (c.gasPedal - c.brakePedal)
}

func (c *car) newSpeed(speed float64){
	c.topSpeeKmh = speed
}

func main(){
	var sampleCar car = car{gasPedal: 1000,
						brakePedal:200,
						steeringWheel: 12,
						topSpeeKmh: 12345.43}

	// fmt.Println("SHOULD BE 1000: ", sampleCar.gasPedal)
	// fmt.Println("TOTAL PEDAL: ", sampleCar.totalPedal())
	fmt.Println("OLD SPEED: ", sampleCar.topSpeeKmh)
	sampleCar.newSpeed(500.43)
	fmt.Println("NEW SPEED: ", sampleCar.topSpeeKmh)

}