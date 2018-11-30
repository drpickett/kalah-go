package main

import "fmt"

func renderFrame(s *State) string {
	var ret = "+----+----"
	var i = 0
	for i < s.houseCount {
		ret += "+----"
		i++
	}
	ret += "+\n"
	return ret
}

func renderHeader(s *State) string {
	var ret = "       "
	var i = s.houseCount - 1
	for i >= 0 {
		ret += fmt.Sprintf("N%1d   ", i)
		i--
	}
	ret += "\n"
	return ret
}

func renderNorth(s *State) string {
	var ret = fmt.Sprintf("| %02d ", s.northEndZone.count)
	var i = s.houseCount - 1
	for i >= 0 {
		ret += fmt.Sprintf("| %02d ", s.northHouses[i].count)
		i--
	}
	ret += "|    |\n"
	return ret
}

func renderSouth(s *State) string {
	var ret = "|    "
	var i = 0
	for i < s.houseCount {
		ret += fmt.Sprintf("| %02d ", s.southHouses[i].count)
		i++
	}
	ret += fmt.Sprintf("| %02d |\n", s.southEndZone.count)
	return ret
}

func renderFooter(s *State) string {
	var ret = "       "
	var i = 0
	for i < s.houseCount {
		ret += fmt.Sprintf("S%1d   ", i)
		i++
	}
	ret += "\n"
	return ret
}

func renderText(s *State) string {
	return renderHeader(s) + renderFrame(s) + renderNorth(s) + renderSouth(s) + renderFrame(s) + renderFooter(s)
}
