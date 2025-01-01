package vehicles

import "github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"

type Vehicle struct {
	LicenseNumber string
	Type          vehicle_types.VehicleType
	Cost          int
}

func NewVehicle(licenseNumber string, vehicleType vehicle_types.VehicleType) *Vehicle {
	return &Vehicle{
		LicenseNumber: licenseNumber,
		Type:          vehicleType,
		Cost:          vehicle_types.GetCost(vehicleType),
	}
}

type IVehicle interface {
	GetLicenseNumber() string
	GetType() vehicle_types.VehicleType
	GetCost() int
}

func (v *Vehicle) GetLicenseNumber() string {
	return v.LicenseNumber
}

func (v *Vehicle) GetType() vehicle_types.VehicleType {
	return v.Type
}

func (v *Vehicle) GetCost() int {
	return v.Cost
}
