package structural

import (
	"fmt"
)

// FileSystemComponent defines the component interface
type FileSystemComponent interface {
	Display(indent string)
	GetSize() int64
}

// File is a leaf class
type File struct {
	name string
	size int64
}

// NewFile creates a new File
func NewFile(name string, size int64) *File {
	return &File{
		name: name,
		size: size,
	}
}

// Display implements the FileSystemComponent interface
func (f *File) Display(indent string) {
	fmt.Println(indent + "File: " + f.name + " (" + fmt.Sprintf("%d", f.size) + " bytes)")
}

// GetSize implements the FileSystemComponent interface
func (f *File) GetSize() int64 {
	return f.size
}

// Directory is a composite class
type Directory struct {
	name     string
	children []FileSystemComponent
}

// NewDirectory creates a new Directory
func NewDirectory(name string) *Directory {
	return &Directory{
		name:     name,
		children: make([]FileSystemComponent, 0),
	}
}

// Add adds a component to the directory
func (d *Directory) Add(component FileSystemComponent) {
	d.children = append(d.children, component)
}

// Remove removes a component from the directory
func (d *Directory) Remove(component FileSystemComponent) {
	for i, child := range d.children {
		if child == component {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
}

// Display implements the FileSystemComponent interface
func (d *Directory) Display(indent string) {
	fmt.Println(indent + "Directory: " + d.name)
	for _, child := range d.children {
		child.Display(indent + "  ")
	}
}

// GetSize implements the FileSystemComponent interface
func (d *Directory) GetSize() int64 {
	var totalSize int64
	for _, child := range d.children {
		totalSize += child.GetSize()
	}
	return totalSize
}

// CompositeDemo demonstrates the Composite pattern
func CompositeDemo() {
	// Create files
	file1 := NewFile("document.txt", 100)
	file2 := NewFile("image.jpg", 200)
	file3 := NewFile("data.csv", 150)

	// Create directories
	dir1 := NewDirectory("Documents")
	dir2 := NewDirectory("Pictures")
	root := NewDirectory("Root")

	// Build the tree structure
	dir1.Add(file1)
	dir2.Add(file2)
	dir2.Add(file3)
	root.Add(dir1)
	root.Add(dir2)

	// Display the structure
	fmt.Println("File System Structure:")
	root.Display("")

	// Calculate total size
	fmt.Println("\nTotal size:", root.GetSize(), "bytes")
} 