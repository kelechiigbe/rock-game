package models

import (
	"fmt"
	"log"
	"math"

	"github.com/kelechiigbe/rock-game/pkg/utils"
)

const (
	ArmRock     = "rock"
	ArmPaper    = "paper"
	ArmScissors = "scissors"
	ArmLizard   = "lizard"
	ArmSpock    = "spock"
)

type Choice struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PlayerChoice struct {
	Player int `json:"player"`
}

type GameResult struct {
	Results  string `json:"results"`
	Player   int    `json:"player"`
	Computer int    `json:"computer"`
}
type Rock string
type Paper string
type Scissor string
type Lizard string
type Spock string

type Arm interface {
	Defeats(val string) bool
}

func (r *Rock) Defeats(val string) bool {
	return val == ArmScissors || val == ArmLizard
}

func (r *Paper) Defeats(val string) bool {
	return val == ArmRock || val == ArmSpock
}

func (r *Scissor) Defeats(val string) bool {
	return val == ArmPaper || val == ArmLizard
}

func (r *Lizard) Defeats(val string) bool {
	return val == ArmPaper || val == ArmSpock
}

func (r *Spock) Defeats(val string) bool {
	return val == ArmScissors || val == ArmRock
}

func getArm(arm string) (Arm, error) {
	switch arm {
	case ArmRock:
		r := Rock(arm)
		return &r, nil
	case ArmPaper:
		p := Paper(arm)
		return &p, nil
	case ArmScissors:
		s := Scissor(arm)
		return &s, nil
	case ArmLizard:
		l := Lizard(arm)
		return &l, nil
	case ArmSpock:
		s := Spock(arm)
		return &s, nil
	}

	return nil, fmt.Errorf("invalid value '%v' to get as arm", arm)
}

var choices = []Choice{{Id: 1, Name: ArmRock}, {Id: 2, Name: ArmPaper}, {Id: 3, Name: ArmScissors}, {Id: 4, Name: ArmLizard}, {Id: 5, Name: ArmSpock}}

func GetAllChoices() []Choice {
	return choices
}

func GetChoice() Choice {
	var i = float64((utils.GetRandomNumber().RandomNumber)) / 20
	index := math.Ceil(i)
	return choices[int(index)]
}

func GetChoiceById(choice_id int) Choice {

	c := findChoiceByIndexId(choice_id)
	if c == -1 {
		log.Fatal("Invalid id")
	}

	return choices[c]
}

func PlayGame(myChoice PlayerChoice) GameResult {
	player1 := GetChoiceById(myChoice.Player)
	player2 := GetChoice()
	var gameResult = GameResult{Results: getResult(player1, player2), Player: player1.Id, Computer: player2.Id}

	return gameResult
}

func getResult(player1 Choice, player2 Choice) string {

	lh, err := getArm(player1.Name)
	if err != nil {
		log.Fatal(err)
	}

	if player1.Name == player2.Name {
		return "tie"
	} else if lh.Defeats(player2.Name) {
		return "win"
	} else {
		return "lose"
	}
}

func findChoiceByIndexId(choice_id int) int {
	for i, v := range choices {
		if v.Id == choice_id {
			return i
		}
	}
	return -1
}
