package server

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var welcome string = `
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    '.       | '' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     '-'       '--'
`

func (s *Server) ReadLoop(conn net.Conn) {
	defer s.ClientExit(conn)

	// Send the welcome message
	conn.Write([]byte(welcome))

	buffer := make([]byte, 1024)
	name := make([]byte, 20000)
	var trimmedName string
	nameMessage := "Enter Your Name: "
	for {
		conn.Write([]byte(nameMessage))
		n1, err := conn.Read(name)
		if err != nil {
			fmt.Print("Error in read: ", err)
			return
		}
		trimmedName = strings.TrimSpace(string(name[:n1]))
		val, _ := s.NameValidation(trimmedName)
		if val {
			break
		}
		_, nameMessage = s.NameValidation(trimmedName)
	}

	// Register the client's name
	s.mu.Lock()
	s.clients[conn] = trimmedName
	s.mu.Unlock()

	// Send a personalized welcome message to the user
	conn.Write([]byte(fmt.Sprintf("Welcome, %s! You have joined the chat.\n", trimmedName)))

	// Announce to everyone that a new user has joined
	s.msgch <- Message{msg: fmt.Sprintf("%s has joined our chat...", trimmedName), sender: "System", time: time.Now()}

	// Send the chat history to the new user
	s.mu.Lock()
	for _, msg := range s.history {
		conn.Write([]byte(fmt.Sprintf("[%s][%s]: %s\n", msg.time.Format("2006-01-02 15:04:05"), msg.sender, msg.msg)))
	}
	s.mu.Unlock()

	// Read and broadcast user messages
	for {
		n2, err := conn.Read(buffer)
		if err != nil {
			fmt.Print(err)
			break
		}
		messageContent := strings.TrimSpace(string(buffer[:n2]))
		if messageContent == "" {
			continue // Skip empty messages
		}
		message := Message{
			sender: trimmedName,
			msg:    messageContent,
			time:   time.Now(),
		}
		s.msgch <- message

		// Clear buffer for next read
		buffer = make([]byte, 1024)
	}
}

func (s *Server) BroadcastMessages() {
	for message := range s.msgch {
		s.mu.Lock()
		// Append the new message to the chat history
		s.history = append(s.history, message)
		for client := range s.clients {
			_, err := client.Write([]byte(fmt.Sprintf("[%s][%s]: %s\n", message.time.Format("2006-01-02 15:04:05"), message.sender, message.msg)))
			if err != nil {
				fmt.Println("Error sending message to client: ", err)
			}
		}
		s.mu.Unlock()
	}
}
