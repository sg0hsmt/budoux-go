package main

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func TestDemo(t *testing.T) {
	compo := &demo{}

	disp := app.NewClientTester(compo)
	defer disp.Close()

	const initialValue = "日本語の文章をいい感じに分割します。"

	disp.Consume()
	if compo.input != initialValue {
		t.Fatalf("unmatch value, want %q, got %q", initialValue, compo.input)
	}

	if err := app.TestMatch(compo, app.TestUIDescriptor{
		Path: app.TestPath(0, 1, 0),
		Expected: app.Textarea().Text(initialValue).
			Placeholder("日本語の文章を入力してください。").
			Style("width", "100%").
			Style("height", "6em").
			OnInput(nil),
	}); err != nil {
		t.Fatalf("unmatch textarea: %v", err)
	}

	if err := app.TestMatch(compo, app.TestUIDescriptor{
		Path:     app.TestPath(0, 1, 1),
		Expected: app.Ol(),
	}); err != nil {
		t.Fatalf("unmatch list: %v", err)
	}

	for i, v := range []string{"日本語の", "文章を", "いい", "感じに", "分割します。"} {
		if err := app.TestMatch(compo, app.TestUIDescriptor{
			Path:     app.TestPath(0, 1, 1, i),
			Expected: app.Li(),
		}); err != nil {
			t.Fatalf("unmatch list item (%d): %v", i, err)
		}

		if err := app.TestMatch(compo, app.TestUIDescriptor{
			Path:     app.TestPath(0, 1, 1, i, 0),
			Expected: app.Text(v),
		}); err != nil {
			t.Fatalf("unmatch list item text (%d): %v", i, err)
		}
	}
}
