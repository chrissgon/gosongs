package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

type Song struct {
	Track          string `csv:"Track.Name"`
	Artist         string `csv:"Artist.Name"`
	Genre          string `csv:"Genre"`
	BeatsPerMinute string `csv:"Beats.Per.Minute"`
	Length         string `csv:"Length."`
	Popularity     string `csv:"Popularity"`
}

var songs []Song

func init() {
	readDataSet()
}

func main() {
	startRouter()
}

func readDataSet() {
	file, err := os.OpenFile("./dataset.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err := gocsv.UnmarshalFile(file, &songs); err != nil {
		panic(err)
	}
}

func startRouter() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, songs)
	})

	r.GET("/:i", getSong)

	r.Run()
}

func getSong(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("i"))

	i--

	if i >= 0 && i < 50 {
		c.JSON(200, songs[i])
		return
	}

	c.String(404, "undefined song")
}
