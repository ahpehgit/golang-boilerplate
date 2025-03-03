# My Golang Project

This is a simple Go application structured with a command line interface and a package for reusable components.

## Project Structure

```
golang-boilerplate
├── bin
│   └── Dockerfile
├── src
│   └── main.go          # Entry point of the application
│   ├── <mypackage>
│   │       └── mypackage.go # Package containing reusable components
│   ├── go.mod               # Module definition file
|   ├── config.yml           # Config file to store env variables   
│   └── README.md            # Documentation for the project
```

## Getting Started

To get started with this project, ensure you have Go installed on your machine. You can download it from the official Go website.

### Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd golang-boilerplate/src
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

### Running the Application

To run the application, use the following command:
```
go run src/main.go
```

### Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

### License

This project is licensed under the MIT License. See the LICENSE file for details.