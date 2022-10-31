package main

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"math"
	"net/http"
)

type Item struct {
	Success bool        `json:"success"`
	Data1   interface{} `json:"data1"`
	Data2   interface{} `json:"data2"`
}

type Student struct {
	Email string `json:"email"`
}

type User struct {
	Name string `json:"name"`
}

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 21)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	//file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 666)
	//if err != nil {
	//	panic(err)
	//}
	//log.SetOutput(file)

	log.SetFormatter(&log.JSONFormatter{})

	log.Printf("test")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("users/:name/:email", getUser)
	//e.Start(":1323")

	e.Logger.Fatal(e.Start(":1323"))

}

func getUser(c echo.Context) error {
	//test := c.QueryParam("test")
	//param := c.Param("name")
	//return c.String(http.StatusOK, test)
	//return c.JSON(http.StatusOK, test)
	name := c.Param("name")
	email := c.Param("email")
	return c.JSON(http.StatusOK, Item{
		Success: true,
		Data1:   User{Name: name},
		Data2:   Student{Email: email},
	})
}
