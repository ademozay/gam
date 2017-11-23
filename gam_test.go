package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateTmpAliasFile(t *testing.T) {
	aliasFile = "/tmp/.gam_profile"
	err := ioutil.WriteFile(aliasFile, []byte{}, 0644)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateAlias(t *testing.T) {
	alias := alias{name: "foo", value: "bar"}
	err := alias.create()
	if err != nil {
		t.Error(err)
	}
}

func TestReadAlias(t *testing.T) {
	alias, err := readOne("foo")
	if err != nil {
		t.Error(err)
	}
	if alias.value != "bar" {
		t.Error(`expected "bar" as alias value`)
	}

}
func TestUpdateAlias(t *testing.T) {
	alias := alias{name: "foo", value: "baz"}
	err := alias.update()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteAlias(t *testing.T) {
	alias := alias{name: "foo"}
	err := alias.delete()
	if err != nil {
		t.Error(err)
	}
}
func TestReadAll(t *testing.T) {
	aliases, err := readAll()
	if err != nil {
		t.Error(err)
	}
	if len(aliases) > 0 {
		t.Error("no any aliases expected")
	}
}

func TestDeleteAliasFile(t *testing.T) {
	err := os.Remove(aliasFile)
	if err != nil {
		t.Error(err)
	}
}
