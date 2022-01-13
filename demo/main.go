package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/sg0hsmt/budoux-go"
	"github.com/sg0hsmt/budoux-go/models"
)

type demo struct {
	app.Compo
	input string
}

func (h *demo) Render() app.UI {
	container := map[string]string{
		"font-family": "sans-serif",
		"margin":      "0 auto",
		"max-width":   "960px",
		"padding":     "1rem",
	}

	data := budoux.Parse(models.DefaultJapaneseModel(), h.input)

	return app.Div().Body(
		app.Header().Styles(container).Body(
			app.H1().Text("BudouX-Go Demo"),
			app.P().Body(
				app.Text("BudouX-Go is a golang port of "),
				app.A().Text("BudouX").Href("https://github.com/google/budoux"),
				app.Text(" (machine learning powered line break organizer tool)."),
			),
			app.P().Body(
				app.Text("["),
				app.A().Href("https://github.com/sg0hsmt/budoux-go").Text("GitHub"),
				app.Text("]"),
				app.Text("["),
				app.A().Href("https://pkg.go.dev/github.com/sg0hsmt/budoux-go").Text("GoDoc"),
				app.Text("]"),
			),
		),
		app.Main().Styles(container).Body(
			app.Textarea().Text(h.input).
				Placeholder("日本語の文章を入力してください。").
				Style("width", "100%").
				Style("height", "6em").
				OnInput(h.ValueTo(&h.input)),
			app.If(h.input != "" && len(data) > 0,
				app.Ol().Body(
					app.Range(data).Slice(func(i int) app.UI {
						return app.Li().Text(data[i])
					}),
				),
			),
		),
	)
}

func (h *demo) OnMount(ctx app.Context) {
	h.input = "日本語の文章をいい感じに分割します。"
}

func main() {
	var ssg string
	var prefix string

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.StringVar(&ssg, "ssg", "", "output path for static site generator")
	flags.StringVar(&prefix, "prefix", "", "prefix for deploy destination")

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalln("parse flags:", err)
	}

	app.Route("/", &demo{})
	app.RunWhenOnBrowser()

	if ssg == "" {
		serve()
	} else {
		generate(ssg, prefix)
	}
}

func serve() {
	http.Handle("/", &app.Handler{
		Name:        "BudouX-Go Demo",
		Description: "Demo page for BudouX-Go",
		Icon: app.Icon{
			Default: "https://github.com/identicons/sg0hsmt.png",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}
}

func generate(dist, prefix string) {
	var res app.ResourceProvider
	if prefix != "" {
		res = app.GitHubPages(prefix)
	}

	err := app.GenerateStaticWebsite(dist, &app.Handler{
		Name:        "BudouX-Go Demo",
		Description: "Demo page for BudouX-Go",
		Icon: app.Icon{
			Default: "https://github.com/identicons/sg0hsmt.png",
		},
		Resources: res,
	})

	if err != nil {
		log.Fatalln(err)
	}
}
