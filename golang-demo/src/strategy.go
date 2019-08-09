package main

import (
	"math/rand"

	kpb "./kunpengBattle"
)

type ramdomStrategy struct {
	TeamID         int
	TeamName       string
	Allies         map[int]*kpb.KunPengPlayer
	Enemies        map[int]*kpb.KunPengPlayer
	Teams          map[int]*kpb.KunPengTeam
	Map            kpb.KunPengMap
	CurrentRoundID int
	MatrixMap      [][][]int
}

func (s *ramdomStrategy) Registrate(registration kpb.KunPengRegistration) error {
	s.TeamID = registration.TeamID
	s.TeamName = registration.TeamName
	return nil
}

func (s *ramdomStrategy) LegStart(legStart kpb.KunPengLegStart) error {
	s.Teams = make(map[int]*kpb.KunPengTeam)
	s.Allies = make(map[int]*kpb.KunPengPlayer)
	s.Enemies = make(map[int]*kpb.KunPengPlayer)
	for _, t := range legStart.Teams {
		s.Teams[t.ID] = &t

		if s.TeamID == t.ID {
			for _, playerID := range t.Players {
				s.Allies[playerID] = &kpb.KunPengPlayer{ID: playerID, Team: t.ID}
			}
		} else {
			for _, playerID := range t.Players {
				s.Enemies[playerID] = &kpb.KunPengPlayer{ID: playerID, Team: t.ID}
			}
		}

	}

	s.Map = legStart.Map
	return nil
}

func (s *ramdomStrategy) LegEnd(legEnd kpb.KunPengLegEnd) error {
	for _, t := range legEnd.Teams {
		team := s.Teams[t.ID]
		team.Point = t.Point
	}
	return nil
}

func (s *ramdomStrategy) React(round kpb.KunPengRound) (kpb.KunPengAction, error) {
	s.CurrentRoundID = round.ID

	action := new(kpb.KunPengAction)
	action.ID = s.CurrentRoundID
	action.Actions = make([]kpb.KunPengMove, 0, len(s.Allies))
	move := [5]string{"up", "down", "right", "left", ""}

	for _, player := range s.Allies {

		ac := kpb.KunPengMove{Team: s.TeamID, PlayerID: player.ID, Move: []string{move[rand.Intn(5)]}}
		action.Actions = append(action.Actions, ac)
	}
	return *action, nil

}
