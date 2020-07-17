package provinsi

// Provinsi struct
type Provinsi struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListProvinsi struct
type ListProvinsi struct {
	Provinces []Provinsi
}
