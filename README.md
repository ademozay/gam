# gam [![Build Status](https://travis-ci.org/ademozay/gam.svg?branch=master)](https://travis-ci.org/ademozay/gam)

gam is a CLI tool to create, update and delete bash aliases permanently.

### Installing
* Create a file `.gam_aliases` in home directory.

* Source `.gam_aliases` in startup script(`.bashrc`, `.zshrc`).
  ```bash
  source ~/.gam_aliases
  ```

* `go get github.com/ademozay/gam`



#### Samples

* create an alias

  `gam -c -n gitlab -v "ssh admin@10.0.8.13"`

* update an alias

  `gam -u -n gitlab -v "ssh admin@10.0.8.14"`

* delete an alias

  `gam -d gitlab`

* print an alias

  `gam -p gitlab`

* print all aliases created by gam

  `gam -P`

Either run `source ~/.bashrc` or `source ~/.zshrc` according to current shell or open a new terminal for changes to take place.

#### Contribution

All contributions are welcome.

##### TODO

- [ ] Source bash in current session. (Find a way to find the right startup script(`.bashrc`, `.zshrc`).
