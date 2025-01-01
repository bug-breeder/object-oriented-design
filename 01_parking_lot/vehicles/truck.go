package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Truck struct {
	*BaseVehicle
}

func NewTruck(licenseNumber string) *Truck {
	return &Truck{
		BaseVehicle: NewBaseVehicle(licenseNumber, vehicle_types.Truck),
	}
}
