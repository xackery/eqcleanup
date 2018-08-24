export VERSION="0.15"
export NAME="eqcleanup"
echo Building Linux $VERSION
GOOS=linux GOARCH=amd64 go build mainmenu.go
mv main bin/$NAME-$VERSION-linux-x64
GOOS=linux GOARCH=386 go build mainmenu.go
mv main bin/$NAME-$VERSION-linux-x86

echo Building Windows $VERSION
GOOS=windows GOARCH=amd64 go build mainmenu.go
mv main.exe bin/$NAME-$VERSION-windows-x64.exe
GOOS=windows GOARCH=386 go build mainmenu.go
mv main.exe bin/$NAME-$VERSION-windows-x86.exe

echo Building OSX $VERSION
GOOS=darwin GOARCH=amd64 go build mainmenu.go
mv main bin/$NAME-$VERSION-osx-x64
