package rest

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

// ReqExchange binding exchange API query parameter
type ReqExchange struct {
	Source string `form:"source"`
	Target string `form:"target"`
	Amount string `form:"amount"`
}

// AmountValue will parse amount string to float64
func (r *ReqExchange) AmountValue() (float64, error) {
	v := strings.Replace(r.Amount[1:], ",", "", -1)
	return strconv.ParseFloat(v, 64)
}

// QueryExchange implement GET exchange API
func QueryExchange(r *gin.Engine) {
	r.GET("/api/exchange", func(c *gin.Context) {
		var req ReqExchange
		err := c.ShouldBindQuery(&req)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		amount, err := req.AmountValue()
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		result := exchange(req.Source, req.Target, amount)

		p := message.NewPrinter(message.MatchLanguage("en"))

		c.JSON(http.StatusOK, gin.H{
			"msg":    "success",
			"amount": p.Sprintf("$%.2f", result),
		})
	})
}

var rateMap map[string]map[string]float64

// LoadRateFile will load rate.json
// MUST initial before run server
func LoadRateFile(filename string) error {
	var rateTemp map[string]map[string]map[string]float64
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	jsonData, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &rateTemp)
	if err != nil {
		return err
	}

	rateMap = make(map[string]map[string]float64)

	for i := range rateTemp["currencies"] {
		rateMap[i] = make(map[string]float64)
		rateMap[i] = rateTemp["currencies"][i]
	}

	return nil
}

func exchange(source, target string, amount float64) float64 {
	return math.Round(rateMap[source][target]*amount*100) / 100
}
