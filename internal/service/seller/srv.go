package seller

import (
	"github.com/c-kissel/SellerMDM.git/internal/core/data/seller"
	"github.com/c-kissel/SellerMDM.git/internal/core/errs"
	"github.com/c-kissel/SellerMDM.git/specs"
	"github.com/google/uuid"
)

type sellerSrv struct {
	SellerStorage
}

func NewSellerSrv(s *SellerStorage) *sellerSrv {
	return &sellerSrv{
		SellerStorage: *s,
	}
}

type SellerStorage interface {
	Get(id uuid.UUID) (*seller.SellerModel, error)
	All() ([]seller.SellerModel, error)
	Search(name string) ([]seller.SellerModel, error)
	Insert(data seller.SellerModel) error
	Update(data seller.SellerModel) error
	Delete(id uuid.UUID) error
}

func (s *sellerSrv) Get(id uuid.UUID) (specs.SellerResponse, error) {
	var result specs.SellerResponse

	data, err := s.SellerStorage.Get(id)
	if err != nil {
		return result, err
	}

	result = data.ToSpecs()
	return result, nil
}

// Creates seller with new ID in DB
func (s *sellerSrv) Create(newSeller specs.NewSellerRequest) (specs.SellerResponse, error) {
	data := seller.FromNewRequest(newSeller)

	// Check for duplicates
	_, err := s.GetByName(data.Name)
	if err != errs.ErrNotFound {
		return specs.SellerResponse{}, errs.ErrDuplicateNotAllowed
	}

	// Create new Id for seller

	data.Id = uuid.New()
	err = s.Insert(data)
	if err != nil {
		return specs.SellerResponse{}, err
	}

	created, err := s.Get(data.Id)
	if err != nil {
		return specs.SellerResponse{}, err
	}

	return created, nil
}

// Updates existing seller data
func (s *sellerSrv) Update(id uuid.UUID, sellerRequest specs.EditSellerRequest) (specs.SellerResponse, error) {
	data := seller.FromEditRequest(sellerRequest)

	data.Id = id
	err := s.SellerStorage.Update(data)
	if err != nil {
		return specs.SellerResponse{}, err
	}

	created, err := s.Get(data.Id)
	if err != nil {
		return specs.SellerResponse{}, err
	}

	return created, nil
}

// Returns all sellers from DB
func (s *sellerSrv) GetAll() ([]specs.SellerResponse, error) {
	data, err := s.All()
	if err != nil {
		return nil, err
	}

	result := make([]specs.SellerResponse, len(data))
	for i, item := range data {
		result[i] = item.ToSpecs()
	}

	return result, nil
}

// Searches for all sellers by name, returns array of found sellers.
func (s *sellerSrv) GetByName(name string) ([]specs.SellerResponse, error) {
	data, err := s.Search(name)
	if err != nil {
		return nil, err
	}

	result := make([]specs.SellerResponse, len(data))
	for i, item := range data {
		result[i] = item.ToSpecs()
	}

	return result, nil
}

func (s *sellerSrv) Delete(id uuid.UUID) error {
	err := s.SellerStorage.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
