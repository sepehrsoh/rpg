# Golang Reverse Proxy for Port Forwarding

Welcome to the Golang Reverse Proxy for Port Forwarding project! This repository contains a simple Go application that acts as a reverse proxy, allowing you to forward traffic from a local port to a target IP and port. This can be useful for various networking scenarios, such as exposing a service running on a local machine to the internet or forwarding requests to a backend server.

## How to Use

Follow the steps below to set up and use the reverse proxy:

### Prerequisites

Before you begin, ensure you have the following installed on your system:

1. Go 
2. Git

### Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/golang-reverse-proxy.git
   ```

2. Use the provided Makefile to download and vendor dependencies, and build the project:

   ```bash
   make all
   ```

### Running the Reverse Proxy

To start the reverse proxy, use the following command:

```bash
./rpg run -ip TARGETIP --to TARGETPORT --from LOCALPORT
```

- Replace `TARGETIP` with the IP address of the target server where you want to forward the traffic.
- Replace `TARGETPORT` with the port on the target server where you want to forward the traffic.
- Replace `LOCALPORT` with the local port on your machine from which you want to forward the traffic.

Example:

```bash
./rpg run -ip 192.168.1.100 --to 8080 --from 8888
```

This will set up the reverse proxy to forward traffic from `localhost:8888` to `192.168.1.100:8080`.

### Stopping the Reverse Proxy

To stop the reverse proxy, press `Ctrl+C` in the terminal where it's running. The proxy will gracefully shut down and stop forwarding traffic.

## Contributing

Feel free to contribute to this project by submitting pull requests or reporting issues. Your contributions are greatly appreciated!

## License

This project is licensed under the [MIT License](LICENSE).

Happy port forwarding!
