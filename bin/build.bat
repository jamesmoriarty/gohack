go version

REM configure csgo.exe dist is 386
set GOARCH=386
set GOOS=windows
set CGO_ENABLED=1
set PATH=%PATH%;C:\MinGW\bin

REM build .\test\dll\csgo.exe
cd .\test\dll
gcc -c -o client.o client.c
gcc -o client.dll -s -shared client.o -Wl,--subsystem,windows
gcc -o csgo csgo.c -L. -lclient
cd ..\..

REM build .\test\nodll\csgo.exe
cd .\test\nodll
go build -v -o csgo.exe main.go
cd ..\..

REM build .\gohack.exe
set GOARCH=amd64

for /f %%a in ('powershell -Command "git rev-parse --short HEAD"') do set VERSION=%%a
for /f %%a in ('powershell -Command "Get-Date -format yyyyMMdd.HHmmss"') do set DATE=%%a

go build -v -ldflags "-X github.com/jamesmoriarty/gohack.Version=%VERSION% -X github.com/jamesmoriarty/gohack.Date=%DATE%" cmd/gohack.go

if %errorlevel% neq 0 exit /b %errorlevel%

REM test
go test -covermode=atomic -coverprofile='cover.out' -v ./...
