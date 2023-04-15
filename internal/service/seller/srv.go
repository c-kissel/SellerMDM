package seller

import (
	"github.com/c-kissel/SellerMDM.git/internal/core/data/seller"
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

	err := s.Insert(data)
	if err != nil {
		return specs.Seller{}, err
	}

	created, err := s.Get(data.ID)
	if err != nil {
		return specs.Seller{}, err
	}

	return created, nil
}
