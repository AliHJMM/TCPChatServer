package server

func (s *Server) NameValidation(name string) (bool, string) {
	for _, v := range s.clients {
		if v == name {
			return false, "The Name Is Already Taken\nEnter Your Name: "
		}
	}
	if name == "" {
		return false, "The Name Is invalid\nEnter Your Name: "
	}
	if len(name) > 30 {
		return false, "The Name Is Too Long\nEnter Your Name: "
	}
	return true, ""
}
