package main

import (
	"context"
	"log"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(titleComponent("Second")),
		fx.Provide(titleComponent("First")),
		fx.Provide(
			fx.Annotate(
				NewPublisher,
				fx.As(new(IPublisher)),
				fx.ParamTags(`group:"titles"`),
			),
		),
		fx.Provide(NewMainService),

		fx.Invoke(func(service *MainService) {
			service.Run()
		}))

	Run(app)
	log.Println("Goodbye")
}

// A way of running the app like a server
func Run(app *fx.App) {
	app.Run()
}

// A way of executing a job and stop is by app.Start with a Context
func Start(app *fx.App) {
	app.Start(context.Background())
}

type MainService struct {
	publisher IPublisher
}

func NewMainService(publisher IPublisher) *MainService {
	return &MainService{publisher: publisher}
}

func (s *MainService) Run() {
	log.Println("Hello world")
	s.publisher.Publish()
}

type IPublisher interface {
	Publish()
}

// dependecy
type Publisher struct {
	titles []*Title
}

func NewPublisher(titles ...*Title) *Publisher {
	return &Publisher{titles: titles}
}

func (p *Publisher) Publish() {
	log.Println("Publish")
	for _, title := range p.titles {
		log.Println("Publish", *title)
	}
}

type Title string

func titleComponent(title string) any {
	return fx.Annotate(
		func() *Title {
			t := Title(title)
			return &t
		},
		fx.ResultTags(`group:"titles"`),
	)
}
