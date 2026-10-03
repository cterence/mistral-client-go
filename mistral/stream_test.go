package mistral

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) (*APIClient, *httptest.Server) {
	t.Helper()
	ts := httptest.NewServer(handler)
	t.Cleanup(ts.Close)
	cfg := NewConfiguration()
	cfg.Scheme = "http"
	cfg.Host = ts.Listener.Addr().String()
	return NewAPIClient(cfg), ts
}

func TestChatCompletionStream(t *testing.T) {
	var gotBody ChatCompletionRequest
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" {
			t.Errorf("path = %s, want /v1/chat/completions", r.URL.Path)
		}
		if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
			t.Errorf("decode request: %v", err)
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.(http.Flusher).Flush()
		io.WriteString(w, "data: {\"id\":\"a1\",\"model\":\"m\",\"choices\":[]}\n\n")
		io.WriteString(w, "data: {\"data\":{\"id\":\"a2\",\"model\":\"m\",\"choices\":[]}}\n\n")
		io.WriteString(w, sseDone)
	})

	req := ChatCompletionRequest{Model: "m"}
	var ids []string
	for chunk, err := range client.ChatAPI.ChatCompletionV1ChatCompletionsPost(context.Background()).ChatCompletionRequest(req).ExecuteStream() {
		if err != nil {
			t.Fatalf("stream error: %v", err)
		}
		ids = append(ids, chunk.Id)
	}

	if gotBody.Stream == nil || !*gotBody.Stream {
		t.Errorf("request stream = %v, want true", gotBody.Stream)
	}
	want := []string{"a1", "a2"} // raw chunk and {data: chunk} envelope both accepted
	if len(ids) != len(want) {
		t.Fatalf("got %d chunks (%v), want %d", len(ids), ids, len(want))
	}
	for i := range want {
		if ids[i] != want[i] {
			t.Errorf("chunk %d id = %s, want %s", i, ids[i], want[i])
		}
	}
}

func TestChatCompletionStreamError(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message":"bad key"}`)
	})

	req := ChatCompletionRequest{Model: "m"}
	for _, err := range client.ChatAPI.ChatCompletionV1ChatCompletionsPost(context.Background()).ChatCompletionRequest(req).ExecuteStream() {
		if err == nil {
			t.Fatal("expected error, got chunk")
		}
		return
	}
	t.Fatal("expected one error event")
}

// TestChatCompletionStreamLive runs against the real API; skipped
// unless MISTRAL_API_KEY is set:
//
//	MISTRAL_API_KEY=... nix develop -c go test -run Live -v .
func TestChatCompletionStreamLive(t *testing.T) {
	apiKey := os.Getenv("MISTRAL_API_KEY")
	if apiKey == "" {
		t.Skip("MISTRAL_API_KEY not set")
	}
	cfg := NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	client := NewAPIClient(cfg)

	req := ChatCompletionRequest{
		Model: "mistral-small-latest",
		Messages: []MessagesInner{UserMessageAsMessagesInner(&UserMessage{
			Role:    PtrString("user"),
			Content: *NewNullableContent3(&Content3{String: PtrString("Count from 1 to 5.")}),
		})},
	}

	var text strings.Builder
	for chunk, err := range client.ChatAPI.ChatCompletionV1ChatCompletionsPost(context.Background()).ChatCompletionRequest(req).ExecuteStream() {
		if err != nil {
			t.Fatalf("stream error: %v", err)
		}
		for _, choice := range chunk.Choices {
			if c := choice.Delta.Content.Get().String; c != nil {
				text.WriteString(*c)
			}
		}
	}
	t.Logf("streamed: %q", text.String())
	if text.Len() == 0 {
		t.Fatal("no content streamed")
	}
}
