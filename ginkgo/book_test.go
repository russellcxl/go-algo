package ginkgo_test

import (
	"github.com/onsi/ginkgo/example/books"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Books", func() {
	var foxInSocks, lesMis *books.Book

	BeforeEach(func() {
		lesMis = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		foxInSocks = &books.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing books", func() {
		Context("with more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(lesMis.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(foxInSocks.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})

	// single spec
	It("can extract the author's last name", func() {
		book := &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		Expect(book.Author).To(Equal("Victor Hugo"))
	})
})
