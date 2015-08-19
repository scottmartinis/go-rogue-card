package main

type State int

// Overall game state machine
const (
  var GSStart State = iota // First state
  GSPickEncounter
  GSDrawInitialHand
  GSDrawCard
  GSDrawItem
  GSDrawMonster
  GSFleeMonster
  GSFightMonster
  GSEncounterFinish
)

// Combat state machine
const (
  var GSPlayerStart State = iota + 100
  GSPlayerPlayCreature
  GSPlayerPlayItem
  GSPlayerAttack
  GSPlayerDead
  GSPlayerMonsterDead
  GSPlayerCapture
  GSPlayerDiscard
)

type ActionType int

const (
  ACPlayCard = iota
  ACAttack
)

type Action interface {
  func Apply(Board)
  func GetType() ActionType
  func GetArgs() interface{}
}

type StateFunc func (Board, Action) int

var gameStateFuncMap map[int]StateFunc
gameStateFuncMap := make(map[int]StateFunc)

gameStateFuncMap[GSStart] = func (b Board, a Action) int {
  return GSPlayerStart
}

gameStateFuncMap[GSPlayerStart] = func (b Board, a Action) int {

}
