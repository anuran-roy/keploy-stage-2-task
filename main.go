package main

import (
	"echo-gorm/database"
	"echo-gorm/database/models"
	"echo-gorm/middlewares"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getNewJoke(ctx echo.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	url := fmt.Sprintf(
		`https://v2.jokeapi.dev/joke/%s?blacklistFlags=nsfw,religious,political,racist,sexist,explicit&type=single`,
		ctx.QueryParam("category"),
	)

	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("\n\nThe HTTP request failed with error %s", err)
		log.Fatal(err)
		return err
	}

	resp_parsed, parsing_err := ioutil.ReadAll(resp.Body)

	if parsing_err != nil {
		fmt.Printf("\n\nThe HTTP response parsing failed with error %s", err)
		log.Fatal(parsing_err)
		return parsing_err
	}
	fmt.Println()
	// fmt.Println(resp_parsed)
	fmt.Println()
	sb := string(resp_parsed)
	// fmt.Printf("resp_parsed: %v\n", resp_parsed)
	sb = strings.Replace(sb, "\"type\":", "\"joke_type\":", -1)
	sb = strings.Replace(sb, "\"error\":", "\"errors\":", -1)

	// bytes_modified_sb := []byte(sb)

	var resp_struct models.JokeResponse
	fmt.Println("Bytes of modified sb = ")
	// fmt.Println(bytes_modified_sb == resp_parsed)
	parse_error := json.Unmarshal(json.RawMessage(sb), &resp_struct)
	fmt.Println(resp_struct.Flags)
	fmt.Println(resp_struct.Id)

	if parse_error != nil {
		panic(parse_error)
	}
	fmt.Println("\n\nJSON version = ")
	fmt.Println(resp_struct)

	fmt.Printf("\n\nThe HTTP response is %s", sb)

	var joke models.Joke
	if db.Find(&joke, "joke = ?", resp_struct.Joke).RowsAffected != 0 {
		// fmt.Println("Joke already exists")
		return ctx.JSON(http.StatusOK, joke)
	} else {
		// fmt.Println("Joke does not exist")
		var new_joke models.Joke
		new_joke.Id = resp_struct.Id
		new_joke.Joke = resp_struct.Joke
		new_joke.Category = resp_struct.Category

		msg := db.Select("Joke", "Category", "Id").Create(&new_joke)
		fmt.Printf("\n\nRows affected = %d", msg.RowsAffected)
		return ctx.JSON(http.StatusOK, new_joke)
	}
}

func getJoke(ctx echo.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	var db_len int64
	var joke models.Joke

	if db.Model(&models.Joke{}).Find(&joke).RowsAffected == 0 {
		joke.Joke = "No jokes found"
		return ctx.JSON(http.StatusNotFound, joke)
	} else {
		var joke models.Joke
		db.Model(&models.Joke{}).Count(&db_len)
		// rand_len := rand.Int63n(db_len + 1)
		// fmt.Printf("\n\nRandom number generated by rand.Int63n = %d\n\n", rand_len) // rand.Int63n(db_len))
		db.Model(&models.Joke{}).Last(&joke, "joke LIKE ?", "%")

		fmt.Printf("\nJoke found, returning sample...\n")
		return ctx.JSON(http.StatusOK, joke)
	}
}

func init() {
	godotenv.Load("./.env")
}

func main() {
	e := echo.New()
	db, _ := database.Connect()

	e.Use(middlewares.ContextDB(db))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/joke", getJoke)
	e.GET("/new_joke", getNewJoke)
	e.Logger.Fatal(e.Start(":1323"))
}
