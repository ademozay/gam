package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type alias struct {
	name, value string
}

type aliases map[string]string

func readAll() (aliases, error) {
	aliases := make(aliases)

	file, err := os.Open(aliasFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), "alias", "", 1)
		line = strings.TrimSpace(line)
		parts := strings.Split(line, "=")
		aliases[parts[0]] = strings.Trim(parts[1], "\"")
	}
	return aliases, nil
}

func (a alias) exists(aliases aliases) bool {
	_, exists := aliases[a.name]
	return exists
}

func (a alias) create() error {
	aliases, err := readAll()
	if err != nil {
		return err
	}
	if a.exists(aliases) {
		alias := alias{name: a.name, value: aliases[a.name]}
		return fmt.Errorf("%s exists", alias.string())
	}

	file, err := os.OpenFile(aliasFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(a.string() + "\n"); err != nil {
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
	if !a.exists(aliases) {
		return errors.New("no such alias")
	}

	var oldAlias alias
	var newAlias alias
	var fileContent bytes.Buffer

	for name := range aliases {
		if name == a.name {
			if aliases[name] == a.value {
				return errors.New("no changes")
			}
			oldAlias = alias{name: name, value: aliases[name]}
			newAlias = alias{name: name, value: a.value}
			aliases[name] = a.value
		}
		alias := alias{name: name, value: aliases[name]}
		fileContent.WriteString(alias.string() + "\n")
	}

	err = ioutil.WriteFile(aliasFile, fileContent.Bytes(), 0600)
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
	if !a.exists(aliases) {
		return errors.New("no such alias")
	}
	a.value = aliases[a.name]
	delete(aliases, a.name)

	var fileContent bytes.Buffer
	for name := range aliases {
		alias := alias{name: name, value: aliases[name]}
		fileContent.WriteString(alias.string() + "\n")
	}

	err = ioutil.WriteFile(aliasFile, fileContent.Bytes(), 0600)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("deleted: %s", a.string()))
	return nil
}

func (a alias) print() error {
	aliases, err := readAll()
	if err != nil {
		return err
	}
	if !a.exists(aliases) {
		return errors.New("no such alias")
	}
	alias := alias{name: a.name, value: aliases[a.name]}
	fmt.Println(alias.string())
	return nil
}

func (a alias) printAll() error {
	aliases, err := readAll()
	if err != nil {
		return err
	}
	for name := range aliases {
		alias := alias{name: name, value: aliases[name]}
		fmt.Println(alias.string())
	}
	return nil
}

func (a alias) string() string {
	return fmt.Sprintf("alias %s=\"%s\"", a.name, a.value)
}
