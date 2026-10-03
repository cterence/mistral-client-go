USER_ID  := cterence
REPO_ID  := mistral-client-go
SDK_DIR  := mistral
GEN_DIR  := .gen
HAND_GO  := $(addprefix $(SDK_DIR)/,fixup_types.go stream.go stream_test.go)
GEN_GO   := $(filter-out $(HAND_GO),$(wildcard $(SDK_DIR)/*.go))
MOD_PATH := github.com/$(USER_ID)/$(REPO_ID)

# openapi-generator emits a few invalid literals and skips a few inline
# types for this spec; fixup patches/generates those post-gen.
.PHONY: generate fixup build test

generate:
	rm -rf $(GEN_DIR) $(SDK_DIR)/api $(SDK_DIR)/docs $(SDK_DIR)/test $(GEN_GO)
	nix develop -c openapi-generator-cli generate -g go -i openapi.yaml -o $(GEN_DIR) \
		--git-user-id $(USER_ID) --git-repo-id $(REPO_ID) \
		--additional-properties=packageName=mistral,packageVersion=0.1.0,enumClassPrefix=true
	rm -f $(GEN_DIR)/git_push.sh $(GEN_DIR)/.travis.yml
	mv $(GEN_DIR)/go.mod go.mod
	mv $(GEN_DIR)/* $(SDK_DIR)/
	# tests import the module root; the package lives in $(SDK_DIR)
	nix develop -c sed -i 's|"$(MOD_PATH)"|"$(MOD_PATH)/$(SDK_DIR)"|g' $(SDK_DIR)/test/*.go $(SDK_DIR)/README.md
	$(MAKE) fixup
	nix develop -c sh -c 'go mod tidy && go build ./...'

fixup:
	nix develop -c sed -i -z \
		-e 's/\t\tvar defaultValue ConnectorsQueryFilters = {[^}]*}\n\t\tparameterAddToHeaderOrQuery(localVarQueryParams, "query_filters", defaultValue, "form", "")\n\t\tr.queryFilters = \&defaultValue\n//g' \
		-e 's/\tvar toolChoice ToolChoice = auto\n\tthis.ToolChoice = \&toolChoice/\tvar toolChoice ToolChoice\n\tauto := TOOLCHOICEENUM_AUTO\n\ttoolChoice.ToolChoiceEnum = \&auto\n\tthis.ToolChoice = \&toolChoice/g' \
		-e 's/\tvar prediction Prediction = {type=content, content=}\n\tthis.Prediction = \&prediction\n//g' \
		$(SDK_DIR)/api_beta_connectors.go $(SDK_DIR)/model_chat_completion_request.go $(SDK_DIR)/model_agents_completion_request.go
	nix develop -c sed -i 's/\*OsFile/OsFile/g' $(SDK_DIR)/model_input_audio.go
	# live API sends undocumented fields; strict decoding breaks on spec lag
	nix develop -c sed -i '/decoder.DisallowUnknownFields()/d' $(SDK_DIR)/model_*.go

build:
	nix develop -c go build ./...

test:
	nix develop -c go test ./...
