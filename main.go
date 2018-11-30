package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type outcome int

const (
	outcomeNoDecision outcome = iota
	outcomeNorthWins  outcome = iota
	outcomeSouthWins  outcome = iota
	outcomeTie        outcome = iota
)

func evaluate(s *State) outcome {
	var northTotal = 0
	var southTotal = 0
	var i = 0
	for i < s.houseCount {
		northTotal += s.northHouses[i].count
		southTotal += s.southHouses[i].count
		i++
	}
	if (0 == northTotal) || (0 == southTotal) {
		var northScore = s.northEndZone.count
		var southScore = s.southEndZone.count
		if northScore > southScore {
			return outcomeNorthWins
		}
		if southScore > northScore {
			return outcomeSouthWins
		}
		return outcomeTie
	}
	return outcomeNoDecision
}

func humanTurn(s *State, isNorth bool) *State {
	var side = "South"
	if isNorth {
		side = "North"
	}
	fmt.Println("Human is playing " + side)
	var move = 0
	fmt.Scanf("%d", &move)
	return applyMove(s, move, isNorth)
}

func computerTurn(s *State, isNorth bool) *State {
	var side = "South"
	if isNorth {
		side = "North"
	}
	fmt.Println("Computer is playing " + side)
	var move = analyze(s, isNorth)
	fmt.Println("Computer plays " + strconv.Itoa(move))
	return applyMove(s, move, isNorth)
}

func main() {
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)
	var player = randGen.Intn(2) == 0
	var humanIsNorth = randGen.Intn(2) == 0
	s := newState(6, 4)

	if player {
		fmt.Print("Human")
	} else {
		fmt.Print("Computer")
	}
	fmt.Println(" goes first")
	fmt.Println(renderText(s))

	for {
		for {
			if player {
				s = humanTurn(s, humanIsNorth)
			} else {
				s = computerTurn(s, !humanIsNorth)
			}
			fmt.Println(renderText(s))
			if (evaluate(s) != outcomeNoDecision) || (!s.anotherTurn) {
				break
			}
		}
		player = !player
		if evaluate(s) != outcomeNoDecision {
			break
		}
	}
	fmt.Println("----- Game Over")
	fmt.Println(renderText(s))
	switch evaluate(s) {
	case outcomeNorthWins:
		fmt.Println("North Wins")
	case outcomeSouthWins:
		fmt.Println("South Wins")
	case outcomeTie:
		fmt.Println("Tie")
	}
}
