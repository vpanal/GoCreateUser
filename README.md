# GoCreateUser

GoCreateUser is a Go-based utility for creating a Windows user and adding the user to a specified group using Windows API functions.

## Features

- Creates a new user with specified credentials.
- Adds the user to a specified local group.
- Debugging support for detailed logs.

## Prerequisites

- Go 1.24 or later.
- Windows operating system.
- `golang.org/x/sys` package (already included in `go.mod`).

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/vpanal/GoCreateUser
   cd gocreateuser
   ```

2. Build the project:
   ```sh
   go build -o gocreateuser.exe main.go
   ```

## Usage

1. Update the `username`, `password`, and `group` constants in `main.go` with the desired values.

2. Run the executable:
   ```sh
   ./gocreateuser.exe
   ```

3. The program will:
   - Create a user with the specified credentials.
   - Add the user to the specified group.
   - Print debug messages if the `dbg` constant is set to `true`.

## Debugging

To enable debugging, set the `dbg` constant in `main.go` to `true`. Debug messages will be printed to the console.

## Notes

- Ensure you have administrative privileges to run the program.
- The `username` and `password` constants should be securely managed in a production environment.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
