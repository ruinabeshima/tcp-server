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
}