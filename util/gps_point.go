package util

type IGPSPoint interface {
	Walk()
}

type GPSPointStruct struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

func NewGPSPoint(lat, lng float64)*GPSPointStruct  {
	return &GPSPointStruct{
		Latitude: lat,
		Longitude: lng,
	}
}

func (G *GPSPointStruct) Walk(strip float64)*GPSPointStruct {
	return NewGPSPoint(G.Latitude + RandomFloat(-strip, strip), G.Longitude + RandomFloat(-strip, strip))
}

