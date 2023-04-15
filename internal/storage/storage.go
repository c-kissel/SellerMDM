package storage

import (
	"github.com/c-kissel/SellerMDM.git/internal/service/seller"
	"github.com/c-kissel/SellerMDM.git/internal/storage/db/postgres/sellerdb"
	"github.com/jmoiron/sqlx"
)

type storage struct {
	seller.SellerStorage
}

func NewStorage(d *sqlx.DB) *storage {
	return &storage{
		SellerStorage: sellerdb.NewSellerSQL(d),
	}
}
