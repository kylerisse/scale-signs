package server

type rooms map[string][]sponsors

func newRooms(schedule *Schedule) *rooms {
	rms := make(rooms)
	for _, p := range schedule.Presentations {
		rn := p.Location
		_, ok := rms[rn]
		if !ok {
			var s []sponsors
			rms[rn] = s
		}
	}
	return &rms
}
