# urlendecode
go cli app for url encode/decode
## Usage
> NOTE: wrap strings in single quotes  

urlendecode 
  -d string
        pass the string to decode
  -e string
        pass the string to encode

## Build
`go build .`

### Build examples for other architectures and OS'

#### Linux x86\_64

`env GOOS=linux GOARCH=amd64 go build .`

#### MacOS with arm CPU

`env GOOS=darwin GOARCH=arm64 go build .`

#### Windows

`env GOOS=windows go build .`


## Deploy to your user
`cp urlendecode ~/.local/bin/`


