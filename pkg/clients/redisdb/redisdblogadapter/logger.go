package redisdblogadapter

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type rlogAdapter struct {
}

func NewRlogAdapter() rlogAdapter {
	return rlogAdapter{}
}

func (rla rlogAdapter) Printf(ctx context.Context, format string, v ...any) {

	message := fmt.Sprintf(format, v...)
	rlog.Info("Redis: " + message)
}
