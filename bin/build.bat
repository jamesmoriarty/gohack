go version

set PATH=%PATH%;C:\MinGW\bin

set GOARCH=386
set GOOS=windows
set CGO_ENABLED=1

cd .\test\dll
go build -buildmode=c-archive client_panorama.go
gcc -shared -pthread -o client_panorama.dll client_panorama.a -lWinMM -lntdll -lWS2_32
gcc client_panorama.def client_panorama.a -shared -lwinmm -lWs2_32 -o client_panorama.dll -Wl,--out-implib,client_panorama.lib
go build -v -o csgo.exe main.go
cd ..\..

cd .\test\nodll
go build -v -o csgo.exe main.go
cd ..\..

go test -v -coverprofile cover.out

set GOARCH=amd64

for /f %%a in ('powershell -Command "git rev-parse --short HEAD"') do set VERSION=%%a
for /f %%a in ('powershell -Command "Get-Date -format yyyyMMdd.HHmmss"') do set DATE=%%a

go build -v -ldflags "-X github.com/jamesmoriarty/gohack.Version=%VERSION% -X github.com/jamesmoriarty/gohack.Date=%DATE%" cmd/gohack.go
