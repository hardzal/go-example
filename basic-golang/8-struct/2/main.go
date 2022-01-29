package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel uint16
	topSpeedKmh   float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmhMultiple)
}

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func newerTopSpeed(c car, speed float64) car {
	c.topSpeedKmh = speed
	return c
}

func main() {
	aCar := car{
		gasPedal:      65500,
		brakePedal:    0,
		steeringWheel: 1256,
		topSpeedKmh:   225.0}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())

	// aCar.newTopSpeed(250)
	aCar = newerTopSpeed(aCar, 600)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}
