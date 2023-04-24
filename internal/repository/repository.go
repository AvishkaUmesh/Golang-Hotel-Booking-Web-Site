package repository

import (
	"time"

	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(reservation models.Reservation) (int, error)
	InsertRoomRestriction(restriction models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomId int) (bool, error)
}
