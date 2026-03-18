# DealMaxxingCLI - Finding Discounted Game
<img width="1228" height="720" alt="dmx_gitcover" src="https://github.com/user-attachments/assets/3236fb00-ec7e-439f-8bcf-d46b8c29d294" />

An interactive CLI (Command-Line Interface) tool that handling APIs logic such as searching, and querying. The final output is the links of all available deals from multiple stores such as Steam, Gamersgate, Fanatical, etc. Implemented the basic caching to record search history, hashmap to search for stores data, and user input sanitization to prevent unforseen error and search results.

## Features
- help: Display the available commands.
- id: Find the searched title identifier.
- get: Use the unique ID to find all of the available deals in different stores.

## Getting started

### File Structure

```
dealmaxxingCLI/
├── cmd/                              # cmd files
|   └── cli/                          # cli package file
|       ├── get.go                    # Get command logic
|       ├── id.go                     # ID command logic
|       ├── parser.go                 # Input parser logic
|       └── root.go                   # CLIs logics
|
├── internal/
|	├── api                           # api package file
|	|	└── api.go                    # HTTP request and response
|	|
|	├── cache                         # cache package file
|	|	└── cache.go                  # Caching search data
|	|
|	├── model                         # model package
|	|	└── data.go                   # Define data model
|   |
|   └── service                       # service package
|       └── deal.go                   # Handling filter, sort, etc.
|
├── go.mod
├── go.sum
└── main.go                           # Entry point file
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
