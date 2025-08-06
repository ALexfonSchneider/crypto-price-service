package removeCoin

import "context"

type Coins interface {
	Deactivate(ctx context.Context, symbol string) error
}
