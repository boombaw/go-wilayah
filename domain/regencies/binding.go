package regencies

// Regency struct
type Regency struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProvinceID string `json:"province_id" db:"province_id"`
}
