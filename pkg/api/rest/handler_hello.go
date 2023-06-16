// Package rest is port handler.
package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	pkgRest "github.com/kubuskotak/asgard/rest"
	pkgTracer "github.com/kubuskotak/asgard/tracer"

	"github.com/kubuskotak/king/pkg/entity"
	"github.com/kubuskotak/king/pkg/persist/crud"
)

// Hello handler instance data.
type Hello struct {
	*crud.Database
}

// Register is endpoint group for handler.
func (h *Hello) Register(router chi.Router) {
	router.Route("/hello", func(r chi.Router) {
		r.Get("/", pkgRest.HandlerAdapter(h.HelloAll).JSON)

		r.Route("/{id:[0-9-]+}", func(id chi.Router) {
			id.Get("/", pkgRest.HandlerAdapter(h.HelloGet).JSON)
			id.Post("/", pkgRest.HandlerAdapter(h.HelloPost).JSON)
			id.Put("/", pkgRest.HandlerAdapter(h.HelloPut).JSON)
			id.Delete("/", pkgRest.HandlerAdapter(h.HelloDelete).JSON)
		})
	})
}

// RequestHelloAll handler request .
type RequestHelloAll struct {
	entity.Hello
}

// ResponseHelloAll handler response.
type ResponseHelloAll struct {
	Message string
}

// HelloAll [GET /] hello endpoint func.
func (h *Hello) HelloAll(w http.ResponseWriter, r *http.Request) (ResponseHelloAll, error) {
	_, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "Hello")
	defer span.End()

	l.Info().Str("Hello", "World").Msg("this")

	return ResponseHelloAll{
		Message: "Hello everybody",
	}, nil
}

// RequestHelloGet handler request .
type RequestHelloGet struct {
	entity.Hello
}

// ResponseHelloGet handler response.
type ResponseHelloGet struct {
	Message string
}

// HelloGet [GET :id]  hello endpoint func.
func (h *Hello) HelloGet(w http.ResponseWriter, r *http.Request) (ResponseHelloGet, error) {
	_, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "Hello")
	defer span.End()

	l.Info().Str("Hello", "World").Msg("this")

	return ResponseHelloGet{
		Message: "Hello everybody",
	}, nil
}

// RequestHelloPost handler request .
type RequestHelloPost struct {
	entity.Hello
}

// ResponseHelloPost handler response.
type ResponseHelloPost struct {
	Message string
}

// HelloPost [POST :id] hello endpoint func.
func (h *Hello) HelloPost(w http.ResponseWriter, r *http.Request) (ResponseHelloPost, error) {
	_, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "Hello")
	defer span.End()

	l.Info().Str("Hello", "World").Msg("this")

	return ResponseHelloPost{
		Message: "Hello everybody",
	}, nil
}

// RequestHelloPut handler request .
type RequestHelloPut struct {
	entity.Hello
}

// ResponseHelloPut handler response.
type ResponseHelloPut struct {
	Message string
}

// HelloPut [PUT :id] hello endpoint func.
func (h *Hello) HelloPut(w http.ResponseWriter, r *http.Request) (ResponseHelloPut, error) {
	_, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "Hello")
	defer span.End()

	l.Info().Str("Hello", "World").Msg("this")

	return ResponseHelloPut{
		Message: "Hello everybody",
	}, nil
}

// RequestHelloDelete handler request .
type RequestHelloDelete struct {
	Id int `json:"id" validate:"required"`
	entity.Hello
}

// ResponseHelloDelete handler response.
type ResponseHelloDelete struct {
	Message string
}

// HelloDelete [DELETE :id] hello endpoint func.
func (h *Hello) HelloDelete(w http.ResponseWriter, r *http.Request) (resp ResponseHelloDelete, err error) {
	_, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "Hello")
	defer span.End()
	var (
		request RequestHelloDelete
		b       *pkgRest.Binder[RequestHelloDelete]
	)
	b, err = pkgRest.Bind(r, &request)
	if err != nil {
		l.Error().Err(err).Msg("Bind LoginHandler")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	if err = b.Validate(); err != nil {
		l.Error().Err(err).Msg("Validate LoginHandler")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}

	l.Info().Int("id", request.Id).Msg("HelloDelete")

	return ResponseHelloDelete{
		Message: "Hello everybody",
	}, nil
}
