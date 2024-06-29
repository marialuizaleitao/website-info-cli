# Website Info CLI

Website Info CLI is a command-line application written in Go that allows you to search for IP addresses and server names on the Internet using DNS lookup.

## Features

- **Search IP Addresses**: Retrieve IP addresses associated with a given host name.
- **Search Server Names**: Retrieve server names associated with a given host name.

## Usage

### Installation

Make sure you have Go installed. Then, clone the repository:

```bash
git clone https://github.com/marialuizaleitao/website-info-cli.git
cd website-info-cli
```

### Running the Application
To run the application, use the following command:

```bash
go run main.go
```

#### Commands
**ip Command**
Searches for IP addresses associated with a given host name.

```bash
go run main.go ip --host example.com
```
Replace example.com with the host name you want to search.

**servers Command**
Searches for server names associated with a given host name.

```bash
go run main.go servers --host example.com
```
Replace example.com with the host name you want to search.

## Dependencies

`github.com/urfave/cli`: A powerful package for creating command-line interfaces in Go.      
`net`: Standard Go package for network operations.

## Contributing
Contributions are welcome! If you have suggestions or find any issues, please open an issue or submit a pull request.

