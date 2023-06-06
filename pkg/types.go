package hltv

type Map int

const (
	Mirage Map = iota
	Inferno
	Nuke
	Anubis
	Ancient
	Overpass
	Vertigo
)

type Veto int

const (
	Remove Veto = iota
	Pick
)

type MapsStat struct {
	Map    Map
	Wins   int
	Draws  int
	Losses int
}

type VetoStat struct {
	Map     Map
	PickBan Veto
}
