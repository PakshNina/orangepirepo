package service

import "net/http"

type srv struct {
	httpSrv *http.Server
	b       board
}

type board interface {
	MicName() string
}

func NewServer(addr string, b board) *srv {
	s := http.Server{
		Addr:                         addr,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
	}
	return &srv{
		httpSrv: &s,
		b:       b,
	}
}

func (s *srv) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.MainHandle)
	mux.HandleFunc("/mic", s.BoardMicHandle)
	s.httpSrv.Handler = mux
	return s.httpSrv.ListenAndServe()
}

func (s *srv) MainHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func (s *srv) BoardMicHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.b.MicName()))
}
