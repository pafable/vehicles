package plane

type Plane struct {
	specs map[string]string
	mach  float32
}

func (p *Plane) GetMfr() string {
	return p.specs["mfr"]
}

func (p *Plane) SetDetail(key string, value string) {
	p.specs[key] = value
}

func (p *Plane) GetSpec(query string) string {
	return p.specs[query]
}

func (p *Plane) GetMach() float32 {
	return p.mach
}
