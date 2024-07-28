package server

// NameValidation checks if the provided name is valid and unique among the connected clients.
func (s *Server) NameValidation(name string) (bool, string) {
	// Check if the name is already taken by another client
	for _, v := range s.clients {
		if v == name {
			return false, "The Name Is Already Taken\nEnter Your Name: "
		}
	}
	// Check if the name is empty
	if name == "" {
		return false, "The Name Is invalid\nEnter Your Name: "
	}
	// Check if the name exceeds the maximum allowed length
	if len(name) > 30 {
		return false, "The Name Is Too Long\nEnter Your Name: "
	}
	// If all checks pass, the name is valid
	return true, ""
}
