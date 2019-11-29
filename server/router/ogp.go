package router

import (
	"bytes"
	"html/template"
	"io"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strconv"

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
	html, err := buildHTML(shareID, siteURL)
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

	filename := shareID.String() + ".png"
	dirPath := "ogp_canvas/screenshots"
	fullpath := dirPath + "/" + filename

	// ファイルが既に存在する場合
	if _, err := os.Stat(fullpath); err == nil {
		return c.File(fullpath)
	}

	// if isNodeAvailable, err := isCommandAvailable("node"); !isNodeAvailable || err != nil {
	// 	c.Logger().Error(err)
	// 	return c.NoContent(http.StatusBadRequest)
	// }

	cmd := exec.Command("node", "ogp_canvas/main.js", shareID.String(), strconv.Itoa(share.Count), strconv.Itoa(share.Tsundoku), strconv.Itoa(share.Kidoku))
	out, err := cmd.CombinedOutput()
	c.Logger().Printf("%s", out)
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.File(fullpath)
}

func isCommandAvailable(name string) (bool, error) {
	cmd := exec.Command(name, "-v")
	if err := cmd.Run(); err != nil {
		return false, err
	}
	return true, nil
}

func calcTsundoku(books []*model.Book) *TsundokuData {
	tsundokuData := &TsundokuData{
		Tsundoku: 0,
		Kidoku:   0,
		Count:    0,
	}

	for _, book := range books {
		if book.TotalPages == 0 || book.ReadHistories[0].ReadPage < book.TotalPages {
			if book.TotalPages == 0 {
				tsundokuData.Tsundoku += book.Price
			} else {
				tsundokuData.Tsundoku += int(math.Floor((1.0-float64(book.ReadHistories[0].ReadPage)/float64(book.TotalPages))*float64(book.Price)) + 0.5)
			}
		} else {
			tsundokuData.Kidoku += book.Price
		}
		tsundokuData.Count++
	}

	return tsundokuData
}

func buildHTML(shareID string, siteURL string) (string, error) {
	tmpl := template.Must(template.ParseFiles("/tsundoku/server/static/share.html"))
	buff := new(bytes.Buffer)
	fw := io.Writer(buff)
	dat := struct {
		SiteURL  string
		ImageURL string
	}{
		SiteURL: siteURL + "/share/" + shareID,

		ImageURL: siteURL + "/ogp/" + shareID,
	}
	if err := tmpl.ExecuteTemplate(fw, "base", dat); err != nil {
		return "", err
	}

	html := string(buff.Bytes())
	return html, nil
}
