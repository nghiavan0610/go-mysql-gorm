# go-mysql-gorm

This is an example for Go application project using MySQL and Gorm.

## Install go(lang)

with [homebrew](http://mxcl.github.io/homebrew/):

```Shell
sudo brew install go
```

with [apt](http://packages.qa.debian.org/a/apt.html)-get:

```Shell
sudo apt-get install golang
```

[install Golang manually](https://golang.org/doc/install)
or
[compile it yourself](https://golang.org/doc/install/source)

## Clone the project

```
git clone https://github.com/nghiavan0610/go-mysql-gorm.git
```

## Start MySQL using Docker

```
docker-compose up -d
```

## Start Application

```
go run ./cmd/main/main.go
```

## Compile

```
go build ./cmd/main

./main
```
