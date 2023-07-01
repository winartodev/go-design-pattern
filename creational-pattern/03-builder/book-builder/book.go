package bookbuilder

import "fmt"

type Book struct {
	Title           string
	Author          string
	ISBN            string
	PublicationYear int
	Publisher       string
	Description     string
}

type BookBuilder struct {
	Title           string
	Author          string
	ISBN            string
	PublicationYear int
	Publisher       string
	Description     string
}

func (b *BookBuilder) SetTitle(title string) *BookBuilder {
	b.Title = title
	return b
}

func (b *BookBuilder) SetAuthor(author string) *BookBuilder {
	b.Author = author
	return b
}

func (b *BookBuilder) SetISBN(isbn string) *BookBuilder {
	b.ISBN = isbn
	return b
}

func (b *BookBuilder) SetPublicationYear(publicationYear int) *BookBuilder {
	b.PublicationYear = publicationYear
	return b
}

func (b *BookBuilder) SetPublisher(publisher string) *BookBuilder {
	b.Publisher = publisher
	return b
}

func (b *BookBuilder) SetDescription(description string) *BookBuilder {
	b.Description = description
	return b
}

// Implement methods to set other properties of the book...

func (b *BookBuilder) Build() *Book {
	return &Book{
		Title:           b.Title,
		Author:          b.Author,
		ISBN:            b.ISBN,
		PublicationYear: b.PublicationYear,
		Publisher:       b.Publisher,
		Description:     b.Description,
	}
}

func (b *Book) PrintInfo() {
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("ISBN:", b.ISBN)
	fmt.Println("Publication Year:", b.PublicationYear)
	fmt.Println("Publisher:", b.Publisher)
	fmt.Println("Description:", b.Description)
}
