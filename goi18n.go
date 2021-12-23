package goi18n

import (
	"context"
	"sync"
)

type locale struct {
	mu          sync.RWMutex
	id          int64
	lang        string
	translation map[string]string
}

type Goi18n struct {
	Mu        sync.RWMutex
	Langs     []string
	LangDescs map[string]string
	LocaleMap map[string]*locale
}

func New() *Goi18n {
	return &Goi18n{
		Mu:        sync.RWMutex{},
		Langs:     []string{},
		LangDescs: map[string]string{},
		LocaleMap: map[string]*locale{},
	}
}

func (g *Goi18n) init(ctx context.Context) {

}

func (g *Goi18n) SetLangDesc(lang string, desc string) {
	g.LangDescs[lang] = desc
}

func (g *Goi18n) SetLanguage(lang string, filename string, desc string) bool {
	return true
}

func (g *Goi18n) Translate(lang string, key string) string {
	return ""
}

func (g *Goi18n) GetLanguageLength() int {
	return len(g.Langs)
}
