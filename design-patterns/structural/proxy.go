package structural

import "fmt"

// Image defines the subject interface
type Image interface {
	Display()
}

// RealImage is the real subject
type RealImage struct {
	filename string
}

// NewRealImage creates a new RealImage
func NewRealImage(filename string) *RealImage {
	image := &RealImage{filename: filename}
	image.loadFromDisk()
	return image
}

// loadFromDisk loads the image from disk
func (r *RealImage) loadFromDisk() {
	fmt.Println("Loading image:", r.filename)
}

// Display displays the image
func (r *RealImage) Display() {
	fmt.Println("Displaying image:", r.filename)
}

// ProxyImage is the proxy
type ProxyImage struct {
	realImage *RealImage
	filename  string
}

// NewProxyImage creates a new ProxyImage
func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{
		filename: filename,
	}
}

// Display displays the image
func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display()
}

// ProxyDemo demonstrates the Proxy pattern
func ProxyDemo() {
	image := NewProxyImage("test.jpg")

	// Image will be loaded from disk only when Display() is called
	fmt.Println("First time display:")
	image.Display()

	// Image will not be loaded from disk again
	fmt.Println("\nSecond time display:")
	image.Display()
} 