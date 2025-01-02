package parking_lot

import (
	"time"

	"github.com/bug-breeder/object-oriented-design/01_parking_lot/vehicles"
)

// ParkingTicket represents a parking ticket issued to a vehicle
type ParkingTicket struct {
	EntryTime time.Time
	ExitTime  time.Time
	Vehicle   vehicles.Vehicle
	Spot      *ParkingSpot
}

// NewParkingTicket creates a new parking ticket
func NewParkingTicket(vehicle vehicles.Vehicle, spot *ParkingSpot) *ParkingTicket {
	return &ParkingTicket{
		EntryTime: time.Now(),
		Vehicle:   vehicle,
		Spot:      spot,
	}
}

// SetExitTime sets the exit time on the parking ticket
func (pt *ParkingTicket) SetExitTime() {
	pt.ExitTime = time.Now()
}

// CalculateTotalCharge calculates the total charge for the parking ticket
func (pt *ParkingTicket) CalculateTotalCharge() float64 {
	if pt.ExitTime.IsZero() {
		pt.SetExitTime()
	}

	// Calculate the duration for which the vehicle was parked
	duration := pt.ExitTime.Sub(pt.EntryTime)
	hours := duration.Hours()

	// Calculate the total charge based on the vehicle type
	return hours * float64(pt.Vehicle.GetCost())
}
