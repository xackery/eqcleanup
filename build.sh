echo Building Linux
GOOS=linux GOARCH=amd64 go build main.go
mv main bin/eqcleanup-linux64
GOOS=linux GOARCH=386 go build main.go
mv main bin/eqcleanup-linux386
echo Building Windows
GOOS=windows GOARCH=amd64 go build main.go
mv main.exe bin/eqcleanup-windows64.exe
GOOS=windows GOARCH=386 go build main.go
mv main.exe bin/eqcleanup-windows386.exe
echo Building OSX
GOOS=darwin GOARCH=amd64 go build main.go
mv main bin/eqcleanup-osx-64