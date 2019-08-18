package router

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/unipota/tsundoku/server/model"
)

type TsundokuData struct {
	Tsundoku int `json:"tsundoku"`
	Kidoku   int `json:"kidokoku"`
	Count    int `json:"count"`
}

func GetShareURLHandler(c echo.Context) error {
	deviceID := c.Get("deviceID").(uuid.UUID)

	books, err := model.GetBooksByDeviceID(deviceID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with get books"})
	}

	tsundokuData := calcTsundoku(books)
	share := &model.Share{
		DeviceID: deviceID,
		Tsundoku: tsundokuData.Tsundoku,
		Kidoku:   tsundokuData.Kidoku,
		Count:    tsundokuData.Count,
	}
	newShare, err := model.NewShare(share)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with create share"})
	}

	shareURL := struct {
		URL string `json:url`
	}{
		URL: c.Scheme() + "://" + c.Request().Host + "/share/" + newShare.ID.String(),
	}

	return c.JSON(http.StatusOK, shareURL)
}

func GetSharePageHandler(c echo.Context) error {
	shareID := c.Param("shareID")
	siteURL := c.Scheme() + "://" + c.Request().Host
	html, err := buildHtml(shareID, siteURL)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.HTML(http.StatusOK, html)
}

func GetOGPImageHandler(c echo.Context) error {
	pathParam := c.Param("shareID")
	shareID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	if !model.IsExistShareID(shareID) {
		return c.NoContent(http.StatusBadRequest)
	}

	share, err := model.GetShareByshareID(shareID)
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	filename := "shareimage.png"
	dirPath := "/tsundoku/ogp/screenshots"
	fullpath := dirPath + "/" + filename

	values := url.Values{}
	values.Add("tsundoku", strconv.Itoa(share.Tsundoku))
	values.Add("kidoku", strconv.Itoa(share.Kidoku))
	values.Add("count", strconv.Itoa(share.Count))
	values.Add("filename", filename)

	_, err = os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0777)
		if err != nil {
			fmt.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	_, err = os.Stat(fullpath)
	if !os.IsNotExist(err) {
		err = os.Remove(fullpath)
		if err != nil {
			fmt.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	resp, err := http.Get("http://ogp:4000/ss" + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	// 新規に画像が生成されたのを確認できるまで待つ
	wait := 20
	for i := 0; i < wait; i++ {
		_, err := os.Stat(fullpath)
		if !os.IsNotExist(err) {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}

	return c.File(fullpath)
}

func calcTsundoku(books []*model.Book) *TsundokuData {
	tsundokuData := &TsundokuData{
		Tsundoku: 0,
		Kidoku:   0,
		Count:    0,
	}

	for _, book := range books {
		thisTsundoku := 0
		if book.ReadHistories[0].ReadPage < book.TotalPages {
			thisTsundoku = int(math.Floor((1.0-float64(book.ReadHistories[0].ReadPage)/float64(book.TotalPages))*float64(book.Price)) + 0.5)
		}

		tsundokuData.Tsundoku += thisTsundoku
		tsundokuData.Kidoku += book.Price - thisTsundoku
		tsundokuData.Count += 1
	}

	return tsundokuData
}

func buildHtml(shareID string, siteURL string) (string, error) {
	tmpl := template.Must(template.ParseFiles("/tsundoku/server/static/share.html"))
	buff := new(bytes.Buffer)
	fw := io.Writer(buff)
	dat := struct {
		SiteURL  string
		ImageURL string
	}{
		SiteURL:  siteURL,

		ImageURL: siteURL + "/ogp/" + shareID,

	}
	if err := tmpl.ExecuteTemplate(fw, "base", dat); err != nil {
		return "", err
	}

	html := string(buff.Bytes())
	return html, nil
}
