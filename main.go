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
	}

}