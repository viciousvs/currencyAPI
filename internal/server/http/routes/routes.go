package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viciousvs/currencyAPI/internal/model"
	"github.com/viciousvs/currencyAPI/internal/usecase"
)

type Handler struct {
	CurrencyUseCase usecase.CurrencyUseCase
}

func NewHandler(cu usecase.CurrencyUseCase) Handler {
	return Handler{CurrencyUseCase: cu}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/update", h.update)
	r.GET("/currencies", h.currencies)
	r.GET("/currencies/:name", h.currency)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	return r
}
func (h Handler) update(ctx *gin.Context) {
	res, err := http.Get("https://www.cbr-xml-daily.ru/latest.js")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	currencies := model.Currencies{}
	err = json.Unmarshal(resp, &currencies)
	if err != nil {
		log.Fatal(err)
	}
	err = h.CurrencyUseCase.UpdateCurrencies(ctx.Request.Context(), currencies)
	if err != nil {
		log.Printf(err.Error())
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Updated"})
}
func (h Handler) currency(ctx *gin.Context) {
	name := ctx.Param("name")
	currency, err := h.CurrencyUseCase.GetCurrencyByName(ctx.Request.Context(), name)
	if err != nil {
		log.Printf(err.Error())
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}
	ctx.JSON(http.StatusOK, currency)
}
func (h Handler) currencies(ctx *gin.Context) {
	currencies, err := h.CurrencyUseCase.GetLastRates(ctx.Request.Context())
	if err != nil {
		log.Printf(err.Error())
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}
	ctx.JSON(http.StatusOK, currencies)
}
