# gofind
Traverse all children directories on a file system looking for any file and print their contents.

## Usage
```
$ gofind --help
-n string
      Filename
-p string
      Starting Path (default "/home/testuser1")
```
```
$ gofind -n .mysql_history
Found .mysql_history in: /home/testuser1
show databases;
use testdatabase;
select * from testtable;
...
```
```
$ sudo su - && gofind -n .mysql_history -p /home
Found .mysql_history in: /home/testuser1
...

Found .mysql_history in: /home/testuser2
...
```

## Install
- You can build binaries from source if `go` is already installed:
  ```
  go get github.com/pmbenjamin/gofind
  ```
- Or download precompiled binaries from [here](https://github.com/pmbenjamin/gofind/releases)
