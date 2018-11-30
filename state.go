package main

// State ...
//
type State struct {
	houseCount   int
	seedCount    int
	northEndZone *House
	southEndZone *House
	anotherTurn  bool
	northHouses  []*House
	southHouses  []*House
}

func newState(houseCount int, seedCount int) *State {
	p := new(State)
	p.houseCount = houseCount
	p.seedCount = seedCount
	p.northEndZone = newHouse(0, true, true)
	p.southEndZone = newHouse(0, true, false)
	p.anotherTurn = false
	p.northHouses = make([]*House, houseCount)
	p.southHouses = make([]*House, houseCount)
	var i = 0
	for i < houseCount {
		p.northHouses[i] = newHouse(seedCount, false, true)
		p.southHouses[i] = newHouse(seedCount, false, false)
		i++
	}
	p.northEndZone.next = p.southHouses[0]
	p.southEndZone.next = p.northHouses[0]
	i = 0
	for i < houseCount {
		p.northHouses[i].opposite = p.southHouses[houseCount-i-1]
		p.southHouses[i].opposite = p.northHouses[houseCount-i-1]
		if i == houseCount-1 {
			p.northHouses[i].next = p.northEndZone
			p.southHouses[i].next = p.southEndZone
		} else {
			p.northHouses[i].next = p.northHouses[i+1]
			p.southHouses[i].next = p.southHouses[i+1]
		}
		i++
	}

	return p
}

func copyState(s *State) *State {
	p := newState(s.houseCount, s.seedCount)
	p.northEndZone.count = s.northEndZone.count
	p.southEndZone.count = s.southEndZone.count
	var i = 0
	for i < s.houseCount {
		p.northHouses[i].count = s.northHouses[i].count
		p.southHouses[i].count = s.southHouses[i].count
		i++
	}
	return p
}
