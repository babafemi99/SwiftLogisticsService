package bikesEntity

type Bikes struct {
	BikeId            string          `json:"bike_id"`
	BikeName          string          `json:"bike_name"`
	Colour            string          `json:"colour"`
	Picture           string          `json:"picture"`
	EngineNumber      string          `json:"engine_number"`
	PlateNumber       string          `json:"plate_number"`
	NextServicingDate string          `json:"next_servicing_date"`
	Rider             *AllocatedRider `json:"rider"`
	DatePurchased     string          `json:"date_purchased"`
	DurationOfBike    string          `json:"duration_of_bike"`
	History           []BikesHistory  `json:"history"`
}

type BikesHistory struct {
	BikeHistoryId string `json:"bike_history_id"`
	BikeId        string `json:"bike_id"`
	RidersName    string `json:"riders_name"`
	Duration      string `json:"duration"`
}

type CreateBike struct {
	BikeId        string `json:"bike_id"`
	BikeName      string `json:"bike_name"`
	Colour        string `json:"colour"`
	Picture       string `json:"picture"`
	EngineNumber  string `json:"engine_number"`
	PlateNumber   string `json:"plate_number"`
	DatePurchased string `json:"date_purchased"`
}

type UpdateBikeHistory struct {
	BikeId     string `json:"bike_id"`
	RidersName string `json:"riders_name"`
	Duration   string `json:"duration"`
}

type UpdateBike struct {
	BikeName          string `json:"bike_name"`
	Colour            string `json:"colour"`
	Picture           string `json:"picture"`
	EngineNumber      string `json:"engine_number"`
	PlateNumber       string `json:"plate_number"`
	NextServicingDate string `json:"next_servicing_date"`
}

type AllocatedRider struct {
	RiderId        string `json:"rider_id"`
	Name           string `json:"rider_name"`
	ProfilePicture string `json:"profile_picture"`
}

type AssignBikeReq struct {
	RiderId string `json:"rider_id"`
	BikeId  string `json:"bike_id"`
}
