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
	"google.golang.org/api/books/v1"
)

var (
	RAKUTEN_APPLICATION_ID     = os.Getenv("RAKUTEN_APPLICATION_ID")
	RAKUTEN_APPLICATION_SECRET = os.Getenv("RAKUTEN_APPLICATION_SECRET")
)

type ISBN struct {
	ISBN string `json:"isbn"`
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

	bookRecords := searchBooks(true, isbn)

	return c.JSON(http.StatusOK, bookRecords)
}

func SearchWithWord(c echo.Context) error {
	search := c.QueryParam("search")

	bookRecords := searchBooks(false, search)

	return c.JSON(http.StatusOK, bookRecords)
}

func volume2BookRecord(volume *books.Volume) *BookRecord {
	bookRecord := &BookRecord{
		ID:            "",
		ISBN:          "",
		Title:         volume.VolumeInfo.Title,
		TotalPages:    int(volume.VolumeInfo.PageCount),
		Price:         0,
		Caption:       "",
		Publisher:     volume.VolumeInfo.Publisher,
		CoverImageURL: "",
		ReadPages:     0,
		Memo:          "",
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
				bookRecord.CoverImageURL = imageLink
				break
			}
		}

		authorList := make([]string, 0)
		if volume.VolumeInfo.Authors != nil {
			for _, author := range volume.VolumeInfo.Authors {
				authorList = append(authorList, author)
			}
		}
		bookRecord.Author = authorList
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
		bookRecord.ISBN = isbn13
	} else if isbn10 != "" {
		bookRecord.ISBN = isbn10
	}

	return bookRecord
}

func volumes2BookRecords(volumes *books.Volumes) []*BookRecord {
	bookRecords := make([]*BookRecord, 0)

	for _, volume := range volumes.Items {
		bookRecord := volume2BookRecord(volume)
		bookRecords = append(bookRecords, bookRecord)
	}

	return bookRecords
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

func rakutenBook2BookRecord(book *RakutenBook) *BookRecord {
	bookRecord := &BookRecord{
		ID:            "",
		ISBN:          book.ISBN,
		Title:         book.Title,
		TotalPages:    0,
		Price:         book.ItemPrice,
		Caption:       "",
		Publisher:     book.PublisherName,
		CoverImageURL: "",
		ReadPages:     0,
		Memo:          "",
	}

	authorList := make([]string, 0)
	authors := strings.Split(book.Author, "/")
	for _, author := range authors {
		if author != "" {
			authorList = append(authorList, author)
		}
	}
	bookRecord.Author = authorList

	imageLinks := [...]string{
		book.LargeImageURL,
		book.MediumImageURL,
		book.SmallImageURL,
	}

	for _, imageLink := range imageLinks {
		if imageLink != "" {
			bookRecord.CoverImageURL = imageLink
			break
		}
	}

	return bookRecord
}

func rakutenBooks2BookRecords(books *RakutenBooks) []*BookRecord {
	bookRecords := make([]*BookRecord, 0)
	bookList := books.Items

	for _, bookItem := range bookList {
		bookRecord := rakutenBook2BookRecord(bookItem.Item)
		bookRecords = append(bookRecords, bookRecord)
	}

	return bookRecords
}

func withGoogle(isIsbn bool, searchWord string, bookRecordsChan chan []*BookRecord) {
	values := url.Values{}

	if isIsbn {
		values.Add("q", "isbn:"+searchWord)
	} else {
		values.Add("q", searchWord)
	}

	volumes, err := searchWithGoogle(values)
	if err != nil {
		bookRecordsChan <- nil
	}

	bookRecordsChan <- volumes2BookRecords(volumes)
}

func withRakuten(isIsbn bool, searchWord string, bookRecordsChan chan []*BookRecord) {
	values := url.Values{}

	if isIsbn {
		values.Add("isbn", searchWord)
	} else {
		values.Add("title", searchWord)
	}

	rakutenBooks, err := searchWithRakuten(values)
	if err != nil {
		bookRecordsChan <- nil
	}

	bookRecordsChan <- rakutenBooks2BookRecords(rakutenBooks)
}

func searchBooks(isIsbn bool, searchWord string) []*BookRecord {
	googleChan := make(chan []*BookRecord)
	rakutenChan := make(chan []*BookRecord)

	go withGoogle(isIsbn, searchWord, googleChan)
	go withRakuten(isIsbn, searchWord, rakutenChan)
	googleBooks := <-googleChan
	rakutenBooks := <-rakutenChan

	return mergeBookRecords(rakutenBooks, googleBooks)
}

func mergeBookRecord(mainBook *BookRecord, subBook *BookRecord) *BookRecord {
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

func mergeBookRecords(mainBooks []*BookRecord, subBooks []*BookRecord) []*BookRecord {
	bookRecords := make([]*BookRecord, 0)

	for _, mainBook := range mainBooks {
		bookRecord := mainBook
		if mainBook.ISBN != "" {
			for _, subBook := range subBooks {
				if subBook.ISBN == mainBook.ISBN {
					bookRecord = mergeBookRecord(mainBook, subBook)
					break
				}
			}
		} else {
			for _, subBook := range subBooks {
				if subBook.Title == mainBook.Title {
					bookRecord = mergeBookRecord(mainBook, subBook)
					break
				}
			}
		}
		bookRecords = append(bookRecords, bookRecord)
	}

	return bookRecords
}
