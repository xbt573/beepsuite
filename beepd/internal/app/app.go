package app

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/xbt573/beepsuite/beepd/internal/beep"
	"github.com/xbt573/beepsuite/pkg/models"
)

type App struct {
}

type Options struct {
}

func New() *App {
	return &App{}
}

func (a *App) Serve(ctx context.Context, address string) error {
	f := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	v1 := f.Group("/api/v1")
	v1.Post("/beep", a.beep)
	v1.Post("/beep/random", a.random)
	v1.Post("/beep/dice", a.dice)

	errch := make(chan error)

	slog.Info("running server on", "address", address)

	go func() {
		if err := f.Listen(address); err != nil {
			errch <- err
		}
	}()

	select {
	case <-ctx.Done():
		break
	case err := <-errch:
		return err
	}

	if err := f.Shutdown(); err != nil {
		return err
	}

	return nil
}

func (a *App) beep(ctx *fiber.Ctx) error {
	var commands []models.Beep

	body := ctx.Body()
	if len(body) == 0 {
		commands = []models.Beep{{Frequency: 440, Length: 200}} // populate with default
	} else {
		if err := json.Unmarshal(body, &commands); err != nil {
			return ctx.Status(400).SendString(err.Error())
		}
	}

	if err := beep.Beep(commands); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return nil
}

func (a *App) random(ctx *fiber.Ctx) error {
	param := ctx.Query("probability", "50")

	probability, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	if rand.IntN(101) <= probability {
		return a.beep(ctx)
	}

	return nil
}

func (a *App) dice(ctx *fiber.Ctx) error {
	if rand.IntN(101) <= 33 {
		return a.beep(ctx)
	}

	return nil
}
