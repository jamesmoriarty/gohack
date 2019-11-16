gofmt -s -w .
go get -u golang.org/x/lint/golint
golint .
echo Exit Code is %errorlevel%
go.exe build -v