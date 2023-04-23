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
}

func (s *sellerSrv) Get(id uuid.UUID) (specs.Seller, error) {
	var result specs.Seller

	data, err := s.SellerStorage.Get(id)
	if err != nil {
		return result, err
	}

	result = data.ToSpecs()
	return result, nil
}

// Creates seller with new ID in DB
// TODO: check seller name for duplicates
func (s *sellerSrv) Create(newSeller specs.Seller) (specs.Seller, error) {
	data := seller.FromSpecs(newSeller)

	// Check for duplicates
	_, err := s.GetByName(data.Name)
	if err != errs.ErrNotFound {
		return specs.Seller{}, errs.ErrDuplicateNotAllowed
	}

	err = s.Insert(data)
	if err != nil {
		return specs.Seller{}, err
	}

	created, err := s.Get(data.ID)
	if err != nil {
		return specs.Seller{}, err
	}

	return created, nil
}

func (s *sellerSrv) GetAll() ([]specs.Seller, error) {
	data, err := s.All()
	if err != nil {
		return nil, err
	}

	result := make([]specs.Seller, len(data))
	for i, item := range data {
		result[i] = item.ToSpecs()
	}

	return result, nil
}

// Searches for all sellers by name, returns array of found sellers.
func (s *sellerSrv) GetByName(name string) ([]specs.Seller, error) {
	data, err := s.Search(name)
	if err != nil {
		return nil, err
	}

	result := make([]specs.Seller, len(data))
	for i, item := range data {
		result[i] = item.ToSpecs()
	}

	return result, nil
}
