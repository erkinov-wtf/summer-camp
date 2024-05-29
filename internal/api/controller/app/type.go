package app

import (
	"camp-summer/internal/api/interfaces"
)

type appController struct{}

func NewAppController() interfaces.AppInterface {
	return &appController{}
}
