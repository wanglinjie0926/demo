package kunpengBattle

type KunPengMsg struct {
	Name string      `json:"msg_name"`
	Data KunPengData `json:"msg_data"`
}

type KunPengData interface{}

type KunPengRegistration struct {
	TeamID   int    `json:"team_id"`
	TeamName string `json:"team_name"`
}

type KunPengLegStart struct {
	Map   KunPengMap    `json:"map"`
	Teams []KunPengTeam `json:"teams"`
}

type KunPengLegEnd struct {
	Teams []KunPengTeam `json:"teams"`
}

type KunPengRound struct {
	ID      int             `json:"round_id"`
	Mode    string          `json:"mode"`
	Power   []KunPengPower  `json:"power"`
	Players []KunPengPlayer `json:"players"`
	Teams   []KunPengTeam   `json:"teams"`
}

type KunPengAction struct {
	ID      int           `json:"round_id"`
	Actions []KunPengMove `json:"actions"`
}

type KunPengMove struct {
	Team     int      `json:"team"`
	PlayerID int      `json:"player_id"`
	Move     []string `json:"move"`
}

type KunPengMap struct {
	Width    int               `json:"width"`
	Height   int               `json:"height"`
	Vision   int               `json:"vision"`
	Meteor   []KunPengMeteor   `json:"meteor"`
	Tunnel   []KunPengTunnel   `json:"tunnel"`
	Wormhole []KunPengWormhole `json:"wormhole"`
}

type KunPengMeteor struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type KunPengTunnel struct {
	Direction string `json:"direction"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
}

type KunPengWormhole struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type KunPengTeam struct {
	ID          int    `json:"id"`
	Players     []int  `json:"players"`
	Force       string `json:"force"`
	Point       int    `json:"point"`
	Remain_life int    `json:"remain_life"`
}

type KunPengPower struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	Point int `json:"point"`
}

type KunPengPlayer struct {
	ID    int `json:"id"`
	Score int `json:"score"`
	Sleep int `json:"sleep"`
	Team  int `json:"team"`
	X     int `json:"x"`
	Y     int `json:"y"`
}
