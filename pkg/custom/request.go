package custom

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoRequest interface {
	Bind(obj any) error
}
type customEchoRequest struct {
	ctx       echo.Context
	validator *validator.Validate
}

func (r *customEchoRequest) Bind(obj any) error {
	if err := r.ctx.Bind(obj); err != nil {
		return err
	}
	if err := r.validator.Struct(obj); err != nil {
		return err
	}
	return nil
}

var once sync.Once
var validatorInstance *validator.Validate

func NewCustomEchoRequest(echoRequest echo.Context) EchoRequest {
	once.Do(func() {
		validatorInstance = validator.New()
	})
	return &customEchoRequest{
		ctx:       echoRequest,
		validator: validatorInstance,
	}
}
