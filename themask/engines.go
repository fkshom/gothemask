package themask

import (
	"github.com/dlclark/regexp2"
)

type Engine interface {
	Replace(string) string
}

type PlainText struct {
	Patterns []string
	To       string
}

func NewPlainText(param map[string]interface{}) PlainText {
	conf := PlainText{
		param["patterns"].([]string),
		param["to"].(string),
	}
	return conf
}

func (conf PlainText) Replace(text string) string {
	for _, pattern := range conf.Patterns {
		re := regexp2.MustCompile(regexp2.Escape(pattern), regexp2.Multiline)
		text, _ = re.Replace(text, conf.To, -1, -1)
	}
	return text
}

type Regexp struct {
	Patterns []string
	To       string
}

func NewRegexp(param map[string]interface{}) Regexp {
	conf := Regexp{
		param["patterns"].([]string),
		param["to"].(string),
	}
	return conf
}

func (conf Regexp) Replace(text string) string {
	for _, pattern := range conf.Patterns {
		re := regexp2.MustCompile(pattern, regexp2.Multiline)
		text, _ = re.Replace(text, conf.To, -1, -1)
	}
	return text
}

type SubstringRegexp struct {
	Patterns []string
	To       map[string]string
}

func NewSubstringRegexp(param map[string]interface{}) SubstringRegexp {

	conf := SubstringRegexp{
		param["patterns"].([]string),
		param["to"].(map[string]string),
	}
	return conf
}

func (conf SubstringRegexp) Replace(text string) string {
	for _, pattern := range conf.Patterns {
		re := regexp2.MustCompile(pattern, regexp2.Multiline)
		text, _ = re.ReplaceFunc(text, func(m regexp2.Match) string {
			innertext := m.String()
			re2 := regexp2.MustCompile(pattern, regexp2.Multiline)
			m2, _ := re2.FindStringMatch(innertext)
			for i := m2.GroupCount() - 1; i > 0; i-- {
				name := m2.GroupByNumber(i).Name
				idx := m2.GroupByNumber(i).Index
				len := m2.GroupByNumber(i).Length
				to := conf.To[name]
				innertext = innertext[:idx] + to + innertext[idx+len:]
			}
			return innertext
		}, -1, -1)
	}
	return text
}
