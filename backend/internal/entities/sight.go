package entities

type Sight struct {
	ID          int     `json:"id"`
	Rating      float32 `json:"rating"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CityID      int     `json:"cityID"`
	CountryID   int     `json:"countryID"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Path        string  `json:"url"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
}

func (h Sight) Validate() error {
	return nil
}

type Sights struct {
	Sight []Sight `json:"sights"`
}
