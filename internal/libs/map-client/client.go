package mapClient

type MapClient interface {
	Geocode(address string) (string, error)
}
