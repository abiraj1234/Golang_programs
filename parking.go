
package main 

import "fmt"


type ParkingSystem struct {
    big    int
    medium int
    small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{big: big, medium: medium, small: small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
    switch carType {
    case 1:
        if this.big > 0 {
            this.big--
            return true
        }
    case 2:
        if this.medium > 0 {
            this.medium--
            return true
        }
    case 3:
        if this.small > 0 {
            this.small--
            return true
        }
    }
    return false
}

func main () {
	ps := Constructor(1, 1, 0)
    fmt.Println(ps.AddCar(1)) 
    fmt.Println(ps.AddCar(2)) 
    fmt.Println(ps.AddCar(3)) 
}