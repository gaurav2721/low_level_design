package main

import "fmt"

// Memento stores the state
type Memento struct {
	state string
}

// Originator
type Editor struct {
	state string
}

func (e *Editor) Write(text string) {
	e.state = text
}

func (e *Editor) Save() Memento {
	return Memento{state: e.state}
}

func (e *Editor) Restore(m Memento) {
	e.state = m.state
}

func (e *Editor) Show() {
	fmt.Println("Current State:", e.state)
}

// Caretaker
type History struct {
	mementos []Memento
}

func (h *History) Push(m Memento) {
	h.mementos = append(h.mementos, m)
}

func (h *History) Pop() Memento {
	if len(h.mementos) == 0 {
		return Memento{}
	}
	m := h.mementos[len(h.mementos)-1]
	h.mementos = h.mementos[:len(h.mementos)-1]
	return m
}

func main() {
	editor := &Editor{}
	history := &History{}

	editor.Write("Version 1")
	history.Push(editor.Save())

	editor.Write("Version 2")
	history.Push(editor.Save())

	editor.Write("Version 3")
	editor.Show()

	// Undo
	editor.Restore(history.Pop())
	editor.Show()

	// Undo again
	editor.Restore(history.Pop())
	editor.Show()
}
