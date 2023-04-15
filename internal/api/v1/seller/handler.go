package seller

import (
	"encoding/json"
	"net/http"

	"github.com/c-kissel/SellerMDM.git/internal/api/v1/httperr"
	"github.com/c-kissel/SellerMDM.git/internal/core/errs"
	"github.com/c-kissel/SellerMDM.git/specs"
	"github.com/google/uuid"
)

type sellerHandle struct {
	SellerServer
}

func NewSellerHandler(s *SellerServer) *sellerHandle {
	return &sellerHandle{
		SellerServer: *s,
	}
}

type SellerServer interface {
	Get(id uuid.UUID) (specs.Seller, error)
	Create(seller specs.Seller) (specs.Seller, error)
}

func (h *sellerHandle) GetSeller(w http.ResponseWriter, r *http.Request, id string) {
	api := "GetSeller (/api/v1/sellers/id/{id})"
	// Convert to UUID
	uuid, err := uuid.Parse(id)
	if err != nil {
		httperr.Send(w, http.StatusBadRequest, "%s: failed to decode ID: %s", api, err.Error())
		return
	}

	// Get Seller
	seller, err := h.Get(uuid)
	if err == errs.ErrNotFound {
		httperr.Send(w, http.StatusNotFound, "%s: %s", api, err.Error())
		return
	}
	if err != nil {
		httperr.Send(w, http.StatusInternalServerError, "%s: failed: %s", api, err.Error())
		return
	}

	// Convert to JSON
	data, err := json.Marshal(seller)
	if err != nil {
		httperr.Send(w, http.StatusInternalServerError, "%s: failed: %s", api, err.Error())
		return
	}

	// Return to caller
	_, err = w.Write(data)
	if err != nil {
		httperr.Send(w, http.StatusInternalServerError, "%s: failed: %s", api, err.Error())
		return
	}
}

func (h *sellerHandle) PostSeller(w http.ResponseWriter, r *http.Request) {
	api := "PostSeller (/api/v1/sellers)"

	// Get seller data from JSON
	var sellerData specs.Seller
	err := json.NewDecoder(r.Body).Decode(&sellerData)
	if err != nil {
		httperr.Send(w, http.StatusBadRequest, "%s: failed to decode JSON: %s", api, err.Error())
		return
	}

	// Create new seller
	result, err := h.Create(sellerData)
	if err != nil {
		httperr.Send(w, http.StatusBadRequest, "%s: failed to create seller: %s", api, err.Error())
		return
	}

	// Convert to JSON
	data, err := json.Marshal(result)
	if err != nil {
		httperr.Send(w, http.StatusInternalServerError, "%s: failed: %s", api, err.Error())
		return
	}

	// Return to caller
	_, err = w.Write(data)
	if err != nil {
		httperr.Send(w, http.StatusInternalServerError, "%s: failed: %s", api, err.Error())
		return
	}

}
