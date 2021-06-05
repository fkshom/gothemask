package main

import (
	"flag"
	"gothemask/themask"
	"log"
)

func main() {
	var (
		typename = flag.String("type", "common", "type name")
		debug    = flag.Bool("debug", false, "Debug Mode enabled?")
	)
	flag.Parse()
	args := flag.Args()
	if *debug {
		themask.SetLevel(themask.DEBUG)
	} else {
		themask.SetLevel(themask.WARN)
	}

	_config := themask.NewConfig("config/config.yaml")
	config := themask.ResolveConfig(_config)
	_, ok := config[*typename]
	if !ok {
		log.Fatalln("typaname " + *typename + " not defined.")
	}
	themask := themask.NewTheMask(config[*typename])
	infilename := ""
	outfilename := ""
	if len(args) == 1 {
		infilename = args[0]
	} else if len(args) == 2 {
		infilename = args[0]
		outfilename = args[1]
	} else {
		log.Fatalln("USAGE: ./a.out infile [outfile]")
	}
	themask.SmokingFile(infilename, outfilename)
}
