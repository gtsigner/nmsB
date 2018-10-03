# Name of executable and shared dll
MAIN = nmsB
# Version and Relase 
VERSION	= 0.0.1
RELEASE = $(shell date '+%d-%m-%Y %H:%M:%S')

# Output directory [default: ./dist]
OUTPUT_DIR = $(CURDIR)/dist

# Server Files
SERVER_FILES = $(CURDIR)/server/main/main.go

# DLL Files
DLL_FILES = $(CURDIR)/dll/main/main.go

# GO commands
GO			= go
GOGET		= $(GO) get
GOBUILD		= $(GO) build
GOTEST		= $(GO) test
GOFMT		= gofmt
_GOOS		= windows
_GOARCH		= amd64

# GO FLAGS
GO_LD_FLAGS  = -X "main.VERSION=$(VERSION)"  -X 'main.RELEASE=$(RELEASE)'

# Default Build variables
INPUT =
OUTPUT =
BUILD_MODE	= exe
OUTPUT_NAME = $(MAIN)-$(_GOOS)-$(_GOARCH)

# +++ Common +++

.PHONY: clean

all: build-server build-dll

clean:
	@-rm -r $(OUTPUT_DIR)

test:
	$(GOTEST) -v ./...

# +++ Release +++

release: GO_LD_FLAGS += -s -w
release: build-server build-dll

# +++ Server +++

build-server: BUILD_MODE = exe
build-server: INPUT = $(SERVER_FILES)
build-server: OUTPUT = $(OUTPUT_DIR)/$(OUTPUT_NAME).exe
build-server:
	$(MAKE) BUILD_MODE="$(BUILD_MODE)" GO_LD_FLAGS="$(GO_LD_FLAGS)" OUTPUT="$(OUTPUT)" INPUT="$(INPUT)" build

# +++ Dll +++

build-dll: BUILD_MODE = c-shared
build-dll: INPUT = $(DLL_FILES)
build-dll: OUTPUT = $(OUTPUT_DIR)/$(OUTPUT_NAME).dll
build-dll:
	$(MAKE) BUILD_MODE="$(BUILD_MODE)" GO_LD_FLAGS="$(GO_LD_FLAGS)" OUTPUT="$(OUTPUT)" INPUT="$(INPUT)" build

# +++ Build Target +++

build: dependencies	
	@GOOS=$(_GOOS) GOARCH=$(_GOARCH) $(GOBUILD) -buildmode=$(BUILD_MODE) -ldflags "$(GO_LD_FLAGS)" -o $(OUTPUT) $(INPUT)

# +++ Dependencies +++

dependencies:
	$(GOGET) gopkg.in/yaml.v2
	$(GOGET) golang.org/x/sys/windows
	$(GOGET) github.com/gorilla/websocket

