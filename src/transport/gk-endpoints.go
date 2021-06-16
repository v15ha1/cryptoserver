package transport

import (
	"context"
	"cryptoserver-clean-app/model"
	"cryptoserver-clean-app/service"

	"github.com/go-kit/kit/endpoint"
)

/*Endpoints ...definition*/
type Endpoints struct {
	GetSymbolEndpoint      endpoint.Endpoint

}

/*MakeServerEndpoints ... definition*/
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetSymbolEndpoint:      MakeGetSymbolEndpoint(s),
	}
}

/*MakeGetSymbolEndpoint ... definition*/
func MakeGetSymbolEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(symbolRequest)
		g, e := s.GetSymbol(ctx, req.Symbol)
		return model.GetSymbolResponse{Body: g, Err: e}, nil
	}
}

type symbolRequest struct {
	Symbol string
}

type symbolResponse struct {
	Body model.GetSymbolResponseBody
	Err  error
}
