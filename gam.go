package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

var usr, _ = user.Current()
var aliasFile = filepath.Join(usr.HomeDir, ".gam_aliases")

const usage = `
gam [...options]

options:
-c		enable create mode.
-u		enable update mode.
-n <string>	name of the alias to create or update.
-v <string>	value of the alias to create or update.
-d <string>	name of the alias to delete.
-p <string>	name of the alias to print.
-P		prints all aliases.

`

var (
	name     = flag.String("n", "", "")
	value    = flag.String("v", "", "")
	create   = flag.Bool("c", false, "")
	update   = flag.Bool("u", false, "")
	del      = flag.String("d", "", "")
	print    = flag.String("p", "", "")
	printAll = flag.Bool("P", false, "")
	help     = flag.Bool("h", false, "")
)

type gam struct {
	action string
	alias  alias
}

func (g gam) execute() error {
	switch g.action {
	case "create":
		if g.alias.name == "" || g.alias.value == "" {
			return errors.New("not enough arguments to create an alias")
		}
		return g.alias.create()
	case "update":
		if g.alias.name == "" || g.alias.value == "" {
			return errors.New("not enough arguments to update an alias")
		}
		return g.alias.update()
	case "delete":
		return g.alias.delete()
	case "print":
		return g.alias.print()
	case "printAll":
		return g.alias.printAll()
	default:
		return errors.New("invalid action. type -h or --help to see usage")
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
	}
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	var a alias
	var action string

	if *create {
		action = "create"
		a = alias{name: *name, value: *value}
	} else if *update {
		action = "update"
		a = alias{name: *name, value: *value}
	} else if *del != "" {
		action = "delete"
		a = alias{name: *del}
	} else if *print != "" {
		action = "print"
		a = alias{name: *print}
	} else if *printAll {
		action = "printAll"
	} else {
		flag.Usage()
		os.Exit(1)
	}

	gam := &gam{
		action: action,
		alias:  a,
	}
	err := gam.execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
