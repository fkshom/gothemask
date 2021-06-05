package main

import (
	"fmt"
	"gothemask/themask"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	_config := themask.NewConfig("config/config.yaml")
	config := themask.ResolveConfig(_config)
	for typename := range _config.Test_rules {
		_, ok := config[typename]
		if !ok {
			fmt.Fprintln(os.Stderr, "WARNING: typename "+typename+" not found in typemap")
			continue
		}
		themask := themask.NewTheMask(config[typename])
		testsets := _config.Test_rules[typename]
		for idx, testset := range testsets {
			text := testset.Text
			expect := testset.Expect
			actual := themask.Smoking(text)
			if actual != expect {
				t.Errorf("\n[typename] %s\n[testnum] %d\n[actual]\n%v\n[expect]\n%v", typename, idx, actual, expect)
			}
		}
	}
}
