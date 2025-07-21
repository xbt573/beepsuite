package beep

import (
	"errors"
	"log/slog"

	"github.com/xbt573/beepsuite/pkg/models"
)

var (
	ErrUndefinedBehavior = errors.New("undefined beep behavior")
)

func Beep(commands []models.Beep) error {
	for _, x := range commands {
		if x.Delay != 0 {
			slog.Info("beep stub", "operation", "delay", "delay", x.Delay)
			continue
		}

		if x.Frequency > 20000 {
			return ErrUndefinedBehavior
		}

		slog.Info("beep stub", "operation", "beep", "frequency", x.Frequency, "length", x.Length)
	}

	return nil
}
