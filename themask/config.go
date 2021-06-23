package themask

import (
	"log"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	jsoniter "github.com/json-iterator/go"
)

type Config struct {
	Rules      map[string][]map[string]interface{}
	Typemap    map[string][]string
	Test_rules map[string][]struct {
		Text   string
		Expect string
	}
}

type Engineconfig struct {
	Engine string
}

func merge(m1, m2 map[string]string) map[string]string {
	ans := map[string]string{}

	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
}

func MapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := jsoniter.Marshal(m)
	if err != nil {
		return err
	}
	err = jsoniter.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}

func NewConfigFromDir(dirname string) (Config, error) {
	pattern := filepath.Join(dirname, "*.yaml")
	files, err := filepath.Glob(pattern)
	config := Config{
		Rules:   make(map[string][]map[string]interface{}),
		Typemap: make(map[string][]string),
		Test_rules: make(map[string][]struct {
			Text   string
			Expect string
		}),
	}
	if err != nil {
		return config, err
	}
	for _, filename := range files {
		log.Println("Loading " + filename + "...")
		_config := NewConfig(filename)
		for k, v := range _config.Rules {
			config.Rules[k] = v
		}
		for k, v := range _config.Typemap {
			config.Typemap[k] = v
		}
		for k, v := range _config.Test_rules {
			config.Test_rules[k] = v
		}
	}
	return config, nil
}

func NewConfig(filename string) Config {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fp.Close()
	d := yaml.NewDecoder(fp)

	var config Config
	err = d.Decode(&config)
	if err != nil {
		logger.Warnf("cant decode from yaml")
		log.Fatalln(err)
	}
	return config
}

func ResolveConfig(config Config) map[string][]Engine {
	var err error

	result_rules := map[string][]Engine{}
	for rulename, rules := range config.Rules {
		for _, rule := range rules {
			var ee Engineconfig
			err = MapToStruct(rule, &ee)
			if err != nil {
				log.Fatalln(err)
			}

			switch ee.Engine {
			case "regexp":
				var v3 Regexp
				err = MapToStruct(rule, &v3)
				if err != nil {
					log.Fatalln(err)
				}
				result_rules[rulename] = append(result_rules[rulename], v3)
			case "plaintext":
				var v3 PlainText
				err = MapToStruct(rule, &v3)
				if err != nil {
					log.Fatalln(err)
				}
				result_rules[rulename] = append(result_rules[rulename], v3)
			case "substringregexp":
				var v3 SubstringRegexp
				err = MapToStruct(rule, &v3)
				if err != nil {
					log.Fatalln(err)
				}
				result_rules[rulename] = append(result_rules[rulename], v3)
			default:
				logger.Warnf("Invalid engin name in %s", rulename)
			}
		}
	}

	result := map[string][]Engine{}
	for typename, rulenames := range config.Typemap {
		for _, rulename := range rulenames {
			rules, ok := result_rules[rulename]
			if !ok {
				logger.Warnf("WARNING: rulename %s not found.", rulename)
			}
			result[typename] = append(result[typename], rules...)
		}
	}
	logger.Debugp("loaded config:", result)
	return result
}
