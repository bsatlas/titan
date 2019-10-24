package titan

import (
	"context"
)

// Live returns true if the Core is running correctly.
func (c *Core) Live(ctx context.Context) bool {
	return true
}
