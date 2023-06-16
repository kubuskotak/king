// Package rest is port handler.
package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/copier"
	pkgRest "github.com/kubuskotak/asgard/rest"
	pkgTracer "github.com/kubuskotak/asgard/tracer"

	"github.com/kubuskotak/king/pkg/entity"
	"github.com/kubuskotak/king/pkg/persist/crud"
	"github.com/kubuskotak/king/pkg/persist/crud/ent"
	"github.com/kubuskotak/king/pkg/persist/crud/ent/hello"
)

// Hello handler instance data.
type Hello struct {
	*crud.Database
}

// Register is endpoint group for handler.
func (h *Hello) Register(router chi.Router) {
	router.Route("/hello", func(r chi.Router) {
		r.Get("/", pkgRest.HandlerAdapter(h.HelloAll).JSON)
		r.Post("/", pkgRest.HandlerAdapter(h.HelloPost).JSON)

		r.Route("/{id:[0-9-]+}", func(id chi.Router) {
			id.Get("/", pkgRest.HandlerAdapter(h.HelloGet).JSON)
			id.Put("/", pkgRest.HandlerAdapter(h.HelloPut).JSON)
			id.Delete("/", pkgRest.HandlerAdapter(h.HelloDelete).JSON)
		})
	})
}

// RequestHelloAll handler request .
type RequestHelloAll struct {
	Title string
	Limit int `validate:"gte=0,default=10"`
	Page  int `validate:"gte=0,default=1"`
}

// ResponseHelloAll handler response.
type ResponseHelloAll struct {
	Hellos []*entity.Hello `json:"hellos"`
}

// HelloAll [GET /] hello endpoint func.
func (h *Hello) HelloAll(w http.ResponseWriter, r *http.Request) (resp ResponseHelloAll, err error) {
	ctxSpan, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "HelloAll")
	defer span.End()
	var (
		request RequestHelloAll
		b       *pkgRest.Binder[RequestHelloAll]
	)
	b, err = pkgRest.Bind(r, &request)
	if err != nil {
		l.Error().Err(err).Msg("Bind RequestHelloAll")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}

	if err = b.Validate(); err != nil {
		l.Error().Err(err).Msg("Validate RequestHelloAll")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	var (
		total int
		query = h.Database.Hello.Query()
	)
	// pagination
	total, err = query.Count(ctxSpan)
	if err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	pkgRest.Paging(r, pkgRest.Pagination{
		Page:  request.Page,
		Limit: request.Limit,
		Total: total,
	})
	var (
		hall   []*ent.Hello
		offset = (request.Page - 1) * request.Limit
		rows   = make([]*entity.Hello, len(hall))
	)
	hall, err = query.
		Limit(request.Limit).
		Offset(offset).
		Order(ent.Desc(hello.FieldID)).
		Where(hello.TitleContains(request.Title)).
		All(ctxSpan)
	if err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	if err = copier.Copy(&rows, &hall); err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	return ResponseHelloAll{
		Hellos: rows,
	}, nil
}

// RequestHelloGet handler request .
type RequestHelloGet struct {
	entity.HelloParams
}

// ResponseHelloGet handler response.
type ResponseHelloGet struct {
	entity.Hello
}

// HelloGet [GET :id]  hello endpoint func.
func (h *Hello) HelloGet(w http.ResponseWriter, r *http.Request) (resp ResponseHelloGet, err error) {
	ctxSpan, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "Hello")
	defer span.End()
	var (
		request RequestHelloGet
		b       *pkgRest.Binder[RequestHelloGet]
	)
	b, err = pkgRest.Bind(r, &request)
	if err != nil {
		l.Error().Err(err).Msg("Bind RequestHelloGet")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	if err = b.Validate(); err != nil {
		l.Error().Err(err).Msg("Validate RequestHelloGet")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	var (
		row *ent.Hello
		hl  entity.Hello
	)
	row, err = h.Database.Hello.Query().
		Where(hello.ID(request.ID)).First(ctxSpan)
	if ent.IsNotFound(err) {
		l.Error().Err(err).Msg("IsNotFound")
		return resp, pkgRest.ErrBadRequest(w, r, fmt.Errorf("data is not found: %d", request.ID))
	}
	l.Info().Msg("RequestHelloGet")
	if err = copier.Copy(&hl, &row); err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	return ResponseHelloGet{Hello: hl}, nil
}

// RequestHelloPost handler request .
type RequestHelloPost struct {
	entity.Hello
}

// ResponseHelloPost handler response.
type ResponseHelloPost struct {
	entity.Hello
}

// HelloPost [POST :id] hello endpoint func.
func (h *Hello) HelloPost(w http.ResponseWriter, r *http.Request) (resp ResponseHelloPost, err error) {
	ctxSpan, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "HelloPost")
	defer span.End()
	var (
		request RequestHelloPost
		b       *pkgRest.Binder[RequestHelloPost]
	)
	b, err = pkgRest.Bind(r, &request)
	if err != nil {
		l.Error().Err(err).Msg("Bind HelloPost")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	if err = b.Validate(); err != nil {
		l.Error().Err(err).Msg("Validate HelloPost")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	var (
		row *ent.Hello
		hl  entity.Hello
	)
	row, err = h.Database.Hello.
		Create().
		SetTitle(request.Title).
		SetBody(request.Body).
		SetDescription(request.Slug).
		SetSlug(request.Slug).
		Save(ctxSpan)
	if err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	l.Info().Interface("Hello", hl).Msg("this")
	if err = copier.Copy(&hl, &row); err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	return ResponseHelloPost{
		Hello: hl,
	}, nil
}

// RequestHelloPut handler request .
type RequestHelloPut struct {
	entity.HelloParams
	entity.Hello
}

// ResponseHelloPut handler response.
type ResponseHelloPut struct {
	entity.Hello
}

// HelloPut [PUT :id] hello endpoint func.
func (h *Hello) HelloPut(w http.ResponseWriter, r *http.Request) (ResponseHelloPut, error) {
	_, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "Hello")
	defer span.End()

	l.Info().Str("Hello", "World").Msg("this")

	return ResponseHelloPut{}, nil
}

// RequestHelloDelete handler request .
type RequestHelloDelete struct {
	entity.HelloParams
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

	l.Info().Int("id", request.ID).Msg("HelloDelete")

	return ResponseHelloDelete{
		Message: "Hello everybody",
	}, nil
}
