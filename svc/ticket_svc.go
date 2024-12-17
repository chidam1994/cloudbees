package svc

import (
	"cloudbees/ticket-api/apperror"
	"cloudbees/ticket-api/datastore"
	"cloudbees/ticket-api/models"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type TicketService interface {
	BookTicket(firstName, lastName, email, section string) (*models.Ticket, error)
	CancelTicket(ticketID uuid.UUID) error
	GetTicket(ticketID uuid.UUID) (*models.Ticket, error)
	ListTickets(section string, pageNum int) ([]models.Ticket, error)
	ModifyTicket(ticketID uuid.UUID, firstName, lastName, email string) (*models.Ticket, error)
}

type TicketSvc struct {
	store datastore.DataStore
}

func NewTicketSvc(store datastore.DataStore) TicketService {
	return &TicketSvc{
		store,
	}
}

func (svc *TicketSvc) BookTicket(firstName, lastName, email, section string) (*models.Ticket, error) {
	ticket := models.Ticket{
		User: models.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		},
		From:     viper.GetString("train_info.from"),
		To:       viper.GetString("train_info.to"),
		Price:    viper.GetInt("train_info.price"),
		TicketID: uuid.New(),
		Section:  strings.ToUpper(section),
	}

	validate := validator.New()

	err := validate.Struct(ticket)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return nil, apperror.ValidationError(fmt.Errorf("Validation error: %s", errors))
	}

	seat, err := svc.store.GetAvailableSeat(section)
	if err != nil {
		return nil, apperror.UnknownError(fmt.Errorf("Error finding Seat: %w", err))
	}

	ticket.SeatNumber = *seat

	err = svc.store.SaveTicket(ticket)
	if err != nil {
		return nil, apperror.UnknownError(fmt.Errorf("Error booking Ticket: %w", err))
	}
	return &ticket, nil
}

func (svc *TicketSvc) GetTicket(ticketID uuid.UUID) (*models.Ticket, error) {
	ticket, err := svc.store.GetTicket(ticketID)
	if err != nil {
		return nil, apperror.NotFoundError(err)
	}
	return ticket, nil
}

func (svc *TicketSvc) CancelTicket(ticketID uuid.UUID) error {
	err := svc.store.CancelTicket(ticketID)
	if err != nil {
		return apperror.NotFoundError(err)
	}
	return nil
}

func (svc *TicketSvc) ModifyTicket(ticketID uuid.UUID, firstName, lastName, email string) (*models.Ticket, error) {
	ticket, err := svc.store.GetTicket(ticketID)
	if err != nil {
		return nil, apperror.NotFoundError(err)
	}
	ticket.User = models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	err = svc.store.UpdateTicket(*ticket)
	if err != nil {
		return nil, apperror.UnknownError(fmt.Errorf("Error modifying Ticket: %w", err))
	}
	return ticket, nil
}

func (svc *TicketSvc) ListTickets(section string, pageNum int) ([]models.Ticket, error) {
	tickets, err := svc.store.ListTickets(section, pageNum, viper.GetInt("list_tickets_page_size"))
	if err != nil {
		return []models.Ticket{}, apperror.UnknownError(err)
	}
	return tickets, nil
}
