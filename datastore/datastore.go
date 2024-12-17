package datastore

import (
	"cloudbees/ticket-api/models"
	"fmt"

	"github.com/google/uuid"
)

type DataStore interface {
	GetAvailableSeat(string) (*int, error)
	SaveTicket(models.Ticket) error
	CancelTicket(uuid.UUID) error
	GetTicket(uuid.UUID) (*models.Ticket, error)
	UpdateTicket(models.Ticket) error
	ListTickets(string, int, int) ([]models.Ticket, error)
}

type Store struct {
	SectionTicketIDMap map[string][]uuid.UUID
	TicketIDTicketMap  map[uuid.UUID]models.Ticket
}

func NewStore(sectionCapacity int) DataStore {
	sectionAArr := make([]uuid.UUID, sectionCapacity)
	sectionBArr := make([]uuid.UUID, sectionCapacity)
	return &Store{
		SectionTicketIDMap: map[string][]uuid.UUID{
			"A": sectionAArr,
			"B": sectionBArr,
		},
		TicketIDTicketMap: map[uuid.UUID]models.Ticket{},
	}
}

func (s *Store) GetAvailableSeat(section string) (*int, error) {
	sectionArr, ok := s.SectionTicketIDMap[section]
	if !ok {
		return nil, fmt.Errorf("Invalid Section %s", section)
	}
	for i, ticketID := range sectionArr {
		if ticketID == uuid.Nil {
			seat := i + 1
			return &seat, nil
		}
	}
	return nil, fmt.Errorf("No more seats available in Section %s", section)
}

func (s *Store) SaveTicket(ticket models.Ticket) error {
	sectionArr, ok := s.SectionTicketIDMap[ticket.Section]
	if !ok {
		return fmt.Errorf("Invalid Section %s", ticket.Section)
	}
	if ticket.SeatNumber < 1 || ticket.SeatNumber > len(sectionArr) {
		return fmt.Errorf("Invalid Seat Number %d", ticket.SeatNumber)
	}
	if sectionArr[ticket.SeatNumber-1] != uuid.Nil {
		return fmt.Errorf("Seat Number %d is already booked", ticket.SeatNumber)
	}
	sectionArr[ticket.SeatNumber-1] = ticket.TicketID
	s.TicketIDTicketMap[ticket.TicketID] = ticket

	return nil
}

func (s *Store) CancelTicket(ticketID uuid.UUID) error {
	ticket, ok := s.TicketIDTicketMap[ticketID]
	if !ok {
		return fmt.Errorf("Invalid TicketID %s", ticketID.String())
	}
	if sectionArr, ok := s.SectionTicketIDMap[ticket.Section]; ok {
		if ticket.SeatNumber > 0 && ticket.SeatNumber <= len(sectionArr) {
			sectionArr[ticket.SeatNumber-1] = uuid.Nil
		}
	}
	delete(s.TicketIDTicketMap, ticketID)
	return nil
}

func (s *Store) GetTicket(ticketID uuid.UUID) (*models.Ticket, error) {
	if ticket, ok := s.TicketIDTicketMap[ticketID]; ok {
		return &ticket, nil
	}
	return nil, fmt.Errorf("Ticket not found")
}

func (s *Store) UpdateTicket(ticket models.Ticket) error {
	_, ok := s.TicketIDTicketMap[ticket.TicketID]
	if !ok {
		return fmt.Errorf("Invalid TicketID %s", ticket.TicketID.String())
	}
	s.TicketIDTicketMap[ticket.TicketID] = ticket
	return nil
}

func (s *Store) ListTickets(section string, pageNum, pageSize int) ([]models.Ticket, error) {
	results := []models.Ticket{}
	sectionArr, ok := s.SectionTicketIDMap[section]
	if !ok {
		return results, fmt.Errorf("Invalid Section %s", section)
	}
	offset := (pageNum - 1) * pageSize
	count := 0
	for _, ticketID := range sectionArr {
		if ticketID != uuid.Nil {
			count += 1
			if count > offset {
				results = append(results, s.TicketIDTicketMap[ticketID])
				if len(results) >= pageSize {
					break
				}
			}
		}
	}
	return results, nil

}
