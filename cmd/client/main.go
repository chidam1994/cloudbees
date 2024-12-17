package main

import (
	"bufio"
	"cloudbees/ticket-api/config"
	"cloudbees/ticket-api/proto"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config.LoadConfig()
	serverAddr := fmt.Sprintf("localhost:%d", viper.GetInt("port"))
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to initialize grpc client: %v", err)
	}
	defer conn.Close()

	client := proto.NewTicketAPIClient(conn)

MAINLOOP:
	for {
		fmt.Println("Enter 1 to Purchase a Ticket")
		fmt.Println("Enter 2 to view details of a Ticket")
		fmt.Println("Enter 3 to list tickets belonging to a section")
		fmt.Println("Enter 4 to cancel a ticket")
		fmt.Println("Enter 5 to modify a ticket")
		fmt.Println("Enter 6 to close grpc-client")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		switch strings.TrimSpace(scanner.Text()) {
		case "1":
			ticket, err := bookTicket(scanner, client)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
				continue
			}
			fmt.Println("Ticket purchased Successfully!!")
			fmt.Println(ticket)
			fmt.Println()
		case "2":
			ticket, err := getTicket(scanner, client)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
				continue
			}
			fmt.Println("Ticket Details:")
			fmt.Println(ticket)
			fmt.Println()
		case "3":
			tickets, err := listTickets(scanner, client)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
				continue
			}
			fmt.Println("Tickets:")
			fmt.Println(tickets)
			fmt.Println()
		case "4":
			err := cancelTicket(scanner, client)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
				continue
			}
			fmt.Println("Ticket Cancelled Successfully")
			fmt.Println()
		case "5":
			ticket, err := modifyTicket(scanner, client)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
				continue
			}
			fmt.Println("Ticket Details:")
			fmt.Println(ticket)
			fmt.Println()
		case "6":
			break MAINLOOP
		}

	}
}

func getTicket(scanner *bufio.Scanner, client proto.TicketAPIClient) (string, error) {
	fmt.Print("TicketID: ")
	scanner.Scan()
	ticketID := strings.TrimSpace(scanner.Text())
	ticket, err := client.GetTicket(context.Background(), &proto.GetTicketRequest{
		TicketId: ticketID,
	})
	if err != nil {
		return "", err
	}
	ticketBytes, err := json.MarshalIndent(ticket, "", "    ")
	if err != nil {
		return "", fmt.Errorf("Error while marshaling ticket: %s", err.Error())
	}
	return string(ticketBytes), nil
}

func cancelTicket(scanner *bufio.Scanner, client proto.TicketAPIClient) error {
	fmt.Print("TicketID: ")
	scanner.Scan()
	ticketID := strings.TrimSpace(scanner.Text())
	_, err := client.CancelTicket(context.Background(), &proto.GetTicketRequest{
		TicketId: ticketID,
	})
	if err != nil {
		return err
	}
	return nil
}

func modifyTicket(scanner *bufio.Scanner, client proto.TicketAPIClient) (string, error) {
	fmt.Print("TicketID: ")
	scanner.Scan()
	ticketID := strings.TrimSpace(scanner.Text())
	fmt.Print("First Name: ")
	scanner.Scan()
	firstName := strings.TrimSpace(scanner.Text())
	fmt.Print("Last Name: ")
	scanner.Scan()
	lastName := strings.TrimSpace(scanner.Text())
	fmt.Print("Email: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())
	ticket, err := client.ModifyTicket(context.Background(), &proto.ModifyTicketRequest{
		TicketId: ticketID,
		User: &proto.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		},
	})
	if err != nil {
		return "", err
	}
	ticketBytes, err := json.MarshalIndent(ticket, "", "    ")
	if err != nil {
		return "", fmt.Errorf("Error while marshaling ticket: %s", err.Error())
	}
	return string(ticketBytes), nil
}

func listTickets(scanner *bufio.Scanner, client proto.TicketAPIClient) (string, error) {
	fmt.Print("Train Section(A/B): ")
	scanner.Scan()
	section := strings.ToUpper(strings.TrimSpace(scanner.Text()))
	grpcSection, ok := proto.Section_value[section]
	if !ok {
		return "", fmt.Errorf("Invalid section: %s, section must be one of A or B", section)
	}
	fmt.Print("Page Number: ")
	scanner.Scan()
	pageNumStr := strings.TrimSpace(scanner.Text())
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		return "", fmt.Errorf("Invalid pageNum %s", pageNumStr)
	}
	tickets, err := client.ListTickets(context.Background(), &proto.ListTicketRequest{
		Section: proto.Section(grpcSection),
		PageNum: int32(pageNum),
	})
	if err != nil {
		return "", err
	}
	ticketBytes, err := json.MarshalIndent(tickets, "", "    ")
	if err != nil {
		return "", fmt.Errorf("Error while marshaling ticket: %s", err.Error())
	}
	return string(ticketBytes), nil
}

func bookTicket(scanner *bufio.Scanner, client proto.TicketAPIClient) (string, error) {
	fmt.Println("Enter the requested details required to purchase ticket")
	fmt.Print("First Name: ")
	scanner.Scan()
	firstName := strings.TrimSpace(scanner.Text())
	fmt.Print("Last Name: ")
	scanner.Scan()
	lastName := strings.TrimSpace(scanner.Text())
	fmt.Print("Email: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())
	fmt.Print("Train Section(A/B): ")
	scanner.Scan()
	section := strings.ToUpper(strings.TrimSpace(scanner.Text()))
	grpcSection, ok := proto.Section_value[section]
	if !ok {
		return "", fmt.Errorf("Invalid section: %s, section must be one of A or B", section)
	}
	ticket, err := client.BookTicket(context.Background(), &proto.BookTicketRequest{
		Section: proto.Section(grpcSection),
		User: &proto.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		},
	})
	if err != nil {
		return "", err
	}

	ticketBytes, err := json.MarshalIndent(ticket, "", "    ")
	if err != nil {
		return "", fmt.Errorf("Error while marshaling ticket: %s", err.Error())
	}
	return string(ticketBytes), nil
}
