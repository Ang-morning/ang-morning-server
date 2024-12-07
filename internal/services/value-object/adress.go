package valueObject

type Address struct {
	City        string
	RoadAddress string
	Latitude    float64
	Longitude   float64
	ZipCode     string
}

func NewAddress(city, roadAddress, zipCode string, latitude, longitude float64) *Address {
	return &Address{
		City:        city,
		RoadAddress: roadAddress,
		ZipCode:     zipCode,
		Latitude:    latitude,
		Longitude:   longitude,
	}
}

type AddressUpdateStruct struct {
	City        *string
	RoadAddress *string
	Latitude    *float64
	Longitude   *float64
	ZipCode     *string
}

func (a *Address) With(update AddressUpdateStruct) *Address {
	address := *a
	if update.City != nil {
		address.City = *update.City
	}
	if update.RoadAddress != nil {
		address.RoadAddress = *update.RoadAddress
	}
	if update.Latitude != nil {
		address.Latitude = *update.Latitude
	}
	if update.Longitude != nil {
		address.Longitude = *update.Longitude
	}
	if update.ZipCode != nil {
		address.ZipCode = *update.ZipCode
	}

	return &address
}
