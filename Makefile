.PHONY := default init clean test test-file release
.DEFAULT_GOAL = default

PROJECT ?= $(shell basename $$PWD)
GOTEST ?= go test

default:
	@ mmake help

# init project
init:
	@ go mod vendor

# remove build assets
clean:
	@ rm -f test.zip

# run tests
test: clean
	@ ${GOTEST} ./...

# run specific test
#
#     make test-file ZippityTestSuite/TestVersion
test-file: clean
	@ ${GOTEST} -run $(filter-out $@, $(MAKECMDGOALS))

# create release VERSION on github
#
# VERSION should being with a v and be in SemVer format.
release: test
	$(eval VERSION=$(filter-out $@, $(MAKECMDGOALS)))
	$(if ${VERSION},@true,$(error "VERSION is required"))
	git commit --allow-empty -am ${VERSION}
	git push
	hub release create -m ${VERSION} -e ${VERSION}

%:
	@ true
