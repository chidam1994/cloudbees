package models

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	TicketID   uuid.UUID `json:"ticket_id"`
	User       User      `json:"user"`
	From       string    `json:"from"`
	To         string    `json:"to"`
	Price      int       `json:"price"`
	CreatedAt  time.Time `json:"create_at"`
	UpdatedAt  time.Time `json:"updatedAt"`
	SeatNumber int       `json:"seat_number"`
	Section    string    `json:"section" validate:"oneof=A B"`
}

type User struct {
	FirstName string `json:"first_name" validate:"required,min=1,max=20"`
	LastName  string `json:"last_name" validate:"required,min=1,max=20"`
	Email     string `json:"email" validate:"required,email"`
}
