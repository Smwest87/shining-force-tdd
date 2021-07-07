package models

import (
	"errors"
	"fmt"
)

type Game struct {
}

type Player struct {
	Name string
	X    int64
	Y    int64
}

type Level struct {
	ID          int64
	MaxX        int64
	MaxY        int64
	OutOfBounds []Coordinates
}

type Coordinates struct {
	X int64
	Y int64
}

func (p *Player) Move(x int64, y int64, level Level) error {
	if p.X+x < 0 {
		message := fmt.Sprintf("player cannot exit the map. X: %d", p.X+x)
		err := errors.New(message)
		return err
	}

	if p.X+x > level.MaxX {
		message := fmt.Sprintf("player cannot exit the map. X: %d\n Level MaxX: %d", p.X+x, level.MaxX)
		err := errors.New(message)
		return err
	}

	if p.Y+y < 0 {
		message := fmt.Sprintf("player cannot exit the map. Y: %d", p.Y+y)
		err := errors.New(message)
		return err
	}

	if p.Y+y > level.MaxY {
		message := fmt.Sprintf("player cannot exit the map. Y: %d\n Level MaxY: %d", p.Y+y, level.MaxY)
		err := errors.New(message)
		return err
	}

	checkCoordinate := Coordinates{
		X: p.X + x,
		Y: p.Y + y,
	}

	for _, point := range level.OutOfBounds {
		if checkCoordinate == point {
			message := fmt.Sprintf("character cannot enter out of bounds.\nattempted coordinate: X: %d\n Y: %d\nout of bounds coordinate: X: %d\nY: %d", checkCoordinate.X, checkCoordinate.Y, point.X, point.Y)
			err := errors.New(message)
			return err
		}
	}

	p.X = p.X + x
	p.Y = p.Y + y
	return nil
}
