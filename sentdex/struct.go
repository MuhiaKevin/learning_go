package main

import "fmt"

const usixteenbitmax = 65535
const kmh_multiple = 1.60934

// struct
type car struct {
	gas_pedal      uint16
	break_pedal    uint16
	steering_wheel int16
	top_speed      float64
}

// methods VALUE RECEIVERS => METHODS THAT RECEIVE VALUES AND DO SOME LOGIC
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed / usixteenbitmax)
}
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed / usixteenbitmax / kmh_multiple)
}

// methods POINTER RECEIVERS MODIFY THE VALUES OF STRUCT
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed = newspeed
}

func main() {
	car := car{gas_pedal: 22341, break_pedal: 5656, steering_wheel: 89, top_speed: 56.65}
	fmt.Println(car.kmh())
	car.new_top_speed(340.23)
	fmt.Println(car.mph())
}
