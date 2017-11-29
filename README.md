# gam [![Build Status](https://travis-ci.org/ademozay/gam.svg?branch=master)](https://travis-ci.org/ademozay/gam)

gam is a CLI tool to create, update and delete bash aliases permanently.

### Installing
* Create a file `.gam_aliases` in home directory.

* Source `.gam_aliases` in startup script(`~/.bashrc`, `~/.zshrc`).
  ```bash
  source ~/.gam_aliases
  ```

* `go get -u github.com/ademozay/gam`


#### Samples

* Create an alias (First parameter is alias name, value is the rest)

  `gam gitlab ssh admin@10.0.8.13`

* Update an alias (First parameter is alias name, value is the rest)

  `gam gitlab ssh admin@10.0.8.14`

* Delete an alias

  `gam -d gitlab`

* Print an alias

  `gam gitlab`

* Print all aliases created by gam

  `gam`

Either run `source ~/.bashrc` or `source ~/.zshrc` according to current shell or open a new terminal for changes to take place.

#### Contribution

All contributions are welcome.

##### TODO

- [ ] Source bash in current session. (Find a way to find the right startup script(`~/.bashrc`, `~/.zshrc`).
