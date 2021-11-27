NAME = box-tailor
GOOS = linux

run:
	go run .
dev:
	$(eval NAME = $(NAME)_debug)
windows:
	$(eval GOOS = windows)
binary:
	GOOS=$(GOOS) go build -o ./bin/$(NAME).exe .
test:
	go run . -o tmp/test -d 100,50,25
