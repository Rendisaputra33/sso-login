package service

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"os"
)

func NewAuth() goth.Providers {
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_KEY"),
			os.Getenv("GOOGLE_CLIENT_SECRET_KEY"),
			os.Getenv("GOOGLE_CALLBACK_URL"),
		),
	)

	return goth.GetProviders()
}
