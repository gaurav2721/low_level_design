
// Book is the item
type Book struct {
    Title string
}

// Iterator interface
type Iterator interface {
    HasNext() bool
    Next() *Book
}

// Aggregate interface
type IterableCollection interface {
    CreateIterator() Iterator
}

type BookCollection struct {
    books []*Book
}

func (b *BookCollection) AddBook(book *Book) {
    b.books = append(b.books, book)
}

func (b *BookCollection) CreateIterator() Iterator {
    return &BookIterator{
        books: b.books,
        index: 0,
    }
}

type BookIterator struct {
    books []*Book
    index int
}

func (it *BookIterator) HasNext() bool {
    return it.index < len(it.books)
}

func (it *BookIterator) Next() *Book {
    if !it.HasNext() {
        return nil
    }
    book := it.books[it.index]
    it.index++
    return book
}

package iterator



// Generic Iterator Interface
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

// Generic Iterable Collection Interface
type IterableCollection[T any] interface {
	CreateIterator() Iterator[T]
}


package iterator

// Generic Slice-based Collection
type SliceCollection[T any] struct {
	items []T
}

func NewSliceCollection[T any](items []T) *SliceCollection[T] {
	return &SliceCollection[T]{items: items}
}

--------------------------


func (c *SliceCollection[T]) Add(item T) {
	c.items = append(c.items, item)
}

func (c *SliceCollection[T]) CreateIterator() Iterator[T] {
	return &SliceIterator[T]{
		items: c.items,
		index: 0,
	}
}

// Generic Iterator over Slice
type SliceIterator[T any] struct {
	items []T
	index int
}

func (it *SliceIterator[T]) HasNext() bool {
	return it.index < len(it.items)
}

func (it *SliceIterator[T]) Next() T {
	item := it.items[it.index]
	it.index++
	return item
}


collection := iterator.NewSliceCollection[*Book](nil)

	collection.Add(&Book{Title: "Go in Action"})
	collection.Add(&Book{Title: "Design Patterns"})
	collection.Add(&Book{Title: "Domain-Driven Design"})

	it := collection.CreateIterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println("Reading:", book.Title)
	}