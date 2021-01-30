GO_TEST_FLAGS:=-timeout 5s
ifeq ($(COVERAGE),1)
GO_TEST_FLAGS:=$(GO_TEST_FLAGS) -coverprofile=coverage.out
endif

.PHONY: all
all:

.PHONY: check
check: 
	go test $(GO_TEST_FLAGS) ./... $(GO_TEST_PIPE)
ifeq ($(COVERAGE),1)
	go tool cover -html=coverage.out
endif
