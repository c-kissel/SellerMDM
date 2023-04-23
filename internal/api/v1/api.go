package v1

import (
	"net/http"

	"github.com/c-kissel/SellerMDM.git/internal/api/v1/seller"
	"github.com/c-kissel/SellerMDM.git/specs"
	"github.com/google/uuid"
)

// Quick check if api specification valid
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
	SellerHandler
}

type SellerHandler interface {
	GetSeller(w http.ResponseWriter, r *http.Request, id string)
	GetSellersAll(w http.ResponseWriter, r *http.Request)
	GetSellersByName(w http.ResponseWriter, r *http.Request, params specs.GetSellersByNameParams)
	PostNewSeller(w http.ResponseWriter, r *http.Request)
	PutSeller(w http.ResponseWriter, r *http.Request, id uuid.UUID)
	DeleteSeller(w http.ResponseWriter, r *http.Request, id uuid.UUID)
}

type Server interface {
	seller.SellerServer
}

func NewAPI(s *Server) *apiServer {
	sell := (*s).(seller.SellerServer)

	return &apiServer{
		SellerHandler: seller.NewSellerHandler(&sell),
	}
}
