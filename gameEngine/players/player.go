package players

type player struct {
	name   string
	marker byte
}

func (p *player) GetName() string {
	return p.name
}

func (p *player) GetMarker() byte {
	return p.marker
}

func (p *player) SetName(name string) (bool, error) {
	p.name = name
	return true, nil
}
func (p *player) SetMarker(marker byte) (bool, error) {
	p.marker = marker
	return true, nil
}
