package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Vehicle interface {
	GetLicenseNumber() string
	GetType() vehicle_types.VehicleType
	GetCost() int
}

type BaseVehicle struct {
	LicenseNumber string
	Type          vehicle_types.VehicleType
	Cost          int
}

func NewBaseVehicle(licenseNumber string, vehicleType vehicle_types.VehicleType) *BaseVehicle {
	return &BaseVehicle{
		LicenseNumber: licenseNumber,
		Type:          vehicleType,
		Cost:          vehicle_types.GetCost(vehicleType),
	}
}

func (b *BaseVehicle) GetLicenseNumber() string {
	return b.LicenseNumber
}

func (b *BaseVehicle) GetType() vehicle_types.VehicleType {
	return b.Type
}

func (b *BaseVehicle) GetCost() int {
	return b.Cost
}
