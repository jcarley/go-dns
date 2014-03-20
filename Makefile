NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS=$(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
GOFILES=$(go list -f '{{range .GoFiles}}{{.}} {{end}}' ./...)
TESTGOFILES=$(go list -f '{{range .TestGoFiles}}{{.}} {{end}}' ./...)
GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=2

all: deps
	@mkdir -p bin/
	@echo "$(OK_COLOR)==> Building$(NO_COLOR)"
	@bash --norc -i ./scripts/devcompile.sh

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d

updatedeps:
	@echo "$(OK_COLOR)==> Updating all dependencies$(NO_COLOR)"
	@go get -d -v -u ./...
	@echo $(DEPS) | xargs -n1 go get -d -u

clean:
	@rm -rf bin/ local/ pkg/ src/
	@rm go-dns go-dns-linuxarm go-dns-linuxamd64 go-dns.exe

format:
	# @echo $(GOFILES) | xargs -n1 gofmt -l -w -tabs=false -tabwidth=2 $(GOFILES)
	@echo "$(OK_COLOR)==> Formatting source files$(NO_COLOR)"
	@echo $(GOFMT) -w ./...

test: deps
	@echo "$(OK_COLOR)==> Testing go-dns...$(NO_COLOR)"
	go test ./...

build:
	@echo "$(OK_COLOR)==> Building go-dns...$(NO_COLOR)"
	@XC_ARCH=${XC_ARCH:-"386 amd64 arm"}
	@XC_OS=${XC_OS:-linux darwin windows}
	@gox -os="${XC_OS}" -arch="${XC_ARCH}" -output "pkg/{{.OS}}_{{.Arch}}/go-dns-{{.Dir}}" ./...
	# @gox -os= "pkg/{{.OS}}_{{.Arch}}/go-dns-{{.Dir}}"
	# @GOOS=windows GOARCH=amd64 go build -o "pkg/{{.OS}}_{{.Arch}}/go-dns-{{.Dir}}"
	# @GOOS=linux GOARCH=amd64 go build -o "pkg/{{.OS}}_{{.Arch}}/go-dns-{{.Dir}}"

.PHONY: all clean deps format test updatedeps
