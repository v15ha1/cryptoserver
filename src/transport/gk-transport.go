package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"cryptoserver-clean-app/model"
	"cryptoserver-clean-app/service"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

//MakeHTTPHandler - Create http handlers
func MakeHTTPHandler(s service.Service, logger log.Logger) http.Handler {

	r := mux.NewRouter()

	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	sr := r.PathPrefix("/api/v1").Subrouter()

	sr.Methods("GET").Path("/currency/{symbol}").Handler(httptransport.NewServer(
		e.GetSymbolEndpoint,
		decodeSymbolRequest,
		encodeSymbolResponse,
		options...,
	))

	return r
}

func decodeSymbolRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	symbol, ok := vars["symbol"]
	if !ok {
		return nil, ErrBadRouting
	}

	if len(symbol) <= 0 {
		return nil, ErrBadRouting
	}

	return symbolRequest{
		Symbol: symbol,
	}, nil
}

func encodeSymbolResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	r, ok := response.(model.GetSymbolResponse)
	if ok && r.Err != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, r.Err, w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(r.Body)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.ErrNotFound:
		return http.StatusNotFound
	case service.ErrBadRequest:
		return http.StatusBadRequest
	case service.ErrInternal:
		return http.StatusInternalServerError
	case service.ErrRecordNoFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
