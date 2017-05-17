package poki

import (
	"net/http"
)

type Rest struct {
	hs http.Server
}

type RestHandler struct {
	http.ServeMux
}

func NewRestHandler() {
	return
}

func NewRest(addr string) *Rest {
	return &Rest{
		hs: http.Server{
			Addr: addr,
			//Handler:
		},
	}
}

func (r Rest) Serve(addr string) error {
	return r.hs.ListenAndServe()
}
