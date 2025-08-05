package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const mapSize = 10

type GameObject rune

const (
	Empty    GameObject = '.'
	Player   GameObject = 'P'
	Wall     GameObject = '#'
	Treasure GameObject = 'T'
)

type Position struct {
	X, Y int
}

type Game struct {
	Map    [mapSize][mapSize]GameObject
	Player Position
	Score  int
}

func NewGame() *Game {
	game := &Game{
		Player: Position{X: 0, Y: 0},
		Score:  0,
	}

	// Initialize map with empty cells
	for y := range game.Map {
		for x := range game.Map[y] {
			game.Map[y][x] = Empty
		}
	}

	// Place walls
	game.Map[3][3] = Wall
	game.Map[3][4] = Wall
	game.Map[3][5] = Wall

	// Place treasures
	game.Map[2][2] = Treasure
	game.Map[5][5] = Treasure

	// Place player
	game.Map[game.Player.Y][game.Player.X] = Player

	return game
}

func (g *Game) PrintMap() {
	fmt.Println("Map:")
	for y := 0; y < mapSize; y++ {
		for x := 0; x < mapSize; x++ {
			fmt.Printf("%c ", g.Map[y][x])
		}
		fmt.Println()
	}
	fmt.Printf("Score: %d\n", g.Score)
}

func (g *Game) MovePlayer(dx, dy int) {
	newX := g.Player.X + dx
	newY := g.Player.Y + dy

	if newX < 0 || newX >= mapSize || newY < 0 || newY >= mapSize {
		fmt.Println("You can't move outside the map!")
		return
	}

	// Check for wall
	if g.Map[newY][newX] == Wall {
		fmt.Println("There's a wall! You can't move there.")
		return
	}

	// Clear old player position
	g.Map[g.Player.Y][g.Player.X] = Empty

	// Update position
	g.Player.X = newX
	g.Player.Y = newY

	// Place player in new position
	g.Map[g.Player.Y][g.Player.X] = Player
}

func (g *Game) Interact() {
	// Check adjacent cells for treasure
	directions := []Position{
		{X: 0, Y: -1}, // up
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
		{X: 1, Y: 0},  // right
	}

	for _, d := range directions {
		x := g.Player.X + d.X
		y := g.Player.Y + d.Y
		if x >= 0 && x < mapSize && y >= 0 && y < mapSize {
			if g.Map[y][x] == Treasure {
				fmt.Println("You found a treasure!")
				g.Map[y][x] = Empty
				g.Score++
				return
			}
		}
	}
	fmt.Println("There's nothing to interact with.")
}

func main() {
	game := NewGame()
	reader := bufio.NewReader(os.Stdin)

	for {
		game.PrintMap()
		fmt.Print("Enter command (w/a/s/d to move, e to interact, q to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "w":
			game.MovePlayer(0, -1)
		case "s":
			game.MovePlayer(0, 1)
		case "a":
			game.MovePlayer(-1, 0)
		case "d":
			game.MovePlayer(1, 0)
		case "e":
			game.Interact()
		case "q":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}
