package titan

import (
	"context"
)

// Ready returns true if the Core is ready to serve requests.
func (c *Core) Ready(ctx context.Context) bool {
	return true
}
