echo "Generating..."
go run *.go -id=48 -name=cazicthule
go run *.go -id=18 -name=paw
cd ../../
echo "Formatting..."
gofmt -s -w .
echo "Done."