package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/seihmd/openbd"
	"google.golang.org/api/books/v1"
)

var (
	RAKUTEN_APPLICATION_ID     = os.Getenv("RAKUTEN_APPLICATION_ID")
	RAKUTEN_APPLICATION_SECRET = os.Getenv("RAKUTEN_APPLICATION_SECRET")
)

type ISBN struct {
	ISBN string `json:"isbn"`
}

type SearchedBook struct {
	ID            string   `json:"id"`
	ISBN          string   `json:"isbn"`
	Title         string   `json:"title"`
	Author        []string `json:"author"`
	Price         int      `json:"price"`
	Caption       string   `json:"caption"`
	Publisher     string   `json:"publisher"`
	CoverImageURL string   `json:"coverImageUrl"`
	Memo          string   `json:"memo"`
}

type RakutenBook struct {
	Title          string `json:"title"`
	TitleKana      string `json:"titleKana"`
	SubTitle       string `json:"subTitle"`
	SubTitleKana   string `json:"subTitleKana"`
	SeriesName     string `json:"seriesName"`
	SeriesNameKana string `json:"seriesNameKana"`
	Contents       string `json:"contents"`
	Author         string `json:"author"`
	AuthorKana     string `json:"authorKkana"`
	PublisherName  string `json:"publisherName"`
	Size           string `json:"size"`
	ISBN           string `json:"isbn"`
	ItemPrice      int    `json:"itemPrice"`
	ListPrice      int    `json:"listPrice"`
	DiscountRate   int    `json:"discountRate"`
	DiscountPrice  int    `json:"discountPrice"`
	itemURL        string `json:"itemUrl"`
	affiliateURL   string `json:"affiliateUrl"`
	SmallImageURL  string `json:"smallImageUrl"`
	MediumImageURL string `json:"mediumImageUrl"`
	LargeImageURL  string `json:"largeImageUrl"`
	ChirayomiURL   string `json:"chirayomiUrl"`
	Availability   string `json:"availability"`
	PostageFlag    int    `json:"postageFlag"`
	LimitedFlag    int    `json:"limitedFlag"`
	ReviewCount    int    `json:"reviewCount"`
	ReviewAverage  string `json:"reviewAverage"`
	BooksGenreID   string `json:"booksGenreId"`
}

type RakutenItem struct {
	Item *RakutenBook `json:"Item"`
}

type RakutenBooks struct {
	Count     int            `json:"count"`
	Page      int            `json:"page"`
	First     int            `json:"first"`
	Last      int            `json:"last"`
	Hits      int            `json:"hits"`
	Carrier   int            `json:"carrier"`
	PageCount int            `json:"pageCount"`
	Items     []*RakutenItem `json:"Items"`
}

func SearchWithISBN(c echo.Context) error {
	isbn := c.QueryParam("isbn")

	searchedBooks := searchBooks(true, isbn)

	return c.JSON(http.StatusOK, searchedBooks)
}

func SearchWithWord(c echo.Context) error {
	search := c.QueryParam("search")

	searchedBooks := searchBooks(false, search)

	return c.JSON(http.StatusOK, searchedBooks)
}

func volume2SearchedBook(volume *books.Volume) *SearchedBook {
	searchedBook := &SearchedBook{
		ID:            "",
		ISBN:          "",
		Title:         volume.VolumeInfo.Title,
		Price:         0,
		Caption:       "",
		Publisher:     volume.VolumeInfo.Publisher,
		CoverImageURL: "",
	}

	if volume.VolumeInfo.ImageLinks != nil {
		imageLinks := [...]string{
			volume.VolumeInfo.ImageLinks.ExtraLarge,
			volume.VolumeInfo.ImageLinks.Large,
			volume.VolumeInfo.ImageLinks.Medium,
			volume.VolumeInfo.ImageLinks.Small,
			volume.VolumeInfo.ImageLinks.Thumbnail,
			volume.VolumeInfo.ImageLinks.SmallThumbnail,
		}
		for _, imageLink := range imageLinks {
			if imageLink != "" {
				searchedBook.CoverImageURL = imageLink
				break
			}
		}

		authorList := make([]string, 0)
		if volume.VolumeInfo.Authors != nil {
			for _, author := range volume.VolumeInfo.Authors {
				authorList = append(authorList, author)
			}
		}
		searchedBook.Author = authorList
	}

	industryIdentifiers := volume.VolumeInfo.IndustryIdentifiers
	isbn10 := ""
	isbn13 := ""
	for _, idetifier := range industryIdentifiers {
		if idetifier.Type == "ISBN_10" {
			isbn10 = idetifier.Identifier
		} else if idetifier.Type == "ISBN_13" {
			isbn13 = idetifier.Identifier
		}
	}
	if isbn13 != "" {
		searchedBook.ISBN = isbn13
	} else if isbn10 != "" {
		searchedBook.ISBN = isbn10
	}

	return searchedBook
}

func volumes2searchedBooks(volumes *books.Volumes) []*SearchedBook {
	searchedBooks := make([]*SearchedBook, 0)

	for _, volume := range volumes.Items {
		searchedBook := volume2SearchedBook(volume)
		searchedBooks = append(searchedBooks, searchedBook)
	}

	return searchedBooks
}

func searchWithGoogle(values url.Values) (*books.Volumes, error) {
	resp, err := http.Get("https://www.googleapis.com/books/v1/volumes" + "?" + values.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	volumes := &books.Volumes{}
	if err := json.Unmarshal(body, volumes); err != nil {
		return nil, err
	}

	return volumes, nil
}

func searchWithRakuten(values url.Values) (*RakutenBooks, error) {
	values.Add("applicationId", RAKUTEN_APPLICATION_ID)
	values.Add("application_secret", RAKUTEN_APPLICATION_SECRET)

	resp, err := http.Get("https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404" + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	books := &RakutenBooks{}
	if err := json.Unmarshal(body, books); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return books, nil
}

func searchWithOpenBD(isbn string) (*openbd.Book, error) {
	o := openbd.New()

	book, err := o.Get(isbn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return book, nil
}

func rakutenBook2SearchedBook(book *RakutenBook) *SearchedBook {
	searchedBook := &SearchedBook{
		ID:            "",
		ISBN:          book.ISBN,
		Title:         book.Title,
		Price:         book.ItemPrice,
		Caption:       "",
		Publisher:     book.PublisherName,
		CoverImageURL: "",
	}

	authorList := make([]string, 0)
	authors := strings.Split(book.Author, "/")
	for _, author := range authors {
		if author != "" {
			authorList = append(authorList, author)
		}
	}
	searchedBook.Author = authorList

	imageLinks := [...]string{
		book.LargeImageURL,
		book.MediumImageURL,
		book.SmallImageURL,
	}

	for _, imageLink := range imageLinks {
		if imageLink != "" {
			searchedBook.CoverImageURL = imageLink
			break
		}
	}

	return searchedBook
}

func rakutenBooks2SearchedBooks(books *RakutenBooks) []*SearchedBook {
	searchedBooks := make([]*SearchedBook, 0)
	bookList := books.Items

	for _, bookItem := range bookList {
		searchedBook := rakutenBook2SearchedBook(bookItem.Item)
		searchedBooks = append(searchedBooks, searchedBook)
	}

	return searchedBooks
}

func openBDBook2SearchedBook(book *openbd.Book) *SearchedBook {
	searchedBook := &SearchedBook{
		ID:            "",
		ISBN:          book.GetISBN(),
		Title:         book.GetTitle(),
		Price:         0,
		Caption:       "",
		Publisher:     book.GetPublisher(),
		CoverImageURL: book.GetCover(),
	}

	authorList := make([]string, 0)
	authors := strings.Split(book.GetAuthor(), "/")
	for _, author := range authors {
		if author != "" {
			authorList = append(authorList, author)
		}
	}
	searchedBook.Author = authorList

	if len(book.ProductSupply.SupplyDetail.Price) > 0 {
		price, _ := book.ProductSupply.SupplyDetail.Price[0].PriceAmount.Int64()
		searchedBook.Price = int(price)
	}

	return searchedBook
}

func withGoogle(isIsbn bool, searchWord string, searchedBooksChan chan []*SearchedBook) {
	values := url.Values{}

	if isIsbn {
		values.Add("q", "isbn:"+searchWord)
	} else {
		values.Add("q", searchWord)
	}

	volumes, err := searchWithGoogle(values)
	if err != nil {
		searchedBooksChan <- nil
	}

	searchedBooksChan <- volumes2searchedBooks(volumes)
}

func withRakuten(isIsbn bool, searchWord string, searchedBooksChan chan []*SearchedBook) {
	values := url.Values{}

	if isIsbn {
		values.Add("isbn", searchWord)
	} else {
		values.Add("title", searchWord)
	}

	rakutenBooks, err := searchWithRakuten(values)
	if err != nil {
		searchedBooksChan <- nil
	}

	searchedBooksChan <- rakutenBooks2SearchedBooks(rakutenBooks)
}

func withOpenBD(isIsbn bool, searchWord string, searchedBookChan chan *SearchedBook) {
	if !isIsbn {
		searchedBookChan <- nil
	}

	openBDBook, err := searchWithOpenBD(searchWord)
	if err != nil {
		searchedBookChan <- nil
	}

	searchedBookChan <- openBDBook2SearchedBook(openBDBook)
}

func searchBooks(isIsbn bool, searchWord string) []*SearchedBook {
	googleChan := make(chan []*SearchedBook)
	rakutenChan := make(chan []*SearchedBook)
	openBDChan := make(chan *SearchedBook)

	go withGoogle(isIsbn, searchWord, googleChan)
	go withRakuten(isIsbn, searchWord, rakutenChan)
	go withOpenBD(isIsbn, searchWord, openBDChan)
	googleBooks := <-googleChan
	rakutenBooks := <-rakutenChan
	openBDBook := <-openBDChan

	return mergeSearchedBooks(rakutenBooks, googleBooks, openBDBook)
}

func mergeBookRecord(mainBook *SearchedBook, subBook *SearchedBook) *SearchedBook {
	if mainBook.ISBN == "" {
		mainBook.ISBN = subBook.ISBN
	}
	if mainBook.Price == 0 {
		mainBook.Price = subBook.Price
	}
	if len(mainBook.Author) == 0 {
		mainBook.Author = subBook.Author
	}
	if mainBook.Publisher == "" {
		mainBook.Publisher = subBook.Publisher
	}
	if mainBook.CoverImageURL == "" {
		mainBook.CoverImageURL = subBook.CoverImageURL
	}

	return mainBook
}

func mergeSearchedBooks(mainBooks []*SearchedBook, subBooks []*SearchedBook, subBook2 *SearchedBook) []*SearchedBook {
	searchedBooks := make([]*SearchedBook, 0)

	for _, mainBook := range mainBooks {
		searchedBook := mainBook
		if mainBook.ISBN != "" {
			for _, subBook := range subBooks {
				if subBook.ISBN == mainBook.ISBN {
					searchedBook = mergeBookRecord(mainBook, subBook)
					break
				}
			}
		} else {
			for _, subBook := range subBooks {
				if subBook.Title == mainBook.Title {
					searchedBook = mergeBookRecord(mainBook, subBook)
					break
				}
			}
		}
		searchedBooks = append(searchedBooks, searchedBook)
	}

	if len(mainBooks) == 0 && subBook2 != nil {
		searchedBooks = append(searchedBooks, subBook2)
	}

	return searchedBooks
}
