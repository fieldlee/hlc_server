package server

type Server struct {
	Code	string
	Msg	string
}

func (this *Server)Error() string {
	return this.Msg
}
