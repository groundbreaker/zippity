include .env
.PHONY := default init clean test test-file cb release env env-push
.DEFAULT_GOAL = default

AWS ?= aws --profile ${AWS_PROFILE}
PROJECT ?= $(shell basename $$PWD)
GOTEST ?= AWS_PROFILE=${AWS_PROFILE} go test

default:
	@ mmake help

# init project
init: env
	@ go mod vendor

# remove build assets
clean:
	@ rm -f test.zip

# run tests
test: clean
	@ ${GOTEST} ./...

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

# run codebuild locally
cb:
	@ gb codebuild -i groundbreaker/ci-go:latest -a .artifacts -c -b buildspec.yml

# pull latest .env file
env:
	@ ${AWS} s3 cp s3://groundbreaker-eng/${PROJECT}/.env .env

# push latest .env file
env-push:
	@ ${AWS} s3 cp .env s3://groundbreaker-eng/${PROJECT}/.env --sse

%:
	@ true
