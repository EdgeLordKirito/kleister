package schedule

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/EdgeLordKirito/wallpapersetter/package/statusserver"
)

type tickerSync struct {
	Context       context.Context
	Cancel        context.CancelFunc
	Waiter        *sync.WaitGroup
	ErrChannel    chan error
	SignalChannel chan os.Signal
}

func newContextBundle(t *tickerSync) (*statusserver.ContextBundle, error) {
	if t == nil {
		return &statusserver.ContextBundle{}, fmt.Errorf("Require non nil pointer")
	}
	result := statusserver.ContextBundle{
		Context:    t.Context,
		CancelFunc: t.Cancel,
		Waiter:     t.Waiter,
	}
	return &result, nil
}
