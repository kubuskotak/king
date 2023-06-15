// Package rest is port handler.
package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	pkgRest "github.com/kubuskotak/asgard/rest"
	pkgTracer "github.com/kubuskotak/asgard/tracer"
)

// Hello handler instance data.
type Hello struct {
}

// Register is endpoint group for handler.
func (h *Hello) Register(router chi.Router) {
	// PLEASE EDIT THIS EXAMPLE, how to register handler to router
	router.Get("/hello", pkgRest.HandlerAdapter(h.Hello).JSON)
	router.Get("/hello-csv", pkgRest.HandlerAdapter(h.HelloCSV).CSV)
}

// ResponseHello Hello handler response. /** PLEASE EDIT THIS EXAMPLE, return handler response */.
type ResponseHello struct {
	Message string
}

// Hello endpoint func. /** PLEASE EDIT THIS EXAMPLE, return handler response */.
func (h *Hello) Hello(w http.ResponseWriter, r *http.Request) (ResponseHello, error) {
	_, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "Hello")
	defer span.End()

	l.Info().Str("Hello", "World").Msg("this")

	return ResponseHello{
		Message: "Hello everybody",
	}, nil
}

// HelloCSV endpoint func. /** PLEASE EDIT THIS EXAMPLE, return handler response */.
func (h *Hello) HelloCSV(w http.ResponseWriter, r *http.Request) (pkgRest.ResponseCSV, error) {
	_, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "HelloCSV")
	defer span.End()

	l.Info().Str("Hello", "World").Msg("this")

	rows := make([][]string, 0)
	rows = append(rows, []string{"SO Number", "Nama Warung", "Area", "Fleet Number", "Jarak Warehouse", "Urutan"})
	rows = append(rows, []string{"SO45678", "WPD00011", "Jakarta Selatan", "1", "45.00", "1"})
	rows = append(rows, []string{"SO45645", "WPD001123", "Jakarta Selatan", "1", "43.00", "2"})
	rows = append(rows, []string{"SO45645", "WPD003343", "Jakarta Selatan", "1", "43.00", "3"})
	return pkgRest.ResponseCSV{
		Filename: "warehouse",
		Rows:     rows,
	}, nil
}
