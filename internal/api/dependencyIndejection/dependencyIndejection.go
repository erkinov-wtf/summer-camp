package dependencyIndejection

import (
	"camp-summer/internal/api/controller/app"
	"camp-summer/internal/api/interfaces"
)

type Container struct {
	AppController interfaces.AppInterface
}

func NewContainer() *Container {
	return &Container{
		AppController: app.NewAppController(),
	}
}
