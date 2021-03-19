package domain

type Comment struct {
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserId    string `json:"user_id"`
}

type Tipping struct {
	TippingAt string `json:"tipping_at"`
	UserId    string `json:"user_id"`
	Value     int64  `json:"value"`
}

type Performance struct {
	ArtistId     string     `json:"artist_id"`
	Comments     []*Comment `json:"comments"`
	Description  string     `json:"description"`
	Finish       string     `json:"finish"`
	GeoLocate    []float64  `json:"geo_locate"`
	Name         string     `json:"name"`
	PlaceId      string     `json:"place_id"`
	PlaceName    string     `json:"place_name"`
	PostalCode   string     `json:"postal_code"`
	Start        string     `json:"start"`
	Thumbnail    string     `json:"thumbnail"`
	Tipping      []Tipping  `json:"tipping"`
	TippingToken string     `json:"tipping_token"`
}

type User struct {
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	IconUrl  string `json:"icon_url"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
}
