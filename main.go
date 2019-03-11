package goarea

import "math"

//Pi é uma proportção numerica...
const Pi = 3.1416

//comentario...
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// retc...
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visivel
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
