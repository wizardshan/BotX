package field

type IDField struct {
	ID int64 `binding:"required,min=1"`
}

//func (field *IDField) Scan(val any) error {
//	switch v := val.(type) {
//	case int64:
//		field.ID = v
//		return nil
//	}
//
//	return errors.New("IDField scan error")
//}
//
//func (field *IDField) MarshalText() ([]byte, error) {
//	if field == nil {
//		return []byte("<nil>"), nil
//	}
//	return numberx.ToBytes(field.ID)
//}
//
//func (field *IDField) UnmarshalText(text []byte) (err error) {
//	field.ID, err = binaryx.ToInt(text)
//	return
//}
