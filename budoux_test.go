package budoux_test

import (
	"reflect"
	"testing"

	"github.com/sg0hsmt/budoux-go"
	"github.com/sg0hsmt/budoux-go/models"
)

func TestParse(t *testing.T) {
	tbl := []struct {
		in  string
		out []string
	}{
		{
			in:  "",
			out: []string{""},
		},
		{
			in:  "日本語",
			out: []string{"日本語"},
		},
		{
			in:  "水と油",
			out: []string{"水と油"},
		},
		{
			in:  "水道水とミネラルウォーター",
			out: []string{"水道水と", "ミネラルウォーター"},
		},
		{
			in:  "PythonとJavaScriptとGolang",
			out: []string{"Pythonと", "JavaScriptと", "Golang"},
		},
		{
			in:  "日本語の文章において語の区切りに空白を挟んで記述すること",
			out: []string{"日本語の", "文章に", "おいて", "語の", "区切りに", "空白を", "挟んで", "記述する", "こと"},
		},
		{
			in:  "これはテストです。",
			out: []string{"これは", "テストです。"},
		},
		{
			in:  "これは美しいペンです。",
			out: []string{"これは", "美しい", "ペンです。"},
		},
		{
			in:  "今日は天気です。",
			out: []string{"今日は", "天気です。"},
		},
		{
			in:  "今日はとても天気です。",
			out: []string{"今日は", "とても", "天気です。"},
		},
		{
			in:  "あなたに寄り添う最先端のテクノロジー。",
			out: []string{"あなたに", "寄り添う", "最先端の", "テクノロジー。"},
		},
		{
			in:  "これはテストです。今日は晴天です。",
			out: []string{"これは", "テストです。", "今日は", "晴天です。"},
		},
		{
			in:  "これはテストです。\n今日は晴天です。",
			out: []string{"これは", "テストです。", "\n今日は", "晴天です。"},
		},
	}

	model := models.DefaultJapaneseModel()
	parser := budoux.New(model, budoux.DefaultThreshold)

	for _, v := range tbl {
		t.Run(v.in, func(t *testing.T) {
			out := parser.Parse(v.in)
			if !reflect.DeepEqual(out, v.out) {
				t.Errorf("unmatch result:\nwant %q\n got %q", v.out, out)
			}
		})
	}
}
