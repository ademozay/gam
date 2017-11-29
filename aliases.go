package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

type aliases map[string]string

func readAll() (aliases, error) {
	file, err := os.Open(aliasFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	aliases := make(aliases)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), "alias", "", 1)
		line = strings.TrimSpace(line)
		parts := strings.Split(line, "=")
		aliases[parts[0]] = strings.Trim(parts[1], "\"")
	}
	return aliases, nil
}

func readOne(name string) (alias, error) {
	aliases, err := readAll()
	if err != nil {
		return alias{}, err
	}
	if !aliases.exists(name) {
		return alias{}, errNoSuchAlias
	}
	alias := alias{name: name, value: aliases[name]}
	return alias, nil
}

func (aliases aliases) exists(name string) bool {
	if _, exists := aliases[name]; exists {
		return true
	}
	return false
}

func (aliases aliases) append(alias alias) error {
	file, err := os.OpenFile(aliasFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(alias.string() + "\n"); err != nil {
		return err
	}
	return nil
}

func (aliases aliases) modify(alias alias) error {
	if !aliases.exists(alias.name) {
		return errNoSuchAlias
	}
	aliases[alias.name] = alias.value
	return aliases.persist()
}

func (aliases aliases) remove(name string) error {
	if !aliases.exists(name) {
		return errNoSuchAlias
	}
	delete(aliases, name)
	return aliases.persist()
}

func (aliases aliases) persist() error {
	var out bytes.Buffer
	for name := range aliases {
		alias := alias{name: name, value: aliases[name]}
		out.WriteString(alias.string() + "\n")
	}
	return ioutil.WriteFile(aliasFile, out.Bytes(), 0600)
}

func (aliases aliases) strings() []string {
	var s []string
	for name := range aliases {
		alias := alias{name: name, value: aliases[name]}
		s = append(s, alias.string())
	}
	return s
}
