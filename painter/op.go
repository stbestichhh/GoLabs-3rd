package painter

import (
	"image/color"
	"image"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

type BlackRectangle struct {
	X1, X2, Y1, Y2 int
}

func (op *BlackRectangle) Do(t screen.Texture) bool {
	if op.X1 > op.X2 {
		op.X1, op.X2 = op.X2, op.X1
	}
	if op.Y1 < op.Y2 {
		op.Y1, op.Y2 = op.Y2, op.Y1
	}

	t.Fill(image.Rect(op.X1, op.Y1, op.X2, op.Y2), color.Black, screen.Src)
	return false
}


type CrossFigure struct {
	X, Y int
}

func (op *CrossFigure) Do(t screen.Texture) bool {
	c := color.RGBA{R: 255, G: 0, B: 0, A: 1}
	t.Fill(image.Rect(op.X-100, op.Y+25, op.X+100, op.Y-25), c, draw.Src)
	t.Fill(image.Rect(op.X-25, op.Y+100, op.X+25, op.Y-100), c, draw.Src)
	return false
}

type MoveOperation struct {
	X, Y int
	Crosses []*CrossFigure
}

func (op *MoveOperation) Do(t screen.Texture) bool {
	for i := range op.Crosses {
		op.Crosses[i].X += op.X
		op.Crosses[i].Y += op.Y
	}
	return false
}

func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src)
}
