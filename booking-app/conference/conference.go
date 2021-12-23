package conference

import (
	"booking-app/booking"
)

type Conference struct {
	Name             string
	AvailableTickets uint
	Bookings         []*booking.Booking
}
