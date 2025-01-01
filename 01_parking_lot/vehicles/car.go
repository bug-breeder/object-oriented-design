package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Car struct {
	*BaseVehicle
}

func NewCar(licenseNumber string) *Car {
	return &Car{
		BaseVehicle: NewBaseVehicle(licenseNumber, vehicle_types.Car),
	}
}
