package main

// House ...
//
type House struct {
	count     int
	isEndZone bool
	isNorth   bool
	next      *House
	opposite  *House
}

func newHouse(count int, isEndZone bool, isNorth bool) *House {
	p := new(House)
	p.count = count
	p.isEndZone = isEndZone
	p.isNorth = isNorth
	p.next = nil
	p.opposite = nil
	return p
}
