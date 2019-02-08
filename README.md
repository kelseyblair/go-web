A very basic implementation of a Go web application to receive and store an image file.

go build webthing.go

go run webthing.go

localhost:8080

curl -X POST -F 'image=@/path/to/imagefile' localhost:8080/upload

