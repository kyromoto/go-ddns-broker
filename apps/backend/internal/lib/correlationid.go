package lib

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ReadCorrelationID(ctx context.Context) (uuid.UUID, error) {
	cid, ok := ctx.Value("correlationid").(uuid.UUID)

	if !ok {
		return uuid.Nil, errors.New("failed to get correlationid from context")
	}

	return cid, nil
}

func WriteCorrelationID(ctx context.Context, correlationID uuid.UUID) context.Context {
	return context.WithValue(ctx, "correlationid", correlationID)
}

func CreateContextWithCorrelationIDFromRequest(c *fiber.Ctx) context.Context {
	requestIDAsStr, ok := c.Locals("requestid").(string)
	// requestID, ok := requestidAsStr.(string)

	if !ok {
		log.Error().Msg("read request id from request failed")
		requestIDAsStr = uuid.Nil.String()
	}

	correlationID, err := uuid.Parse(requestIDAsStr)

	if err != nil {
		log.Error().Err(err).Msg("parse request id to correlationid failed")
		correlationID = uuid.Nil
	}

	return context.WithValue(context.Background(), "correlationid", correlationID)
}

func LoggerWithCorrelationID(ctx context.Context) zerolog.Logger {
	cid, err := ReadCorrelationID(ctx)

	if err != nil {
		return log.With().Err(errors.New("failed to get correlationid from context")).Logger()
	}

	return log.With().Str("correlationid", cid.String()).Logger()
}
