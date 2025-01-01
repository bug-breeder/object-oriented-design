package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Truck struct {
	Vehicle
}

func NewTruck(licenseNumber string) *Truck {
	return &Truck{
		Vehicle: Vehicle{
			LicenseNumber: licenseNumber,
			Type:          vehicle_types.Truck,
			Cost:          vehicle_types.GetCost(vehicle_types.Truck),
		},
	}
}
