package Router

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type earthquake struct {
	ID             int
	Date           string `json:"date"`
	Time           string `json:"time"`
	Depth          string `json:"depth"`
	Violence       string `json:"violence"`
	Location       string `json:"location"`
}

func Initalize(app *fiber.App) {

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/api/earthquakes", func(ctx *fiber.Ctx) error {
		res, err := http.Get("http://www.koeri.boun.edu.tr/scripts/lst8.asp")
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var earthquakes []earthquake
		var id int = 0

		doc.Find("pre").Each(func(i int, s *goquery.Selection) {
			lines := strings.Split(s.Text(), "\n")
			for _, line := range lines[6:] {
				data := strings.Fields(line)
				if len(data) < 11 {
					continue
				}

				e := earthquake{
					ID:       id,
					Date:     data[0] + " " + data[1],
					Time:     data[1],
					Depth:    data[4],
					Violence: data[6],
					Location: data[8] + data[9],
				}
				id++
				earthquakes = append(earthquakes, e)
			}

		})

		return ctx.JSON(earthquakes)
	})

}
