package data

//LocationModel type convert in tile38
type LocationModel struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
