package v1

import (
	"net/http"

	"github.com/c-kissel/SellerMDM.git/internal/api/v1/seller"
	"github.com/c-kissel/SellerMDM.git/specs"
)

// Quick check if api specification valid
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
	SellerHandler
}

type SellerHandler interface {
	GetSeller(w http.ResponseWriter, r *http.Request, id string)
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
