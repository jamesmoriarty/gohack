SETLOCAL ENABLEDELAYEDEXPANSION

go version
gofmt -s -w .
echo Exit Code is %errorlevel%

go get -u golang.org/x/lint/golint
golint .
echo Exit Code is %errorlevel%

go test -coverprofile cover.out
echo Exit Code is %errorlevel%

for /f %%a in ('powershell -Command "git rev-parse --short HEAD"') do set VERSION=%%a
for /f %%a in ('powershell -Command "Get-Date -format yyyyMMdd.HHmmss"') do set DATE=%%a
go.exe build -v -ldflags "-X github.com/jamesmoriarty/gohack/config.Version=%VERSION% -X github.com/jamesmoriarty/gohack/config.Date=%DATE%"