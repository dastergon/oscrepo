# oscrepo
Get .repo URLs in openSUSE without opening the browser.

## Installation

> You can install Go by following [these instructions](https://golang.org/doc/install).

`oscrepo` is written in Go, so if you have Go installed you can install it with `go get`:

    go get github.com/dastergon/oscrepo


This will download the code, compile it, and leave an `oscrepo` binary in `$GOPATH/bin`.

## Configuration

The openSUSE API requires to authenticate before you request resources from it.
Consequently, `oscrepo` parses the credentials either from the `$HOME/.oscrc` file of `osc` or through 
the `--username` and `--password` parameters.

## Usage

Show the available .repo URLs for projects that contain the word containers (using `.oscrc`):

    oscrepo url containers 

Another example:

     oscrepo url Virtualization:containers
  
Show the available .repo URLs for projects that contain the word containers (without `.oscrc`):

    oscrepo url containers --username john --password smith
  
If there are more than one entries that match the keyword, a enumerated list will be returned.
To select the desired URL use the `--entry` option:

     oscrepo url containers --entry 1
  
