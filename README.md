# TCPChatServer: Chat Application

## **Description**

TCPChatServer is a chat application built using Go, designed to emulate the `netcat` server-client model. This application allows users to connect as clients to a server over a specified port, facilitating real-time communication among multiple participants. The server can handle multiple clients simultaneously, enabling group chats with features such as user identification, message broadcasting, and timestamps.

---

## **Features**

### **Core Features**

- **TCP Communication**: Facilitates connections between a server and multiple clients for group chats.
- **User Identification**: Clients are prompted to enter a unique name when joining the chat.
- **Message Broadcasting**: Messages sent by one client are shared with all connected clients.
- **Timestamps and User Tags**: Each message includes the sender's name and timestamp for clarity.
- **Connection Notifications**: All clients are notified when someone joins or leaves the chat.

### **Enhanced Features**

- **Persistent Chat History**: New clients can view the chat history upon joining.
- **Concurrency Management**: Uses goroutines to handle multiple client connections simultaneously.
- **Scalability**: Supports up to 10 clients at a time.
- **Customizable Port**: Defaults to port 8989 but can be changed as needed.

---

## **How to Use**

### **Running the Server**

1. Clone the repository or download the project files.
   ```bash
   git clone https://github.com/AliHJMM/TCPChatServer
   ```
2. Navigate to the project directory and run the server:
   ```bash
   go run server.go <ip-address> <port>
   ```
   - Replace `<ip-address>` and `<port>` with the desired values (e.g., `localhost 8989`).

### **Connecting as a Client**

1. Open another terminal and connect as a client using NetCat:
   ```bash
   nc <ip-address> <port>
   ```
   - Replace `<ip-address>` and `<port>` with the server details.

2. Enter your username when prompted to join the chat.
3. Start sending and receiving messages in real-time.

---

## **Technical Implementation**

- **Built with Go**: The application leverages Go's concurrency features for efficient multi-client handling.
- **Concurrency Management**: Utilizes goroutines for simultaneous client connections.
- **Synchronization**: Channels and mutexes ensure thread-safe communication.

---

## **Learning Objectives**

This project provides insights into:

- **Network Programming**: Understanding TCP/IP protocols and socket programming.
- **Concurrency in Go**: Mastering goroutines, channels, and synchronization techniques.
- **Software Scalability**: Designing systems that efficiently handle multiple connections.

---

## **Authors**

- [Ali Hasan](https://github.com/AliHJMM)
- [Habib Mansoor](https://github.com/7abib04)
- [Husain Ali](https://github.com/hujaafar)

