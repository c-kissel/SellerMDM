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
}

func (h *sellerHandle) GetSeller(w http.ResponseWriter, r *http.Request, id string) {
	api := "GetSeller (/v1/sellers/id/{id})"
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
