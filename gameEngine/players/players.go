package players

type IPlayer interface {
	GetName() string
	GetMarker() byte
	SetName(name string) (bool, error)
	SetMarker(marker byte) (bool, error)
}

func InitializePlayer(name string, marker byte) IPlayer {
	return &player{
		name:   name,
		marker: marker,
	}
}
