# Version and Relase 
VERSION	= 0.0.1
RELEASE = $(shell date '+%d-%m-%Y %H:%M:%S')

# define recursive wildcards
rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))
# GO Files
FILES = $(call rwildcard, ,*.go)
# Main files
MAIN_FILES = $(wildcard *.go)

# Executabels
EXECUTABLE = nmsB

# GO commands
GO			= go
GOGET		= $(GO) get
GOBUILD		= $(GO) build
GOTEST		= $(GO) test
GOFMT		= gofmt
_GOOS		= linux
_GOARCH		= amd64

# GO FLAGS
GO_LD_FLAGS  = -X "main.VERSION=$(VERSION)"  -X 'main.RELEASE=$(RELEASE)'

# $(CURDIR)
BIN_DIR = $(CURDIR)/dist

# Default Build variables
EXE_ENDOING =

.PHONY: clean

all: build-windows

release: GO_LD_FLAGS += -s -w
release: build-windows

test:
	$(GOTEST) -v ./...

fmt:
	$(GOFMT) -l -e -w $(FILES)

build-windows: _GOOS=windows
build-windows: EXE_ENDOING=.exe
build-windows: 	
	$(MAKE) GO_LD_FLAGS="$(GO_LD_FLAGS)" EXE_ENDOING=$(EXE_ENDOING) _GOOS=$(_GOOS) _GOARCH=$(_GOARCH) build

build: OUTOUT=$(EXECUTABLE)-$(_GOOS)-$(_GOARCH)$(EXE_ENDOING)
build: dependencies	
	@GOOS=$(_GOOS) GOARCH=$(_GOARCH) $(GOBUILD) -buildmode=exe -ldflags "$(GO_LD_FLAGS)" -o $(BIN_DIR)/$(OUTOUT) $(MAIN_FILES)

dependencies:
	$(GOGET) gopkg.in/yaml.v2
	$(GOGET) golang.org/x/sys/windows

clean:
	@-rm -r $(BIN_DIR)