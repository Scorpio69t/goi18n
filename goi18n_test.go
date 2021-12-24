package goi18n

import (
	"fmt"
	"testing"
)

func TestGoi18n(t *testing.T) {
	g := New(nil)
	g.SetLanguage("zh", "中文")
	g.SetLanguage("en", "English")
	msg := g.Translate("zh", "title")
	t.Logf("msg: %s", msg)
	fmt.Printf("msg: %s", msg)
}
