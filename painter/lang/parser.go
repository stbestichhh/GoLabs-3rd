package lang

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	lastBgColor painter.Operation
	lastBgRect *painter.BlackRectangle
	figures []*painter.CrossFigure
	moveOps [] painter.Operation
	updateOp painter.Operation
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	if p.lastBgColor == nil {
		p.lastBgColor = painter.OperationFunc(painter.Reset)
	}
	if p.updateOp != nil {
		p.updateOp = nil
	}

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		cmd := scanner.Text()

		error := p.Parse(cmd)
		if error != nil {
			return nil, error
		}

		return p.result(), nil
	}
}

func (p *Parser) result() []painter.Operation {
	var res []painter.Operation
	if p.lastBgColor != nil {
		res = append(res, p.lastBgColor)
	}
	if len(p.moveOps) != 0 {
		res = append(res, p.moveOps...)
	}
	p.moveOps = nil
	if len(p.figures) != 0 {
		for _, figure := range p.figures {
			res = append(res, figure)
		}
	}
	if p.lastBgRect != nil {
		res = append(res, p.lastBgRect)
	}
	if p.updateOp != nil {
		res = append(res, p.updateOp)
	}
	return res
}

func (p *Parser) resetState() {
	p.lastBgColor = nil
	p.lastBgRect = nil
	p.figures = nil
	p.moveOps = nil
	p.updateOp = nil
}

func (p *Parser) parse(commandLine string) error {
	parts := strings.Split(commandLine, " ")
	instruction := parts[0]
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}
	var iArgs []int
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err == nil {
			iArgs = append(iArgs, i)
		}
	}

	switch instruction {
	case "white":
		p.lastBgColor = painter.OperationFunc(painter.WhiteFill)
	case "green":
		p.lastBgColor = painter.OperationFunc(painter.GreenFill)
	case "bgrect":
		p.lastBgRect = &painter.BlackRectangle{X1: iArgs[0], X2: iArgs[1], Y1: iArgs[2], Y2: iArgs[3]}
	case "figure":
		figure := painter.CrossFigure{X: iArgs[0], Y: iArgs[1]}
		p.figures = append(p.figures, &figure)
	case "move":
		moveOp := painter.MoveOperation{X: iArgs[0], Y: iArgs[1], Crosses: p.figures}
		p.moveOps = append(p.moveOps, &moveOp)
	case "reset":
		p.resetState()
		p.lastBgColor = painter.OperationFunc(painter.Reset)
	case "update":
		p.updateOp = painter.UpdateOp
	default:
		return errors.New("could not parse command %v")
	}
	return nil
}
