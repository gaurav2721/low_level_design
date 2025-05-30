package main

import "fmt"

type Character struct {
	symbol    string // Intrinsic state
	font      string
	fontSize  int
	color     string
}

type CharacterFactory struct {
	chars map[string]*Character
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{chars: make(map[string]*Character)}
}

func (f *CharacterFactory) GetCharacter(symbol string) *Character {
	if char, ok := f.chars[symbol]; ok {
		return char
	}
	char := &Character{symbol: symbol}
	f.chars[symbol] = char
	return char
}

// Display applies extrinsic data
func (c *Character) Display(font string, size int, color string) {
	fmt.Printf("Symbol: %s | Font: %s | Size: %d | Color: %s\n",
		c.symbol, font, size, color)
}

func main() {
	factory := NewCharacterFactory()

	a := factory.GetCharacter("A")
	b := factory.GetCharacter("B")
	a2 := factory.GetCharacter("A") // Shared instance

	a.Display("Arial", 12, "Black")
	b.Display("Times", 14, "Blue")
	a2.Display("Verdana", 10, "Red")
}
