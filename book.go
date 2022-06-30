package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
	"time"

	"fmt"
	"log"
	"net/http"
)

type Book struct {
	BookID      int       `json:"BookID,omitempty"`
	Title       string    `json:"Title,omitempty"`
	Language    string    `json:"Language,omitempty"`
	ISBN        string    `json:"ISBN,omitempty"`
	Pages       int       `json:"Pages,omitempty"`
	ReadingAge  int       `json:"ReadingAge,omitempty"`
	Genres      []string  `json:"Genres,omitempty"`
	PublishedAt time.Time `json:"PublishedAt,omitempty"`
	Authors     []Author
	Publisher   Publisher
}

func (b *Book) Scan(src interface{}) error {
	fail := func(err error) error {
		return fmt.Errorf("Book.Scan: %w", err)
	}

	data, ok := src.([]byte)
	if !ok {
		return fail(errors.New("type assertions to []byte failed"))
	}

	if err := json.Unmarshal(data, &b); err != nil {
		return fail(fmt.Errorf("json serialisation failed: %w", err))
	}
	return nil
}

type Author struct {
	AuthorID int    `json:"AuthorID,omitempty"`
	Name     string `json:"Name,omitempty"`
}

type Publisher struct {
	PublisherID int    `json:"PublisherID,omitempty"`
	Name        string `json:"Name,omitempty"`
}

type BookFilter struct {
	Title         string
	Language      string
	ISBN          string
	MaxReadingAge int
	Genres        []string
	Authors       []string
	Publisher     string
}

type BookSorter struct {
	Title       SortOrder
	PublishedAt SortOrder
}

type SortOrder int

const (
	OrderUndefined SortOrder = iota
	OrderAscending
	OrderDescending
)

func (so SortOrder) Value() (driver.Value, error) {
	switch so {
	case OrderAscending:
		return "asc", nil
	case OrderDescending:
		return "desc", nil
	default:
		return nil, nil
	}
}

type BookHandler struct {
	BookStore BookStore
}

func (h BookHandler) GetRoutes() []Route {
	return []Route{
		{"/books", http.MethodGet, h.ListBooks},
	}
}

func (h BookHandler) ListBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fail := func(err error, code int) {
		log.Printf("BookHandler.ListBooks: %v\n", err)
		http.Error(w, http.StatusText(code), code)
	}

	// Validate and parse URL query parameters.
	queryParams := QueryParameters{r.URL.Query()}
	maxReadingAge, err := queryParams.GetInt("max_reading_age")
	if err != nil {
		fail(err, http.StatusBadRequest)
		return
	}
	bookFilter := BookFilter{
		Title:         queryParams.GetString("title"),
		Language:      queryParams.GetString("language"),
		ISBN:          queryParams.GetString("isbn"),
		Publisher:     queryParams.GetString("publisher"),
		Genres:        queryParams.GetStringSlice("genre"),
		Authors:       queryParams.GetStringSlice("author"),
		MaxReadingAge: maxReadingAge,
	}
	var bookSorter BookSorter
	if sortField, sortOrder, err := queryParams.GetSortBy("sort_by"); err != nil {
		fail(err, http.StatusBadRequest)
		return
	} else if len(sortField) != 0 {
		switch sortField {
		case "title":
			bookSorter = BookSorter{Title: sortOrder}
		case "published_at":
			bookSorter = BookSorter{PublishedAt: sortOrder}
		default:
			fail(fmt.Errorf("unknown sort field %q", sortField), http.StatusBadRequest)
			return
		}
	}

	// Get books from storage, filtered, and sorted.
	books, err := h.BookStore.ListBooks(bookFilter, bookSorter)
	if err != nil {
		fail(err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Books []Book
	}{books})
}

type BookStore struct {
	DB *sql.DB
}

func (s BookStore) ListBooks(filter BookFilter, sorter BookSorter) ([]Book, error) {
	fail := func(err error) ([]Book, error) {
		return nil, fmt.Errorf("BookStore.ListBooks: %w", err)
	}

	query := `
SELECT JSON_BUILD_OBJECT(
               'BookID', b.book_id,
               'Title', b.title,
               'Language', b.language,
               'ISBN', b.isbn,
               'Pages', b.pages,
               'ReadingAge', b.reading_age,
               'Genres', (SELECT array_agg(g.genre)
                          FROM book_genre bg
                                   JOIN genre g on bg.genre_id = g.genre_id
                          WHERE bg.book_id = b.book_id),
               'Authors', (SELECT json_agg(json_build_object('AuthorID', a.author_id, 'Name', a.name))
                           FROM book_author ba
                                    JOIN author a on ba.author_id = a.author_id
                           WHERE ba.book_id = b.book_id),
               'Publisher', json_build_object('PublisherID', p.publisher_id, 'Name', p.name),
               'PublishedAt', timestamptz(b.published_at)
           )
FROM book b
         JOIN publisher p on b.publisher_id = p.publisher_id
WHERE (coalesce($1, '') = '' OR b.title ILIKE ('%' || $1 || '%'))
  AND (coalesce($2, '') = '' OR b.language = $2)
  AND (coalesce($3, '') = '' OR b.isbn = $3)
  AND (coalesce($4, 0) = 0 OR b.reading_age IS NOT NULL AND b.reading_age <= $4)
  AND (coalesce($5, '{}') = '{}'
	OR EXISTS(
		   SELECT * FROM book_genre bg
		   JOIN genre g ON bg.genre_id = g.genre_id
		   WHERE bg.book_id = b.book_id AND g.genre = ANY ($5::VARCHAR[])
		   ))
  AND (coalesce($6, '{}') = '{}'
	OR EXISTS(
		   SELECT *
		   FROM book_author ba
		   JOIN author a ON ba.author_id = a.author_id
		   WHERE ba.book_id = b.book_id
			 AND a.name ILIKE ANY
				 ((SELECT array_agg(('%' || author_name || '%')) FROM unnest($6::VARCHAR[]) AS author_name)::VARCHAR[])
		   ))
  AND (coalesce($7, '') = '' OR EXISTS(
	SELECT * FROM publisher p WHERE p.publisher_id = b.publisher_id AND p.name ILIKE ('%' || $7 || '%')
	))
ORDER BY CASE WHEN $8 = 'asc' THEN b.title END ASC,
         CASE WHEN $8 = 'desc' THEN b.title END DESC,
         CASE WHEN $9 = 'asc' THEN b.published_at END ASC,
         CASE WHEN $9 = 'desc' THEN b.published_at END DESC;
`
	rows, err := s.DB.Query(query, filter.Title, filter.Language, filter.ISBN, filter.MaxReadingAge, pq.Array(filter.Genres), pq.Array(filter.Authors), filter.Publisher, sorter.Title, sorter.PublishedAt)
	if err != nil {
		return fail(err)
	}

	books := []Book{}
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book); err != nil {
			return fail(err)
		}
		books = append(books, book)
	}

	if err := rows.Close(); err != nil {
		return fail(err)
	}

	return books, nil
}
