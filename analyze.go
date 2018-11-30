package main

func analyze(s *State, isNorth bool) int {
	var newState = make([]*State, s.houseCount)
	var currentHouses = s.southHouses
	if isNorth {
		currentHouses = s.northHouses
	}

	// Attempt all possible moves - If a move is not allowed, set the newState to null
	//
	var i = 0
	for i < s.houseCount {
		if currentHouses[i].count > 0 {
			newState[i] = applyMove(s, i, isNorth)
		} else {
			newState[i] = nil
		}
		i++
	}

	// Basic strategy:
	//   If a move will result in another turn, take that move
	//   Otherwise take the move that will result in the largest tally in the end zone
	//
	var currentHouse = s.houseCount - 1
	var highestEz = 0
	var highestHouse = 0
	for currentHouse >= 0 {
		var candidateState = newState[currentHouse]
		if candidateState != nil {
			if candidateState.anotherTurn {
				return currentHouse
			}
			var testEz = candidateState.southEndZone.count
			if isNorth {
				testEz = candidateState.northEndZone.count
			}
			if testEz > highestEz {
				highestEz = testEz
				highestHouse = currentHouse
			}
		}
		currentHouse--
	}

	return highestHouse
}
