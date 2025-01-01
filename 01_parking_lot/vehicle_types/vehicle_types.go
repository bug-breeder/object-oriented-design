package vehicle_types

type VehicleType int

const (
	Car VehicleType = iota
	Van
	Truck
	Motorcycle
)

var vehicleTypeCost = map[VehicleType]int{
	Car:        100,
	Van:        200,
	Truck:      300,
	Motorcycle: 50,
}

func GetCost(vehicleType VehicleType) int {
	return vehicleTypeCost[vehicleType]
}

func GetVehicleType(vehicleType string) VehicleType {
	switch vehicleType {
	case "car":
		return Car
	case "van":
		return Van
	case "truck":
		return Truck
	case "motorcycle":
		return Motorcycle
	default:
		return -1
	}
}
