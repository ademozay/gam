package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var usr, _ = user.Current()
var aliasFile = filepath.Join(usr.HomeDir, ".gam_aliases")

const usage = `

usage:
	gam [options] [...arguments]

options:
	-d	delete given alias

examples:
	create alias. first arg is name of the alias, value is the rest
		gam gitlab ssh admin@10.0.8.13

	update alias.
		gam gitlab ssh admin@10.0.8.14

	delete alias.
		gam -d gitlab

	print alias.
		gam gitlab
	
	print all aliases.
		gam

`

type action int

const (
	persist action = iota
	del
	print
	printAll
)

type gam struct {
	action action
	alias  alias
}

func (g gam) execute() error {
	switch g.action {
	case persist:
		if g.alias.name == "" {
			return errMissingName
		}
		if g.alias.value == "" {
			return errMissingValue
		}
		return g.alias.persist()
	case del:
		return g.alias.remove()
	case print:
		alias, err := readOne(g.alias.name)
		if err != nil {
			return err
		}
		fmt.Println(alias.string())
		return nil
	case printAll:
		aliases, err := readAll()
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(aliases.strings(), "\n"))
		return nil
	default:
		return errInvalidAction
	}
}

func main() {
	delete := flag.String("d", "", "")
	help := flag.Bool("h", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
	}
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	var args = flag.Args()
	var action action
	var alias alias

	if *delete != "" {
		action = del
		alias.name = *delete
	} else if len(args) == 0 {
		action = printAll
	} else if len(args) == 1 {
		action = print
		alias.name = args[0]
	} else if len(args) > 1 {
		action = persist
		alias.name = args[0]
		alias.value = strings.Join(args[1:], " ")
	}

	gam := &gam{
		action: action,
		alias:  alias,
	}
	err := gam.execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
