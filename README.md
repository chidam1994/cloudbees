# Ticket Management API
## API Setup
- Install Go 1.22.1 or above
- Clone the repository and change working directory to the repo's root directory
- Run the following command to download the dependencies
```
go mod tidy
```
- Create a ticket-api.json config file by referring to config/ticket-api.example.json file.
- Start the grpc Server by running
```
go run ./cmd/api/main.go
```

## Testing the API
- Start the grpc client by running
```
go run ./cmd/client/main.go
```