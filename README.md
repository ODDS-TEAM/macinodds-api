# Install Golang

## OS X

##### Install brew

Install brew from [brew.sh](http://brew.sh/)

##### Install Git

```
$ brew install git 
```

##### Install Go 1.10+

Install golang binaries using `brew`

```
$ brew install go
```

##### Setup PATH

Add the PATH to your ``~/.bash_profile``. 

```
export PATH=${HOME}/go/bin:$PATH
```

##### Source the new environment

```
$ source ~/.bash_profile
```

##### Testing it all

```
$ go env
$ go version
```

##### Pull package

```
$ go get
```

##### Run

```
$ go run server.go
```