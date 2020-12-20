package main

type Player struct {
}

type Monster struct {
}

func (p *Player) Attack(target BeAttackable) {
	dmg := 1
	target.beAttacked(dmg)
}

type Attackable interface {
	Attack(BeAttackable)
}

type BeAttackable interface {
	beAttacked(int)
}

func Attack(attacker Attackable, defender BeAttackable) {
	attacker.Attack(defender)
}
