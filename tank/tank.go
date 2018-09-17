package tank

type Tank struct {
	id       uint8
	position uint16
	health   uint8
}

var tankIdentifier uint8

func NewTank(position uint16) *Tank {
	tankIdentifier++
	return &Tank{id: tankIdentifier, position: position, health: 100}
}

func (t *Tank) IsAlive() bool {
	return t.health > 0
}

func (t *Tank) Hit(damage uint8) {
	t.health -= damage
}

func (t *Tank) SetDegree(degree int8) {

}

func (t *Tank) SetPower(power uint8) {

}
