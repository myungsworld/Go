package main

// 이렇게 쓰는것보다
// type Actor interface {
// 	Move()
// 	Attack()
// 	Talk()
// }
type Actor struct {
}

// 이게 더 나음
type Talkable interface {
	Talk()
}
type Attackable interface {
	Attack()
}
type Moveable interface {
	Move()
}

func MoveTo(a Actor) {
	a.Move()
	a.Attack()
}

func MoveTo(a Actor) {
	a.Move()

}
