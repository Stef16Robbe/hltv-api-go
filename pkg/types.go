package hltv

// import "github.com/alexflint/go-restructure"

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
	Removed Veto = iota
	Picked
)

type MapsStat struct {
	Map    Map
	Wins   int
	Draws  int
	Losses int
}

type VetoStat struct {
	Team    string
	PickBan Veto
	Map     Map
}
