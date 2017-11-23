package main

import (
	"errors"
	"fmt"
)

type alias struct {
	name, value string
}

func (a alias) create() error {
	aliases, err := readAll()
	if err != nil {
		return err
	}
	if aliases.exists(a.name) {
		alias := alias{name: a.name, value: aliases[a.name]}
		return fmt.Errorf("exists: %s", alias.string())
	}

	err = aliases.append(a)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("created: %s", a.string()))
	return nil
}

func (a alias) update() error {
	aliases, err := readAll()
	if err != nil {
		return err
	}
	if !aliases.exists(a.name) {
		return errors.New("no such alias")
	}
	if aliases[a.name] == a.value {
		return errors.New("no changes")
	}

	oldAlias := alias{name: a.name, value: aliases[a.name]}
	newAlias := alias{name: a.name, value: a.value}

	err = aliases.modify(a.name, a.value)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("old: %s\nnew: %s", oldAlias.string(), newAlias.string()))
	return nil
}

func (a alias) delete() error {
	aliases, err := readAll()
	if err != nil {
		return err
	}
	if !aliases.exists(a.name) {
		return errors.New("no such alias")
	}

	a.value = aliases[a.name]

	err = aliases.remove(a.name)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("deleted: %s", a.string()))
	return nil
}

func (a alias) string() string {
	return fmt.Sprintf("alias %s=\"%s\"", a.name, a.value)
}
