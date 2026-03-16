# DealMaxxingCLI - Finding Discounted Game

An interactive CLI (Command-Line Interface) tool that handling APIs logic such as searching, and querying. The final output is the links of all available deals from multiple stores such as Steam, Gamersgate, Fanatical, etc.

## Getting started

### File Structure

```
dealmaxxingCLI/
├── cmd/                          # cmd package file
|   ├── get.go                    # Get command logic
|   ├── parser.go                 # Input parser logics
|   ├── root.go                   # CLIs logics
|	└── search.go                 # Search command logic
|
├── internal/
|	├── api                       # api package file
|	|	└── api.go                # HTTP request and response
|   |
|	├── model                     # model package
|	|	└── data.go               # Define data model
|   |
|   └── service                   # service package
|       └── deal.go               # Handling filter, sort, etc.
|
├── go.mod
├── go.sum
└── main.go                       # Entry point file
```

### Installation

1. Clone this repository

   ```
   git clone https://github.com/dwisarut/dealmaxxingCLI.git
   cd dealmaxxingCLI
   ```

2. Install dependencies

   Installing all depedencies

   ```
   go mod tidy
   ```

3. Run this project
   ```
   go run ./main.go
   ```
   Alternatively, you can build a main.exe and open it from there.
   ```
   go build ./main.go
   go run ./main.go
   ```
