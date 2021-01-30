package main

import (
	"math"
)

// y = f(x) 형태로 수학에서 쓰는 함수형
type Func func(float64) float64

// Func을 받아서 Func를 돌려주는 함수 , 함수를 다른 함수로 변환하는 함수의 형태
type Transform func(Func) Func

const tolerance = 0.00001
const dx = 0.00001

func Square(x float64) float64 {
	return x * x
}

// 어떤 함수 f를 계속해서 적용했을때 어떤 값으로 수렴하는 경우에 그 수렴 값을 찾는 함수
func FixedPoint(f Func, firstGuess float64) float64 {
	// 두수 v1,v2가 서로 tolerance 이하로 가까워졋으면 참을 돌려준다.
	closeEnough := func(v1, v2, float64) bool {
		return math.Abs(v1-v2) < tolerance
	}

	var try Func
	try = func(guess float64) float64 {
		next := f(guess)
		if closeEnough(guess, next) {
			return next
		} else {
			return try(next)
		}
	}
	return try(firstGuess)
}

func FixedPointOfTransform(g Func, transform Transform, guess float64) float64 {
	return FixedPoint(transform(g), guess)
}

//어떤 함수를 받아서 다른 함수를 돌려주는 형태 각 점에서 기울기를 구한다, 이 함수는 g 함수를 받아서 미분된 함수로 돌려주는 함수
func Deriv(g Func) Func {
	return func(x float64) float64 {
		return (g(x+dx)- g(x))/dx
	}
}

// 뉴턴방법은 f(x) = 0 이 되는 x를 찾는 방법 중 하나이다 g를 받아서 반복적으로 적용하면 g(x) = 0 에 점점 다가가는 x 값을 반환하는 함수를 돌려준다.
func NewtonTransform(g Func) Func {
	return func(x float64) float64 {
		return x - (g(x)/Deriv(g)(x))
	}
}
// ??? 디스커버리 Go 136쪽 이게 뭐임';;;
func Sqrt(x float64) float64 {
	return FixedPointOfTransform(func(y float64) float64) {
		return Square(y)-x
	},NewtonTransform, 1.0)
}