build: clean
	go fmt
	GOOS=windows GOARCH=amd64 go build -o whitenoise.exe main.go

clean:
	rm -f whitenoise.exe