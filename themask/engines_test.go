package themask

import (
	"testing"
)

func TestMain1(t *testing.T) {
	conf := PlainText{
		Patterns: []string{
			`p@ssw0rd`,
			`this\wis\dnumber`,
		},
		To: "XXXX",
	}
	text := `
	this is p@ssw0rd
	this this\wis\dnumber ok
	`
	actual := conf.Replace(text)
	expected := `
	this is XXXX
	this XXXX ok
	`
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestMain2(t *testing.T) {
	conf := Regexp{
		Patterns: []string{
			`p@ssw0rd`,
			`this\wis\dnumber`,
		},
		To: "XXXX",
	}
	text := `
	this is p@ssw0rd
	this this\wis\dnumber ng
	this thisXis1number ok
	`
	actual := conf.Replace(text)
	expected := `
	this is XXXX
	this this\wis\dnumber ng
	this XXXX ok
	`
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestMain2a(t *testing.T) {
	conf := Regexp{
		Patterns: []string{
			`(?<pre>firewall )(?<mask>\w+)(?<post> {)`,
		},
		To: `${pre}XXXX${post}`,
	}
	text := `
	firewall mysecretfilter {
	`
	actual := conf.Replace(text)
	expected := `
	firewall XXXX {
	`
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestMain3(t *testing.T) {
	conf := SubstringRegexp{
		Patterns: []string{
			`firewall (?<mask>\w+) {`,
			`username (?<user>[^ ]+) password \d (?<pass>.+)`,
		},
		To: map[string]string{
			"mask": "XXXX",
			"user": "USERNAME",
			"pass": "PASSWORD",
		},
	}
	text := `
	firewall myfilter {
	username Bob password 0 abcdabcd
	`
	actual := conf.Replace(text)
	expected := `
	firewall XXXX {
	username USERNAME password 0 PASSWORD
	`
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
func TestMain4(t *testing.T) {
	conf := SubstringRegexp{
		Patterns: []string{
			`name: (?<name1>\w+) (?<name2>\w+)`,
		},
		To: map[string]string{
			"name1": "1XXX",
			"name2": "2XXX",
		},
	}
	text := `
	name: foo1 foo2   name: bar1 bar2
	`
	actual := conf.Replace(text)
	expected := `
	name: 1XXX 2XXX   name: 1XXX 2XXX
	`
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
