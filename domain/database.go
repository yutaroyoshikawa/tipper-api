package domain

import (
	"time"

	"google.golang.org/genproto/googleapis/type/latlng"
)

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
	ArtistId     string         `json:"artistId"`
	Comments     []*Comment     `json:"comments"`
	Discription  string         `json:"discription"`
	Finish       *time.Time     `json:"finish"`
	GeoLocate    *latlng.LatLng `json:"geoLocate"`
	Name         string         `json:"name"`
	PlaceId      string         `json:"placeId"`
	PlaceName    string         `json:"placeName"`
	PostalCode   string         `json:"postalCode"`
	Start        *time.Time     `json:"start"`
	Thumbnail    string         `json:"thumbnail"`
	Tipping      []*Tipping     `json:"tipping"`
	TippingToken string         `json:"tippingToken"`
}

type User struct {
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	IconUrl  string `json:"icon_url"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}
