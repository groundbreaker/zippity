include .env
.PHONY := default init test test-file build pkg deploy triggers validate stack-events cb release env env-push
.DEFAULT_GOAL = default

AWS ?= aws --profile ${AWS_PROFILE}
PROJECT ?= $(shell basename $$PWD)
GOTEST ?= AWS_PROFILE=${AWS_PROFILE} go test
ARN_PREFIX_LAMBDA := arn:aws:lambda:us-east-1:140981404942:function:
FUNC := FUNC_NAME
STACK := STACK_NAME

default:
	@ mmake help

# init project
init: env
	@ go mod vendor

# run tests
test:
	@ SENTRY_DSN=${SENTRY_DSN} ${GOTEST} ./...

test-file:
	@ SENTRY_DSN=${SENTRY_DSN} ${GOTEST} -run $(filter-out $@, $(MAKECMDGOALS))

# remove build assets
clean:
	@ rm -rf bin/${FUNC}

# build assets
build: clean
	@ GOOS=linux GOARCH=amd64 go build -o bin/${FUNC}

# package stack
pkg: build
	@ ${AWS} cloudformation package \
	    --template-file src.yml \
	    --output-template-file dist.yml \
	    --s3-bucket gb-cloud

# deploy stack
deploy: validate
	@ ${AWS} cloudformation deploy \
	    --template-file dist.yml \
	    --s3-bucket gb-cloud \
	    --stack-name ${STACK}

# validate cloudformation templates
validate: pkg
	@ ${AWS} cloudformation validate-template --template-body file://src.yml
	@ ${AWS} cloudformation validate-template --template-body file://dist.yml

# list function triggers
triggers:
	@ ${AWS} lambda list-event-source-mappings \
	    --function-name ${ARN_PREFIX_LAMBDA}${STACK}:latest

# create release VERSION on github
#
# VERSION should being with a v and be in SemVer format.
release:
	$(eval VERSION=$(filter-out $@, $(MAKECMDGOALS)))
	$(if ${VERSION},@true,$(error "VERSION is required"))
	git commit --allow-empty -am ${VERSION}
	git push
	hub release create -m ${VERSION} -e ${VERSION}

# show stack-events
stack-events:
	@ ${AWS} cloudformation describe-stack-events \
		--stack-name ${STACK}

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
