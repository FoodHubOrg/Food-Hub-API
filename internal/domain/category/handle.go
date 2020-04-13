package restaurant

import "net/http"

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindById(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

func (s *handler) Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){

}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){

}

func (s *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){

}

func (s *handler) FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){

}

func (s *handler) FindById(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){

}
