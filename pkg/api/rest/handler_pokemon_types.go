// Package rest is port handler.
package rest

import "github.com/kubuskotak/king/pkg/entity"

// ListPokemonRequest Get a Pokemon request.  /** PLEASE EDIT THIS EXAMPLE, request handler */.
type ListPokemonRequest struct {
	*entity.Resource
}

// ListPokemonResponse Get a Pokemon response.  /** PLEASE EDIT THIS EXAMPLE, return handler response */.
type ListPokemonResponse struct {
	*entity.Resource
}
