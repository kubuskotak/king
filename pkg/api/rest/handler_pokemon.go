// Package rest is port handler.
package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	pkgRest "github.com/kubuskotak/asgard/rest"
	pkgTracer "github.com/kubuskotak/asgard/tracer"

	"github.com/kubuskotak/king/pkg/usecase/pokemon"
)

// PokemonOption is a struct holding the handler options.
type PokemonOption func(Pokemon *Pokemon)

// Pokemon handler instance data.
type Pokemon struct {
	uc pokemon.T
}

// WithPokemonUseCase option function to assign on pokemon.
func WithPokemonUseCase(pokemon pokemon.T) PokemonOption {
	return func(p *Pokemon) {
		p.uc = pokemon
	}
}

// NewPokemon creates a new Pokemon handler instance.
//
//	var PokemonHandler = rest.NewPokemon()
//
//	You can pass optional configuration options by passing a Config struct:
//
//	var adaptor = &adapters.Adapter{}
//	var PokemonHandler = rest.NewPokemon(rest.WithPokemonAdapter(adaptor))
func NewPokemon(opts ...PokemonOption) *Pokemon {
	// Create a new handler.
	var handler = &Pokemon{}

	// Assign handler options.
	for o := range opts {
		var opt = opts[o]
		opt(handler)
	}

	// Return handler.
	return handler
}

// Register is endpoint group for handler.
func (p *Pokemon) Register(router chi.Router) {
	router.Route("/pokemon", func(r chi.Router) {
		r.Get("/", pkgRest.HandlerAdapter[ListPokemonRequest](p.ListPokemon).JSON)
	})
}

// ListPokemon endpoint func. /** PLEASE EDIT THIS EXAMPLE, return handler response */.
func (p *Pokemon) ListPokemon(w http.ResponseWriter, r *http.Request) (ListPokemonResponse, error) {
	ctxSpan, span, l := pkgTracer.StartSpanLogTrace(r.Context(), "ListPokemon")
	defer span.End()

	result, err := p.uc.GetAll(ctxSpan)
	if err != nil {
		return ListPokemonResponse{}, pkgRest.ErrStatusConflict(w, r, err)
	}
	l.Info().Str("Hello", "World").Msg("this")

	return ListPokemonResponse{
		Resource: result,
	}, nil
}
