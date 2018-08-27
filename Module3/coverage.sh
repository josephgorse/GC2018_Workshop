echo "creating coverage files"
go test -coverprofile=/tmp/coverage.out
go tool cover -html=/tmp/coverage.out -o /tmp/coverage.html
echo "coverage files created"
