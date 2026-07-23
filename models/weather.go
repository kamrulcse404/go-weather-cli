package models

type Main struct {
	Temp     float64 `json:"temp"`
	Humidity int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type WeatherResponse struct {
	Name    string    `json:"name"`
	Main    Main      `json:"main"`
	Wind    Wind      `json:"wind"`
	Weather []Weather `json:"weather"`
}
