package seller

import (
	"time"

	"github.com/google/uuid"
)

type Seller struct {
	ID          uuid.UUID `db:"id"`
	Contact     uuid.UUID `db:"contact"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	ImageNames  []string  `db:"imagenames"`
	OGRN        string    `db:"ogrn"`
	INN         string    `db:"inn"`
	City        string    `db:"city"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
