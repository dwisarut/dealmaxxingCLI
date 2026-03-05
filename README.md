# DealMaxxingCLI - CLI tool for finding discounted game

An interactive CLI (Command-Line Interface) tool that handling APIs logic such as searching, and querying. The final output is the links of all available deals from multiple stores such as Steam, Gamersgate, Fanatical, etc.

## File Structure
```
dealmaxxingCLI/
├── cmd/                          # cmd package file
|   ├── parser.go                 # input parser logics
|	└── root.go                   # CLIs logics
|
├── internal/
|	├── api                       # api package file
|	|	└── api.go                # http request and response
|   |
|	├── model                     # model package
|	|	└── data.go               # define data model
|   |
|   └── service                   # service package
|       └── deal.go               # handling filter, sort, etc.
|
├── go.mod
├── go.sum
└── main.go                       # entry point file
```
