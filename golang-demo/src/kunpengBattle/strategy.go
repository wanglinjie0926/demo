package kunpengBattle

type Strategy interface {
	Registrate(registration KunPengRegistration) error
	LegStart(legStart KunPengLegStart) error
	LegEnd(legEnd KunPengLegEnd) error
	React(round KunPengRound) (KunPengAction, error)
}
