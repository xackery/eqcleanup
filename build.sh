echo Building Linux
GOOS=linux GOARCH=amd64 go build main.go
mv main bin/linux/amd64/eqcleanup
GOOS=linux GOARCH=386 go build main.go
mv main bin/linux/386/eqcleanup
echo Building Windows
GOOS=windows GOARCH=amd64 go build main.go
mv main.exe bin/windows/amd64/eqcleanup.exe
GOOS=windows GOARCH=386 go build main.go
mv main.exe bin/windows/386/eqcleanup.exe
echo Building OSX
GOOS=darwin GOARCH=amd64 go build main.go
mv main bin/darwin/amd64/eqcleanup