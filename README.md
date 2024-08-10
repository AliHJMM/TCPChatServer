# NETCAT

# 💻 Tech Stack:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)


## DESCRIPTION
TCPChatServer in Go is a command-line application that replicates the functionality of NetCat (nc) for creating a server-client group chat. The server can handle multiple clients simultaneously, allowing them to communicate in real-time. Each client must provide a name to join the chat, and messages are timestamped and identified by the sender's name.

### AUTHORS
- [Ali Hasan](https://github.com/AliHJMM)
- [Habib Mansoor](https://github.com/7abib04)
- [Husain Ali](https://github.com/hujaafar)


## Usage

### How to Run
1. git clone repo-url or download zip file.
2. go run ip/localhost port 
3. nc ip port from another terminal "client"
4. Test the TCP Chat Application


### Implementation Details

#### Algorithm

# TCPChatServer: Chat Application

**TCPChatServer** is a chat application built using Go, designed to emulate the `netcat` server-client model. This application allows users to connect as clients to a server over a specified port, facilitating real-time communication among multiple participants.

## Project Features

- **TCP Communication**: Facilitates connection between a server and multiple clients, enabling a group chat environment.

- **User Identification**: Clients are prompted to enter a unique name, adding a personalized touch to the chat.

- **Client Management**: The server can control the number of active connections, ensuring optimal performance and management.

- **Message Broadcasting**: Messages sent by one client are shared with all other clients, creating a shared communication space.

- **Timestamps and User Tags**: Each message includes the sending time and the sender's name, making it easy to track conversations.

- **Persistent Chat History**: New clients receive the complete chat history upon joining the chat.

- **Connection Notifications**: The server notifies all clients when someone joins or leaves the chat.

- **Customizable Port**: The server defaults to port 8989, but this can be changed as needed.

## Technical Implementation

- **Built with Go**: The application is written in Golang, leveraging its powerful concurrency features.

- **Concurrency Management**: Uses goroutines for handling multiple client connections concurrently.

- **Synchronization**: Employs channels and mutexes to ensure safe and synchronized data handling.

- **Scalability**: Supports up to 10 clients simultaneously, balancing performance and usability.
