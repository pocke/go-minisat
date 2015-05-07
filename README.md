go-minisat
=================

Golang binding for minisat.



Installation
----------------

### Install minisat


Arch linux

```sh
$ yaourt -S minisat-git
```

Other

```sh
$ git clone https://github.com/niklasso/minisat
$ cd minisat
$ make config prefix="/usr"
$ make
$ sudo make install
```

### Install go-minisat

```sh
$ go get github.com/pocke/go-minisat
```


Usage
-----------


```go
package main

import (
  "github.com/pocke/go-minisat"
)

func main() {
  s := minisat.NewSolver()
  v1 := minisat.NewVar()
  v2 := minisat.NewVar()
  s.AddClause(v1, v2)
  s.AddClause(v1, v2.Not())
  s.AddClause(v1.Not(), v2.Not())
  s.Solve()  // => true

  s.ModelValue(v1)  // => true, nil
  s.ModelValue(v2)  // => false, nil
}
```


Example
------------

- [sudoku-solver](https://github.com/pocke/sudoku-solver)


License
-----------

Copyright &copy; 2015 Masataka Kuwabara
Licensed [MIT][mit]
[MIT]: http://www.opensource.org/licenses/mit-license.php
