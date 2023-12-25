package repository

import "mybookings.com/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(reservation models.Reservation) (int, error)

	InsertRoomRestriction(r models.RoomRestriction) error
}
