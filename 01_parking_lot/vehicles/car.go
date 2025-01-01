package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Car struct {
	Vehicle
}

func NewCar(licenseNumber string) *Car {
	return &Car{
		Vehicle: Vehicle{
			LicenseNumber: licenseNumber,
			Type:          vehicle_types.Car,
			Cost:          vehicle_types.GetCost(vehicle_types.Car),
		},
	}
}
