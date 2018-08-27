echo "creating coverage files"
if not exist "%TEMP%\coverage" mkdir %TEMP%\coverage
go test -coverprofile=%TEMP%\coverage\coverage.out
go tool cover -html=%TEMP%\coverage\\coverage.out -o %TEMP%\coverage\\coverage.html
echo "coverage files created"
