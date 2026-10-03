USER_ID  := cterence
REPO_ID  := mistral-client-golang
GEN_GO   := $(filter-out fixup_types.go,$(wildcard *.go))

# openapi-generator emits a few invalid literals and skips a few inline
# types for this spec; fixup patches/generates those post-gen.
.PHONY: generate fixup build test

generate:
	rm -rf api docs test $(GEN_GO) go.mod go.sum .openapi-generator README.md git_push.sh .openapi-generator-ignore .travis.yml
	nix develop -c openapi-generator-cli generate -g go -i openapi.yaml -o . \
		--git-user-id $(USER_ID) --git-repo-id $(REPO_ID) \
		--additional-properties=packageName=mistral,packageVersion=0.1.0,enumClassPrefix=true
	rm -f git_push.sh
	$(MAKE) fixup
	nix develop -c sh -c 'go mod tidy && go build ./...'

fixup:
	nix develop -c sed -i -z \
		-e 's/\t\tvar defaultValue ConnectorsQueryFilters = {[^}]*}\n\t\tparameterAddToHeaderOrQuery(localVarQueryParams, "query_filters", defaultValue, "form", "")\n\t\tr.queryFilters = \&defaultValue\n//g' \
		-e 's/\tvar toolChoice ToolChoice = auto\n\tthis.ToolChoice = \&toolChoice/\tvar toolChoice ToolChoice\n\tauto := TOOLCHOICEENUM_AUTO\n\ttoolChoice.ToolChoiceEnum = \&auto\n\tthis.ToolChoice = \&toolChoice/g' \
		-e 's/\tvar prediction Prediction = {type=content, content=}\n\tthis.Prediction = \&prediction\n//g' \
		api_beta_connectors.go model_chat_completion_request.go model_agents_completion_request.go
	nix develop -c sed -i 's/\*OsFile/OsFile/g' model_input_audio.go

build:
	nix develop -c go build ./...

test:
	nix develop -c go test ./...
