package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	numbers := [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/api/numbers", func(c *fiber.Ctx) error {
		return c.JSON(numbers)
	})

	app.Get("/api/shuffle", func(c *fiber.Ctx) error {
		rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })
		return c.JSON(numbers)
	})

	app.Post("/api/bingocheck", func(c *fiber.Ctx) error {

		userData := &[]int{}

		if err := c.BodyParser(userData); err != nil {
			fmt.Println(err)
			return err
		}

		var points = bingoCheck(userData)
		return c.JSON(points)
	})

	log.Fatal(app.Listen(":4000"))
}

func bingoCheck(array *[]int) int {
	var points int
	for i := 0; i < 4; i++ {
		if (*array)[i] != 0 && (*array)[i+4] != 0 && (*array)[i+8] != 0 && (*array)[i+12] != 0 {
			points += 1
		}
	}

	for i := 0; i < len(*array); i += 4 {
		if (*array)[i] != 0 && (*array)[i+1] != 0 && (*array)[i+2] != 0 && (*array)[i+3] != 0 {
			points += 1
		}
	}

	if (*array)[0] != 0 && (*array)[5] != 0 && (*array)[10] != 0 && (*array)[15] != 0 {
		points += 1
	}

	if (*array)[3] != 0 && (*array)[6] != 0 && (*array)[9] != 0 && (*array)[12] != 0 {
		points += 1
	}
	return points
}
