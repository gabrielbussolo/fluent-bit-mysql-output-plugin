# Define the compiler and compiler flags
GO=go
GOFLAGS=-ldflags="-s -w"

# Define the source files
SRC=$(wildcard *.go)

# Define the output binary name
BIN=mysql-output-plugin

# Define the default target
all: $(BIN)

# Define the build target
$(BIN): $(SRC)
	$(GO) build $(GOFLAGS) -o $(BIN) $(SRC)

# Define the clean target
clean:
	rm -f $(BIN)

container:
	docker build --progress=plain --platform linux/amd64  -t fluent-bit-mysql .

container-run:
	docker run -it --rm --name fluent-bit-mysql fluent-bit-mysql