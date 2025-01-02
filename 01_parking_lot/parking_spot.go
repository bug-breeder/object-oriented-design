package parking_lot

import (
	"sync"

	"github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicle_types"
	"github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicles"
)

// ParkingSpot is a struct that represents a parking spot in a parking lot.
type ParkingSpot struct {
	Id       int
	SpotType vehicle_types.VehicleType
	Vehicle  *vehicles.Vehicle
	lock     sync.Mutex
}

// NewParkingSpot creates a new parking spot.
func NewParkingSpot(id int, spotType vehicle_types.VehicleType) *ParkingSpot {
	return &ParkingSpot{
		Id:       id,
		SpotType: spotType,
	}
}

// Park parks a vehicle in the parking spot.
func (p *ParkingSpot) Park(vehicle vehicles.Vehicle) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.Vehicle != nil {
		panic("Parking spot is already occupied")
	}

	if vehicle.GetType() != p.SpotType {
		panic("Vehicle type does not match parking spot type")
	}

	p.Vehicle = &vehicle
}

// RemoveVehicle removes the vehicle from the parking spot.
func (p *ParkingSpot) RemoveVehicle() {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.Vehicle == nil {
		panic("Parking spot is already empty")
	}

	p.Vehicle = nil
}

// IsOccupied returns true if the parking spot is occupied.
func (p *ParkingSpot) IsOccupied() bool {
	p.lock.Lock()
	defer p.lock.Unlock()

	return p.Vehicle != nil
}

// GetVehicle returns the vehicle parked in the parking spot.
func (p *ParkingSpot) GetVehicle() vehicles.Vehicle {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.Vehicle == nil {
		panic("Parking spot is empty")
	}

	return *p.Vehicle
}
