package common

import "net"

func HandleConnection(conn net.Conn, server *Server) {
	defer conn.Close()

	// Handle the connection. This could involve reading data from the
	// connection, updating the server state, etc.
}
