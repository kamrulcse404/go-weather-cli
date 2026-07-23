package models

type City struct {
	Name string `json:"name"`
}

type ForecastItem struct {
	Main    Main     `json:"main"`
	Weather []Weather `json:"weather"`
	Wind    Wind      `json:"wind"`
	Date    string    `json:"dt_txt"`
}

type ForecastResponse struct {
	City City           `json:"city"`
	List []ForecastItem `json:"list"`
}
