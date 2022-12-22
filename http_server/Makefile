PORT = 5000

install:
	go install github.com/cortesi/modd/cmd/modd@latest
	go mod tidy

dev:
	-fuser -k ${PORT}/tcp
	PORT=${PORT} modd
