package sellerdb

import (
	"fmt"
	"time"

	"github.com/c-kissel/SellerMDM.git/internal/core/data/seller"
	"github.com/c-kissel/SellerMDM.git/internal/storage/db/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type sellerSQL struct {
	db *sqlx.DB
}

func NewSellerSQL(d *sqlx.DB) *sellerSQL {
	return &sellerSQL{
		db: d,
	}
}

// Gets from Seller table mapping of seller code to SKU
func (sql *sellerSQL) Get(id uuid.UUID) (*seller.SellerModel, error) {
	var data seller.SellerModel
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE id=$1`, getFields(), postgres.SELLERS_TABLE)

	err := sql.db.Get(&data, query, id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (sql *sellerSQL) Insert(data seller.SellerModel) error {
	var id uuid.UUID

	query := fmt.Sprintf(`INSERT INTO %s (id, name, ogrn, inn, city, created_at, updated_at)
							VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`, postgres.SELLERS_TABLE)

	created := time.Now()
	updated := created
	data.CreatedAt = created
	data.UpdatedAt = updated

	row := sql.db.QueryRow(query, data.ID, data.Name, data.OGRN, data.INN, data.City, created, updated)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func getFields() string {
	return "id, name, ogrn, inn, city"
}
