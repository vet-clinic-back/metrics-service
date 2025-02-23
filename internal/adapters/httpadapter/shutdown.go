package httpadapter

import (
	"context"
)

func (s *HTTPAdapter) Shutdown(ctx context.Context) error {
	return s.mainServer.Shutdown(ctx)
}
