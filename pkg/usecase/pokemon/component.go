// Package pokemon is implements component logic.
package pokemon

import (
	"context"
	"reflect"

	"github.com/kubuskotak/king/pkg/adapters"
	"github.com/kubuskotak/king/pkg/entity"
	"github.com/kubuskotak/king/pkg/usecase"
)

func init() {
	usecase.Register(usecase.Registration{
		Name: "pokemon",
		Inf:  reflect.TypeOf((*T)(nil)).Elem(),
		New: func() any {
			return &impl{}
		},
	})
}

// T is the interface implemented by all pokemon Component implementations.
type T interface {
	GetAll(ctx context.Context) (*entity.Resource, error)
}

type impl struct {
	adapter *adapters.Adapter
}

// Init initializes the execution of a process involved in a pokemon Component usecase.
func (i *impl) Init(adapter *adapters.Adapter) error {
	i.adapter = adapter
	return nil
}
