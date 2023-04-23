package seller

import (
	"time"

	"github.com/google/uuid"
)

type SellerModel struct {
	City      *string   `db:"city"`
	CreatedAt time.Time `db:"created_at"`
	Id        uuid.UUID `db:"id"`
	INN       *string   `db:"inn"`
	Logo      *string   `db:"logo"`
	Memo      *string   `db:"memo"`
	Name      string    `db:"name"`
	OGRN      *string   `db:"ogrn"`
	Site      *string   `db:"site"`
	UpdatedAt time.Time `db:"updated_at"`
	YML       *string   `db:"yml"`
}
