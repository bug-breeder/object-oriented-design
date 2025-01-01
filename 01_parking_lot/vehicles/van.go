package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Van struct {
	*BaseVehicle
}

func NewVan(licenseNumber string) *Van {
	return &Van{
		BaseVehicle: NewBaseVehicle(licenseNumber, vehicle_types.Van),
	}
}
