GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

GO_PACKAGE_NAME := $(shell \
	git config remote.origin.url | sed -E 's~https://(\.)?([^:]+).git~\2~' \
	)


GO_GIT_USER_ID := $(shell \
	git config remote.origin.url | sed -E 's~https://(\.)?([^:]+)/([^:]+)/([^:]+).git~\3~' \
	)

GO_GIT_REPO_ID := $(shell \
	git config remote.origin.url | sed -E 's~https://(\.)?([^:]+)/([^:]+)/([^:]+).git~\4~' \
	)

GO_BRANCH_NAME := $(shell \
	git branch --show-current \
	)

GO_PATH := $(shell \
	go env GOPATH \
	)

# Shows the variables we just set.
# By prepending `@` we prevent Make
# from printing the command before the
# stdout of the execution.
show:
	@echo "SRC  = $(GO_SRC_DIRS)"
	@echo "TEST = $(GO_TEST_DIRS)"

test: $(GO_TEST_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		go test -v ; \
		popd > /dev/null ; \
	done;

fmt: $(GO_SRC_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		echo $$dir ; \
		echo $(MAKE)  ; \
		echo $(LD)  ; \
		echo $(CC)  ; \
		echo $(@)  ; \
		echo $(%)  ; \
		echo $(?)  ; \
		popd > /dev/null ; \
	done;


build: 
	go build -o bin/main main.go

run: 
	go run main.go


compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

test: 
	go test

all:  build

clean:
	rm go.mod || true
	rm -rf servermock || true
	rm -rf openApi || true
	rm -rf servermock || true


init: 
	echo "initialization of the local go module with pakage name"
	mkdir openApi
	wget https://raw.github.com/marcocipri/RESTful-interfaces/$(GO_BRANCH_NAME)/examples/helloworld.yml -O openApi/helloworld.yml
	echo "mock code generation"
	docker run --rm \
  	-v ${PWD}:/local openapitools/openapi-generator-cli generate \
  	-i /local/openApi/helloworld.yml \
  	-g go-server \
  	-o /local/servermock --git-user-id ${GO_GIT_USER_ID} --git-repo-id ${GO_GIT_REPO_ID}
	${GO_PATH}/bin/goimports -w ./servermock/
	echo "initialization of the local go module with pakage name"
	go mod init $(GO_PACKAGE_NAME) 