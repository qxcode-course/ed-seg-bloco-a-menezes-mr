package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func circle(t *Pen, raio float64) {
	if raio < 3 {
		return 
	}

	for range 6 {
		t.Right(60)
		t.Up()
		t.Walk(raio)
		t.Down()
		t.DrawCircle(raio/3)
		circle(t, raio/3)
		t.Up()
		t.Walk(-raio)
		t.Down()
	}
	
	
}

func main() {
	raio := 333.3
	pen := NewPen(1000, 1000)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(500, 500) // move o pincel para x 250, y 500
	pen.SetHeading(60)
	pen.SetLineWidth(1)        // coloca o pincel apontando para cima
	pen.Down()
	pen.DrawCircle(raio)
	circle(pen, raio)

	pen.SavePNG("circle.png")
	fmt.Println("PNG file created successfully.")
}