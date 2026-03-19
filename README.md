# DealMaxxingCLI - Finding Discounted Game

An interactive CLI (Command-Line Interface) tool that handling APIs logic to search for game ID and use it to find all available stores that sell those game, such as Steam, 2game, Gamersgate, Fanatical, etc. The tool also implemented in-memory caching to record search history, hashmap indexing to optimize data lookup, user input sanitization and error handling to prevent unforseen error and search results.

## Technology

- Go (Stdlib and fatih/color package)

## Features

- help: Display the available commands.
- id: Find the searched title identifier.
- get: Use the unique ID to find all of the available deals in different stores.

## Preview

### Help command

<p align="center"><img width="615" height="361" src="https://github.com/user-attachments/assets/16ac1b65-4f13-48b0-a8d1-d547041ba557" /></p>

### ID command

<p align="center"><img width="615" height="361" src="https://github.com/user-attachments/assets/4d8fd18c-8063-49ee-8eda-4066b0bdd19e"/>

### Get command

<p align="center"><img width="615" height="361" src="https://github.com/user-attachments/assets/96fa0945-cad1-4747-8dae-61182d173eef" /></p>

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
