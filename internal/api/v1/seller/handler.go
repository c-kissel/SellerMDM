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
	Get(id uuid.UUID) (specs.SellerResponse, error)
	GetAll() ([]specs.SellerResponse, error)
	GetByName(name string) ([]specs.SellerResponse, error)
	Create(newSeller specs.NewSellerRequest) (specs.SellerResponse, error)
	Update(id uuid.UUID, sellerRequest specs.EditSellerRequest) (specs.SellerResponse, error)
	Delete(id uuid.UUID) error
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

func (h *sellerHandle) GetSellersAll(w http.ResponseWriter, r *http.Request) {
	api := "GetSeller (/api/v1/sellers/id/all)"

	// Get Sellers
	sellers, err := h.GetAll()
	if err != nil {
		httperr.Send(w, http.StatusInternalServerError, "%s: failed: %s", api, err.Error())
		return
	}

	// Convert to JSON
	data, err := json.Marshal(sellers)
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

func (h *sellerHandle) GetSellersByName(w http.ResponseWriter, r *http.Request, params specs.GetSellersByNameParams) {
	api := "GetByName (/api/v1/sellers/search)"

	if params.Name == "" {
		httperr.Send(w, http.StatusBadRequest, "%s empty seller name", api)
		return
	}

	sellers, err := h.GetByName(params.Name)
	if err == errs.ErrNotFound {
		httperr.Send(w, http.StatusNotFound, "%s seller %s not found", api, params.Name)
		return
	}
	if err != nil {
		httperr.Send(w, http.StatusBadRequest, "%s failed to find sellers", api)
		return
	}

	// Convert to JSON
	data, err := json.Marshal(sellers)
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

func (h *sellerHandle) PostNewSeller(w http.ResponseWriter, r *http.Request) {
	api := "PostSeller (/api/v1/sellers)"

	// Get seller data from JSON
	var sellerData specs.NewSellerRequest
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

func (h *sellerHandle) PutSeller(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	api := "PutSeller (/api/v1/sellers/{id})"

	// Get seller data from JSON
	var sellerData specs.EditSellerRequest
	err := json.NewDecoder(r.Body).Decode(&sellerData)
	if err != nil {
		httperr.Send(w, http.StatusBadRequest, "%s: failed to decode JSON: %s", api, err.Error())
		return
	}

	// Update seller
	result, err := h.Update(id, sellerData)
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

func (h *sellerHandle) DeleteSeller(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	api := "DeleteSeller (/api/v1/sellers/{id})"

	err := h.Delete(id)
	if err != nil {
		httperr.Send(w, http.StatusInternalServerError, "%s: failed: %s", api, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
