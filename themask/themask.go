package themask

import (
	"io/ioutil"
	"os"
)

type RuleSet []Engine

type TheMask struct {
	engines []Engine
}

func NewTheMask(engines []Engine) TheMask {
	themask := TheMask{
		engines,
	}
	return themask
}

func (tm TheMask) Smoking(text string) string {
	logger.Debugf("Start Smoking")
	logger.Debugp("engines dump:", tm.engines)
	for _, engine := range tm.engines {
		text = engine.Replace(text)
	}
	return text
}

func (tm TheMask) SmokingFile(infilename string, outfilename string) error {
	logger.Debugf("Start SmokingFile")
	logger.Debugf("infilename:  %s", infilename)
	logger.Debugf("outfilename: %s", outfilename)

	rfp, err := os.Open(infilename)
	if err != nil {
		return err
	}
	defer rfp.Close()

	buf, err := ioutil.ReadAll(rfp)
	if err != nil {
		return err
	}
	text := string(buf)
	text = tm.Smoking(text)
	ofp, err := os.OpenFile(outfilename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer ofp.Close()
	ofp.WriteString(text)
	return nil
}
