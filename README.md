# GoCreate

GoCreate is a command-line interface (CLI) application built in Go. It's primary purpose is for me to up my Go game. Original concept is a simple learning tool for web development, providing a linear, user-friendly interface to experiment and learn.

## Getting Started

To run the application, navigate to the project directory and use the `go run` command:

```bash
go run . website
```

## Notables

- **Learning Web Development**: GoCreate provides a hands-on approach to learning web development. It’s a simple tool for web development, providing a linear, user-friendly interface to experiment and learn.
- **Command Line Interface (CLI)**: The project uses the Cobra package to provide a powerful, modern CLI experience.
- **Server-Sent Events (SSE)**: With the use of certain packages, the project can handle real-time bidirectional communication between the server and the client.
- **Terminal-based User Interface**: The project uses a specific package to build rich terminal-based user interfaces.
- **Logging**: The project uses a specific package for logging purposes.

Please note that this is a personal project and is still under development (WIP). The features may change as the project evolves.

## Testing

GoCreate uses the `ExecuteCommand` function for testing its commands. This function executes a command and captures its output, allowing for assertions about the command's behavior.

For more information on testing Cobra CLI applications, refer to this [guide](https://jackwrfuller.au/posts/testing-cobra-cli/).

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License.
