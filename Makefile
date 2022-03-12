test: lint unit-test

PLATFORM=local
NAME=myApp 

all: bin/${NAME}

.DEFAULT_GOAL := bin/${NAME}

.PHONY: name
name:
	echo "Printing app name : '${NAME}'"

.PHONY: bin/${NAME}
bin/${NAME}:
	@docker build . --target bin \
	--output bin/ \
	--platform ${PLATFORM} \
	--build-arg go_app_name=${NAME}

.PHONY: unit-test
unit-test:
	echo "unit tests"
	@docker build . --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	echo "Coverage of unit tests"
	@docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	echo "Linting"
	@docker build . --target lint --build-arg go_app_name=${NAME}
