package plane

type Lockheed struct {
	Plane
}

func newLockheed(mach float32) IPlane {
	return &Lockheed{
		Plane: Plane{
			specs: map[string]string{
				"mfr":    "lockheed martin",
				"origin": "US",
			},
			mach: mach,
		},
	}
}
