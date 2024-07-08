# NetCat: TCP-Chat Server-Client Application
# Project Overview
NetCat is a Go-based implementation of the NetCat utility, designed to facilitate a chat server-client architecture. The project enables TCP connections between a server and multiple clients for group chat purposes.

# Features
TCP Connection: Establish TCP connections between a server and multiple clients (1 to many).
Client Naming: Each client must provide a name upon connection.
Connection Control: Manage and limit the number of simultaneous connections (maximum 10).
Message Broadcasting: Clients can send messages to the chat, which are broadcast to all other connected clients.
Message Formatting: Messages are timestamped and identified by the client's name.
Client Join/Leave Notifications: Notify all clients when a new client joins or leaves the chat.
Message History: New clients receive the chat history upon connection.
Error Handling: Robust error handling on both server and client sides.
Default Port: Use port 8989 if no port is specified.

# Start the TCP Server:

$ go run . \n
Listening on the port :8989

$ go run . 2525 \n

Listening on the port :2525

# Client Connection:

$ nc $IP $port
