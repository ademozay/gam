# gam

gam is a CLI tool to create, update and delete bash aliases permanently.

### Installing
* Create a file `.gam_aliases` in home directory.

* Source `.gam_aliases` in bash.
  ```bash
  if [ -f ~/.gam_aliases ]; then
          . ~/.gam_aliases
  fi
  ```

* `go get github.com/ademozay/gam`



#### Samples

*Following commands require new shell session in order to source startup script.*

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


#### Contribution

All contributions are welcome.

##### TODO

- [ ] Source bash in current session. (Find a way to find the right startup script(`.bashrc`, `.zshrc`).
