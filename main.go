package main

import (
	"fmt"
	"log"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
)

func main() {
	// A context for our experiment
	ctx := cuecontext.New()

	// The first arg to this program is the yaml file to validate
	dataFile := os.Args[1]

	// Load the PodSpec schema we will use for validation
	bis := load.Instances([]string{"cue.dev/x/k8s.io/api/core/v1"}, nil)
	pkgV := ctx.BuildInstance(bis[0])
	podSpec := pkgV.LookupPath(cue.ParsePath("#PodSpec"))

	// Load the yaml file to validate
	data, err := yaml.Extract(dataFile, nil)
	if err != nil {
		log.Fatal(err)
	}
	dataV := ctx.BuildFile(data)

	// Unify the yaml value with the schema and validate
	combined := podSpec.Unify(dataV)
	if err := combined.Validate(cue.Concrete(true)); err != nil {
		log.Fatal(err)
	}

	// If we go this far we are good!
	fmt.Printf("%s is good\n", dataFile)
}
