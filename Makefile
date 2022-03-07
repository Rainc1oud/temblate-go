# Makefile easier to remember

.PHONY: test
test: test/temblate_test.go test/messages.go
	go test -v ./$(<D)

test/messages.go: cmd/gen/gen.go $(wildcard testtemplates/*.gotmpl) test/
	go run $< $(word 2,$(^D)) $@

%/:
	mkdir -p $@

.PHONY: clean
clean:
	rm -f test/messages.go