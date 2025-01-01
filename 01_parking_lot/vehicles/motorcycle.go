package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Motorcycle struct {
	Vehicle
}

func NewMotorcycle(licenseNumber string) *Motorcycle {
	return &Motorcycle{
		Vehicle: Vehicle{
			LicenseNumber: licenseNumber,
			Type:          vehicle_types.Motorcycle,
			Cost:          vehicle_types.GetCost(vehicle_types.Motorcycle),
		},
	}
}
