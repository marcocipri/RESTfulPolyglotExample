.PHONY:  clean run init

PY_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

PY_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

PY_PACKAGE_NAME := $(shell \
	git config remote.origin.url | sed -E 's~https://(\.)?([^:]+).git~\2~' \
	)


PY_GIT_USER_ID := $(shell \
	git config remote.origin.url | sed -E 's~https://(\.)?([^:]+)/([^:]+)/([^:]+).git~\3~' \
	)

PY_GIT_REPO_ID := $(shell \
	git config remote.origin.url | sed -E 's~https://(\.)?([^:]+)/([^:]+)/([^:]+).git~\4~' \
	)

PY_BRANCH_NAME := $(shell \
	git rev-parse --abbrev-ref HEAD \
	)



# Shows the variables we just set.
# By prepending `@` we prevent Make
# from printing the command before the
# stdout of the execution.
show:
	@echo "SRC  = $(PY_SRC_DIRS)"
	@echo "TEST = $(PY_TEST_DIRS)"

test: $(PY_TEST_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		go test -v ; \
		popd > /dev/null ; \
	done;



run: 
	cd myclient/client/; python my.py


clean:
	rm -rf servermock || true
	rm -rf openApi || true
	rm -rf generated || true


init: 
	echo "initialization of the local go module with pakage name"
	mkdir openApi
	wget https://raw.github.com/marcocipri/RESTful-interfaces/$(PY_BRANCH_NAME)/examples/pets.yml -O openApi/pets.yml
	echo "mock code generation"
	docker run --rm \
  	-v ${PWD}:/local openapitools/openapi-generator-cli generate \
  	-i /local/openApi/pets.yml \
  	-g python \
	--additional-properties generateInterfaces=true \
	--additional-properties sourceFolder=clientmock \
	--additional-properties packageName=RESTfulPolyglotEamplePythonClient \
  	-o /local/generated --git-user-id ${PY_GIT_USER_ID} --git-repo-id ${PY_GIT_REPO_ID} 
	cd generated ; python setup.py install --user
	rm -rf generated
	
