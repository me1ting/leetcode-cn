package hj17movepoint

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readInputs()
	slices := strings.Split(input, ";")
	pos := Pos{0, 0}
	for _, str := range slices {
		action := ParseAction(str)
		if action != nil {
			pos.DoAction(action.t, action.d)
		}
	}
	fmt.Printf("%d,%d\n", pos.X, pos.Y)
}

func ParseAction(str string) *Action {
	r := regexp.MustCompile(`^(A|W|D|S)(\d+)$`)
	matchs := r.FindStringSubmatch(str)
	if matchs == nil {
		return nil
	}

	distance, _ := strconv.Atoi(matchs[2])

	return &Action{
		t: matchs[1],
		d: distance,
	}
}

type Action struct {
	t string //type
	d int
}

type Pos struct {
	X int
	Y int
}

func (p *Pos) DoAction(t string, d int) {
	if t == "A" {
		p.Left(d)
	} else if t == "D" {
		p.Right(d)
	} else if t == "W" {
		p.Top(d)
	} else if t == "S" {
		p.Bottum(d)
	}
}

func (p *Pos) Left(d int) {
	p.X -= d
}

func (p *Pos) Right(d int) {
	p.X += d
}

func (p *Pos) Bottum(d int) {
	p.Y -= d
}

func (p *Pos) Top(d int) {
	p.Y += d
}

func readInputs() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Main() (int, int) {
	input := "A10;S20;W10;D30;X;A1A;B10A11;;A10;"
	slices := strings.Split(input, ";")
	pos := Pos{0, 0}
	for _, str := range slices {
		action := ParseAction(str)
		if action != nil {
			pos.DoAction(action.t, action.d)
		}
	}
	return pos.X, pos.Y
}
