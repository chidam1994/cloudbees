package datastore

import (
	"testing"

	"cloudbees/ticket-api/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAvailableSeat(t *testing.T) {
	tests := []struct {
		name          string
		section       string
		setupStore    func(*Store)
		expectedSeat  *int
		expectedError string
	}{
		{
			name:    "Empty Section",
			section: "A",
			setupStore: func(s *Store) {
			},
			expectedSeat:  intPtr(1),
			expectedError: "",
		},
		{
			name:    "Partially Filled Section",
			section: "A",
			setupStore: func(s *Store) {
				s.SectionTicketIDMap["A"][0] = uuid.New()
				s.SectionTicketIDMap["A"][1] = uuid.New()
			},
			expectedSeat:  intPtr(3),
			expectedError: "",
		},
		{
			name:    "Invalid Section",
			section: "C",
			setupStore: func(s *Store) {
			},
			expectedSeat:  nil,
			expectedError: "Invalid Section C",
		},
		{
			name:    "Full Section",
			section: "A",
			setupStore: func(s *Store) {
				for i := range s.SectionTicketIDMap["A"] {
					s.SectionTicketIDMap["A"][i] = uuid.New()
				}
			},
			expectedSeat:  nil,
			expectedError: "No more seats available in Section A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore(5)
			tt.setupStore(store.(*Store))

			seat, err := store.GetAvailableSeat(tt.section)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.Nil(t, seat)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedSeat, seat)
			}
		})
	}
}

func TestSaveTicket(t *testing.T) {
	tests := []struct {
		name          string
		ticket        models.Ticket
		setupStore    func(*Store)
		expectedError string
	}{
		{
			name: "Success",
			ticket: models.Ticket{
				TicketID:   uuid.New(),
				Section:    "A",
				SeatNumber: 1,
			},
			setupStore: func(s *Store) {
			},
			expectedError: "",
		},
		{
			name: "Invalid Section",
			ticket: models.Ticket{
				TicketID:   uuid.New(),
				Section:    "C",
				SeatNumber: 1,
			},
			setupStore: func(s *Store) {
			},
			expectedError: "Invalid Section C",
		},
		{
			name: "Invalid Seat Number",
			ticket: models.Ticket{
				TicketID:   uuid.New(),
				Section:    "A",
				SeatNumber: 0,
			},
			setupStore: func(s *Store) {
			},
			expectedError: "Invalid Seat Number 0",
		},
		{
			name: "Invalid Seat Number",
			ticket: models.Ticket{
				TicketID:   uuid.New(),
				Section:    "A",
				SeatNumber: 6,
			},
			setupStore: func(s *Store) {
			},
			expectedError: "Invalid Seat Number 6",
		},
		{
			name: "Seat Already Booked",
			ticket: models.Ticket{
				TicketID:   uuid.New(),
				Section:    "A",
				SeatNumber: 1,
			},
			setupStore: func(s *Store) {
				s.SectionTicketIDMap["A"][0] = uuid.New()
			},
			expectedError: "Seat Number 1 is already booked",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore(5)
			tt.setupStore(store.(*Store))

			err := store.SaveTicket(tt.ticket)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCancelTicket(t *testing.T) {
	testTicket := models.Ticket{
		TicketID:   uuid.New(),
		Section:    "A",
		SeatNumber: 1,
	}
	tests := []struct {
		name          string
		ticketID      uuid.UUID
		setupStore    func(*Store)
		expectedError string
	}{
		{
			name:     "Success",
			ticketID: testTicket.TicketID,
			setupStore: func(s *Store) {
				s.SectionTicketIDMap["A"][0] = testTicket.TicketID
				s.TicketIDTicketMap[testTicket.TicketID] = testTicket
			},
			expectedError: "",
		},
		{
			name:     "InvalidTicketID",
			ticketID: uuid.New(),
			setupStore: func(s *Store) {
			},
			expectedError: "Invalid TicketID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore(5)
			tt.setupStore(store.(*Store))

			err := store.CancelTicket(tt.ticketID)

			if tt.expectedError != "" {
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)

				_, exists := store.(*Store).TicketIDTicketMap[tt.ticketID]
				assert.False(t, exists, "Ticket should be removed from TicketIDTicketMap")

				assert.Equal(t,
					uuid.Nil,
					store.(*Store).SectionTicketIDMap[testTicket.Section][testTicket.SeatNumber-1],
					"Seat should be marked as available")
			}
		})
	}
}

func TestGetTicket(t *testing.T) {
	testTicket := models.Ticket{
		TicketID:   uuid.New(),
		Section:    "B",
		SeatNumber: 1,
	}
	tests := []struct {
		name          string
		ticketID      uuid.UUID
		setupStore    func(*Store)
		expectedError string
	}{
		{
			name:     "Success",
			ticketID: testTicket.TicketID,
			setupStore: func(s *Store) {
				s.TicketIDTicketMap[testTicket.TicketID] = testTicket
			},
			expectedError: "",
		},
		{
			name:     "Ticket Not Found",
			ticketID: uuid.New(),
			setupStore: func(s *Store) {
			},
			expectedError: "Ticket not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore(5)
			tt.setupStore(store.(*Store))

			result, err := store.GetTicket(tt.ticketID)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testTicket, *result)
			}
		})
	}
}

func intPtr(i int) *int {
	return &i
}
