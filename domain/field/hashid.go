package field

import "botx/pkg/idx"

const (
	HashIDUser int64 = iota
	HashIDPost
)

type HashIDField struct {
	HashID string `binding:"required,min=1"`
	ID     int64
}

func (f *HashIDField) En(category int64) string {
	return idx.Encode(f.ID, category)
}

func (f *HashIDField) De(category int64) (int64, error) {
	return idx.Decode(f.HashID, category)
}
