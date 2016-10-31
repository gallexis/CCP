# CCP
Car communication Protocol in Golang


###Cross compile to raspberry 1 (& 2?): 
env GOOS=linux GOARCH=arm GOARM=6 go build -v

###Cross compile to raspberry 3:
env GOOS=linux GOARCH=arm64 GOARM=7 go build -v
