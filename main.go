package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/shareed2k/goth_fiber"
	"log"
	"sso-login/service"
	"sso-login/utils"
	"time"
)

func Auth(ctx *fiber.Ctx) error {
	url, err := goth_fiber.GetAuthURL(ctx)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "redirect_request_id",
		Value:    ctx.Query("redirect_to"),
		Path:     "/",
		HTTPOnly: true,
	})

	return ctx.Redirect(url, fiber.StatusTemporaryRedirect)
}

func AuthCallback(ctx *fiber.Ctx) error {

	user, err := goth_fiber.CompleteUserAuth(ctx)
	if err != nil {
		return err
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return err
	}

	redirectTo := ctx.Cookies("redirect_request_id")

	ctx.Cookie(&fiber.Cookie{
		Name:     "redirect_request_id",
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
	})

	redirectPath := utils.NewURLBuilder("http", "localhost:3000").
		AddPath(redirectTo).
		AddQuery("token", token).
		AddQuery("redirect_to", redirectTo).
		Build()

	return ctx.Redirect(redirectPath)
}

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()

	service.NewAuth()

	app.Get("/auth/:provider", Auth)
	app.Get("/auth/:provider/callback", AuthCallback)

	log.Fatal(app.Listen(":7007"))
}
