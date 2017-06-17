echo "Generating..."
go run *.go -id=48 -name=cazicthule
go run *.go -id=18 -name=paw
go run *.go -id=108 -name=veeshan
go run *.go -id=107 -name=nurga
go run *.go -id=81 -name=droga
cd ../../
echo "Formatting..."
gofmt -s -w .
echo "Done."