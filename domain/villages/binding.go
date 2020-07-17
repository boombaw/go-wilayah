package villages

// Village struct
type Village struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DistrictID string `json:"district_id" db:"district_id"`
}
