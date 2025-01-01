package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Motorcycle struct {
	*BaseVehicle
}

func NewMotorcycle(licenseNumber string) *Motorcycle {
	return &Motorcycle{
		BaseVehicle: NewBaseVehicle(licenseNumber, vehicle_types.Motorcycle),
	}
}
