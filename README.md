# goblog

A fast and simple blog with [Revel](http://revel.github.io/) framework in Golang

Current version is **0.1.0** on 2015.01.10

You can test the latest version at http://goblog.jang.io.

## Getting started

```
$ revel run goblog
```

http://localhost:9000/

## Usage

### 1. Install Go

https://golang.org/doc/install

### 2. Install packages

```
$ go get github.com/revel/revel
$ go get go get github.com/revel/cmd/revel
$ go get github.com/jinzhu/gorm
$ go get github.com/mattn/go-sqlite3
$ go get code.google.com/p/go.crypto/bcrypt
```

### 3. Get the code

```
$ cd $GOPATH/src
$ git clone git@github.com:jaehue/goblog.git
```

### 4. Run

```
$ revel run goblog
```
- Default user information
  - Username: admin
  - Name: Admin
  - Role: admin
  - Password: admin

**YOU MUST CHANGE DEFAULT PASSWROD**

## Deploy to Heroku

### Setup

1. Create .godir
  ```
  $ echo goblog > .godir
  
  $ git add .godir
  
  $ git commit .godir -m 'add .godir'
  ...
  ```
2. Remove `routes/` in .gitignore file
  - .gitignore file   
  ```
  test-results/
  tmp/
  ```
3. Add routes.go file
  ```
  $ git add app/routes/routes.go
  
  $ $ git commit -a -m 'Add routes.go file'
  ...
  ```
4. Create heroku app
  ```
  $ heroku create -b https://github.com/revel/heroku-buildpack-go-revel.git
  ...
  Git remote heroku added
  ```
  
5. Create database
  ```
  $ heroku addons:add heroku-postgresql:hobby-dev
  ...
  
  $ heroku config | grep HEROKU_POSTGRESQL
  HEROKU_POSTGRESQL_GRAY_URL: [YOUR_DATABASE_CONNECTION_URL]
  ```
  **HEROKU_POSTGRESQL_GRAY_URL** is your database connection url
6. Set production database settings in app.conf file
  - conf/app.conf
  ```
  [prod]
  ...
  db.import = github.com/lib/pq
  db.driver = postgres
  db.spec   = ${HEROKU_POSTGRESQL_GRAY_URL}
  ...
  ```

Complete!
  
  

### Deploy

```
$ git push heroku master
Counting objects: 350, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (334/334), done.
Writing objects: 100% (350/350), 118.16 KiB | 0 bytes/s, done.
Total 350 (delta 165), reused 0 (delta 0)
remote: Compressing source files... done.
remote: Building source:
remote:
remote: -----> Fetching custom git buildpack... done
remote: -----> Revel app detected
remote: -----> Installing go1.3... done
remote:        Installing Virtualenv... done
remote:        Installing Mercurial... done
...
```

## Contributing

1. Fork it ( https://github.com/jaehue/goblog/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## Questions

Create issues or pull requests here.
