# TCP Server

A TCP server implemented from scratch with Go. This project serves as an introduction to basic networking and concurrency in Go, demonstrating how to handle multiple clients simultaneously without blocking.

## Features

* **TCP Connection Handling:** Listens on a specific port and accepts incoming connections.
* **Concurrency:** Uses Go's lightweight goroutines to handle multiple clients at the same time.
* **Echo Functionality:** Reads messages sent by the client and echoes them back.
* **Graceful Exit:** Allows clients to disconnect by typing `exit`.

## Prerequisites

* **Go** (version 1.18 or higher recommended)
* **Netcat (`nc`)** for testing client connections

## Getting Started

### 1. Clone the Repository

If you haven't already, clone this repository to your local machine:

```bash
git clone git@github.com:ruinabeshima/tcp-server.git
cd tcp-server
```

### 2. Run the Server

Start the TCP server using the `go run` command. It will begin listening on port 8080:

```bash
go run main.go
```

You should see the output:

```text
TCP server is listening on port 8080...
```

### 3. Connect as a Client

Open a new terminal window (leave the server running in the first one) and connect using Netcat:

```bash
nc localhost 8080
```