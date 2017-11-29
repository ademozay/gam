package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type alias struct {
	name, value string
}

func (a alias) persist() error {
	aliases, err := readAll()
	if err != nil {
		return err
	}

	if aliases.exists(a.name) {
		if aliases[a.name] == a.value {
			return nil
		}
		old := alias{name: a.name, value: aliases[a.name]}
		if yes := old.shouldReplaceWith(a); yes {
			return aliases.modify(a)
		}
		return nil
	}
	return aliases.append(a)
}

func (a alias) remove() error {
	aliases, err := readAll()
	if err != nil {
		return err
	}
	return aliases.remove(a.name)
}

func (a alias) string() string {
	return fmt.Sprintf(`alias %s="%s"`, a.name, a.value)
}

func (a alias) shouldReplaceWith(new alias) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("old value: %s\nnew value: %s\ndo you want to update? [y/n] ", a.value, new.value)
	update, _ := reader.ReadString('\n')
	update = strings.Replace(update, "\n", "", -1) // TODO
	if update == "n" {
		return false
	}
	return true
}
