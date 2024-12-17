package main

import (
	"cloudbees/ticket-api/apperror"
	"cloudbees/ticket-api/config"
	"cloudbees/ticket-api/datastore"
	"cloudbees/ticket-api/proto"
	"cloudbees/ticket-api/svc"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	ticketSvc svc.TicketService
	proto.UnimplementedTicketAPIServer
}

func main() {
	config.LoadConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", viper.GetInt("port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	store := datastore.NewStore(viper.GetInt("train_info.section_capacity"))
	ticketSvc := svc.NewTicketSvc(store)

	grpcServer := grpc.NewServer()
	proto.RegisterTicketAPIServer(grpcServer, &server{
		ticketSvc: ticketSvc,
	})
	grpcServer.Serve(lis)

}

func (s *server) BookTicket(_ context.Context, request *proto.BookTicketRequest) (*proto.Ticket, error) {
	ticket, err := s.ticketSvc.BookTicket(request.User.FirstName, request.User.LastName, request.User.Email, request.Section.String())
	if err != nil {
		return nil, getgrpcError(err)
	}
	return &proto.Ticket{
		TicketId:   ticket.TicketID.String(),
		From:       ticket.From,
		To:         ticket.To,
		Price:      int32(ticket.Price),
		SeatNumber: int32(ticket.SeatNumber),
		User: &proto.User{
			FirstName: ticket.User.FirstName,
			LastName:  ticket.User.LastName,
			Email:     ticket.User.Email,
		},
	}, nil
}

func (s *server) GetTicket(_ context.Context, request *proto.GetTicketRequest) (*proto.Ticket, error) {
	ticketID, err := uuid.Parse(request.TicketId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid TicketID %s", request.TicketId))
	}
	ticket, err := s.ticketSvc.GetTicket(ticketID)
	if err != nil {
		return nil, getgrpcError(err)
	}
	return &proto.Ticket{
		TicketId:   ticket.TicketID.String(),
		From:       ticket.From,
		To:         ticket.To,
		Price:      int32(ticket.Price),
		SeatNumber: int32(ticket.SeatNumber),
		User: &proto.User{
			FirstName: ticket.User.FirstName,
			LastName:  ticket.User.LastName,
			Email:     ticket.User.Email,
		},
	}, nil
}

func (s *server) ListTickets(_ context.Context, request *proto.ListTicketRequest) (*proto.Tickets, error) {
	tickets, err := s.ticketSvc.ListTickets(request.Section.String(), int(request.PageNum))
	if err != nil {
		return nil, getgrpcError(err)
	}
	results := []*proto.Ticket{}
	for _, ticket := range tickets {
		results = append(results, &proto.Ticket{
			TicketId:   ticket.TicketID.String(),
			From:       ticket.From,
			To:         ticket.To,
			Price:      int32(ticket.Price),
			SeatNumber: int32(ticket.SeatNumber),
			User: &proto.User{
				FirstName: ticket.User.FirstName,
				LastName:  ticket.User.LastName,
				Email:     ticket.User.Email,
			},
		})
	}
	return &proto.Tickets{
		Tickets: results,
	}, nil
}

func (s *server) ModifyTicket(_ context.Context, request *proto.ModifyTicketRequest) (*proto.Ticket, error) {
	ticketID, err := uuid.Parse(request.TicketId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid TicketID %s", request.TicketId))
	}
	ticket, err := s.ticketSvc.ModifyTicket(ticketID, request.User.FirstName, request.User.LastName, request.User.Email)
	if err != nil {
		return nil, getgrpcError(err)
	}
	return &proto.Ticket{
		TicketId:   ticket.TicketID.String(),
		From:       ticket.From,
		To:         ticket.To,
		Price:      int32(ticket.Price),
		SeatNumber: int32(ticket.SeatNumber),
		User: &proto.User{
			FirstName: ticket.User.FirstName,
			LastName:  ticket.User.LastName,
			Email:     ticket.User.Email,
		},
	}, nil
}

func (s *server) CancelTicket(_ context.Context, request *proto.GetTicketRequest) (*emptypb.Empty, error) {
	ticketID, err := uuid.Parse(request.TicketId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid TicketID %s", request.TicketId))
	}
	err = s.ticketSvc.CancelTicket(ticketID)
	if err != nil {
		return nil, getgrpcError(err)
	}
	return &emptypb.Empty{}, nil
}

func getgrpcError(err error) error {
	switch appErr := err.(type) {
	case apperror.ValidationError:
		return status.Errorf(codes.InvalidArgument, appErr.Error())
	case apperror.NotFoundError:
		return status.Errorf(codes.NotFound, appErr.Error())
	default:
		return status.Errorf(codes.Unknown, appErr.Error())
	}
}
