package service

import "net/http"

type srv struct {
	httpSrv *http.Server
}

func NewServer(addr string) *srv {
	s := http.Server{
		Addr:                         addr,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
	}
	return &srv{httpSrv: &s}
}

func (s *srv) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.MainHandle)
	s.httpSrv.Handler = mux
	return s.httpSrv.ListenAndServe()
}

func (s *srv) MainHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
