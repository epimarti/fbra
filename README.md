# FBRA

The Fizz Buzz REST API


This is a simple project delivering a configurable Fizz Buzz list. The only purpose of this was to get my hands on Go development.

### Contents
* [Download](#download)
* [Deploy](#deploy)
* [Usage](#usage)
* [License](#license)

#### Download
* `go get github.com/epimarti/fbra`
* `go get github.com/gorilla/mux`
* `cd` into project directory
* `go build`

#### Deploy

* Local
  * run `go build`
  * run the application

* Docker
  * `docker build .`
  * `docker run fbra` _(you might have to map the port 80 to another one)_

#### Usage
This API serve only one resource which requires 5 parameters. It is reqested as follow:

`GET http://localhost/<int1>/<int2>/<limit>/<string1>/<string2>`

It will respond with an array of number from 1 to _limit_ where all multiples of _int1_ are replaced with _string1_, all multiples of _int2_ by _string2_ and all multiples of both _int1_ and _int2_ are replaced by '_string1 string2_'

#### License
This software is under the [Beerware License](LICENSE)
