package main

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
	filename string
}

func (r *RealImage) loadFromDisk() {
	fmt.Println("Loading", r.filename)
}

func (r *RealImage) Display() {
	fmt.Println("Displaying", r.filename)
}

type ProxyImage struct {
	filename string
	real     *RealImage
}

func (p *ProxyImage) Display() {
	if p.real == nil {
		p.real = &RealImage{filename: p.filename}
		p.real.loadFromDisk()
	}
	p.real.Display()
}

func main() {
	img := &ProxyImage{filename: "photo.png"}

	// Image is loaded only when Display is called
	img.Display()
	fmt.Println("Calling again...")
	img.Display()
}
