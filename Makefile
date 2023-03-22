# Define the compiler and compiler flags
GO=go
GOFLAGS=-ldflags="-s -w"

# Define the source files
SRC=$(wildcard *.go)

# Define the output binary name
BIN=myprogram

# Define the default target
all: $(BIN)

# Define the build target
$(BIN): $(SRC)
	$(GO) build $(GOFLAGS) -o $(BIN) $(SRC)

# Define the clean target
clean:
	rm -f $(BIN)