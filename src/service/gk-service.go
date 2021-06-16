package service

import (
	"context"
	"errors"
	"cryptoserver-clean-app/data"
	"cryptoserver-clean-app/model"

)

/*Service ... */
type Service interface {
	GetSymbol(ctx context.Context, symbol string) (model.GetSymbolResponseBody, error)
}

/*CryptoServerSvc ...*/
type CryptoServerSvc struct {
	Config      *data.Config
}

/*NewCryptoServerSvc ...*/
func NewCryptoServerSvc(config *data.Config) Service {
	return &CryptoServerSvc{Config: config}

}

var (
	//ErrNotFound - Resource Not Found
	ErrNotFound = errors.New("Resource Not found")
	//ErrInternal - Internal Server Error
	ErrInternal = errors.New("Internal server error")
	//ErrBadRequest - Bad Request
	ErrBadRequest = errors.New("Bad request")
	//ErrRecordNoFound - Record Not Found
	ErrRecordNoFound = errors.New("Record Not found")
)
