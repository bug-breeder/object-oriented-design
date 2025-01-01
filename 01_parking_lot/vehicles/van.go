package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Van struct {
	Vehicle
}

func NewVan(licenseNumber string) *Van {
	return &Van{
		Vehicle: Vehicle{
			LicenseNumber: licenseNumber,
			Type:          vehicle_types.Van,
			Cost:          vehicle_types.GetCost(vehicle_types.Van),
		},
	}
}
