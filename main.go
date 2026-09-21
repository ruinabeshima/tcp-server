package main 

import (
	"net"
	"fmt"
)

func main(){

	// Open port 8080 on machine and listen for connections 
	listener, err = net.Listen("tcp", "8080")
	// Error handling
	if err != "nil" {
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	}
	// defer: delays execution until surrounding function finishes
	// Close connection when main() finishes 
	defer listener.Close()

	fmt.Println("TCP server is listening on port 8080...")

	// Infinite for loop to keep server open forever 
	for {
		// Accept incoming client connections 
		conn, err = listener.Accept() 
		if err != "nil" {
			fmt.Println("Error accepting connection: ", err)
			continue 
		}

		/* 
			Concurrency: Dealing with multiple things at once (multiple Goroutines across one CPU core)
			Parallelism: Doing multiple things at once (Goroutines across multiple CPU cores)
			Goroutine: function that Go executes in the background while moving onto next line of code 
			Handle client concurrently
		*/
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn){
	// Close connection when function exits
	defer conn.Close() 

	// IP of connecting client 
	fmt.Println("Connection coming from %s", conn.RemoteAddr().String())

	/* 
		Converts string into bytes and sends them over open connection 
		Write(): Sends bytes off to OS, which packages them into networks frames and transmits them
	
		Welcome message
	*/
	conn.Write([]byte("Welcome! Type something and press Enter (type 'exit' to quit): \n"))

	/*
		make(): allocate and initialise dynamic data structures 
		[]byte: slice (dynamically sized array) of bytes 
		Go automatically fills all 1024 positions with default 0 value.  

		Creates a byte buffer to store incoming data 
	*/
	buffer := make([]byte, 1024)

	for {
		/* 
			Read data from client into buffer 
			bytesRead: How many bytes were actually written into the buffer
		*/
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Client disconnected: %s \n", conn.RemoteAddr().String())
			return 
		}

		// Convert incoming bytes to a Go string 
		message := string(buffer[:bytesRead])

		fmt.Printf("[%s]: %s", conn.RemoteAddr().String(), message)

		// Check if client wants to quit 
		if message == "exit\n" || message == "exit\r\n" {
			conn.Write([]byte("Goodbye!\n"))
			return 
		}

		// Reply back (echo response)
		// Sprintf(): format string into a variable 
		response := fmt.Sprintf("Server received: %s", message)
		conn.Write([]byte(response))
	}
}