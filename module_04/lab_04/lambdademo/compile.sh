 
 GOOS=linux GOARCH=amd64 go build -o lambdademo main.go \
 && zip lambdademo.zip lambdademo \
 && rm lambdademo
