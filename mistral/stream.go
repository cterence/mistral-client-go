// Hand-written streaming support for SSE endpoints (stream=true).
// Referenced types are generated; do not delete.

package mistral

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"net/http"
	"net/url"
	"strings"
)

// sseDone is the sentinel line the API sends to close an SSE stream.
const sseDone = "[" + "DONE" + "]"

// ExecuteStream streams chat completions over SSE.
// Forces stream=true and yields CompletionChunk values; iteration ends at
// data error (yielded with a nil chunk).
func (r ApiChatCompletionV1ChatCompletionsPostRequest) ExecuteStream() iter.Seq2[*CompletionChunk, error] {
	if r.chatCompletionRequest == nil {
		return errSeq[*CompletionChunk](fmt.Errorf("chatCompletionRequest is required"))
	}
	r.chatCompletionRequest.Stream = PtrBool(true)
	return streamCompletions(r.ApiService.client, r.ctx, "/v1/chat/completions", r.chatCompletionRequest)
}

// ExecuteStream streams FIM completions over SSE, same contract as chat.
func (r ApiFimCompletionV1FimCompletionsPostRequest) ExecuteStream() iter.Seq2[*CompletionChunk, error] {
	if r.fIMCompletionRequest == nil {
		return errSeq[*CompletionChunk](fmt.Errorf("fIMCompletionRequest is required"))
	}
	r.fIMCompletionRequest.Stream = PtrBool(true)
	return streamCompletions(r.ApiService.client, r.ctx, "/v1/fim/completions", r.fIMCompletionRequest)
}

func errSeq[T any](err error) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) { yield(*new(T), err) }
}

func streamCompletions(c *APIClient, ctx context.Context, path string, body any) iter.Seq2[*CompletionChunk, error] {
	return func(yield func(*CompletionChunk, error) bool) {
		if ctx == nil {
			ctx = context.Background()
		}
		basePath, err := c.cfg.ServerURLWithContext(ctx, "streamCompletions")
		if err != nil {
			yield(nil, err)
			return
		}
		req, err := c.prepareRequest(ctx, basePath+path, http.MethodPost, body,
			map[string]string{"Accept": "text/event-stream"},
			url.Values{}, url.Values{}, nil)
		if err != nil {
			yield(nil, err)
			return
		}
		resp, err := c.callAPI(req)
		if err != nil {
			yield(nil, err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode >= http.StatusMultipleChoices {
			b, _ := io.ReadAll(resp.Body)
			yield(nil, fmt.Errorf("stream %s: HTTP %d: %s", path, resp.StatusCode, b))
			return
		}

		sc := bufio.NewScanner(resp.Body)
		sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
		for sc.Scan() {
			data, ok := strings.CutPrefix(sc.Text(), "data:")
			if !ok {
				continue // comment, event field, or blank line
			}
			data = strings.TrimSpace(data)
			if data == "" {
				continue
			}
			if data == sseDone {
				return
			}
			var chunk CompletionChunk
			if err := json.Unmarshal([]byte(data), &chunk); err != nil || chunk.Id == "" {
				// ponytail: spec models payloads as {data: chunk} (CompletionEvent),
				// live API sends raw chunks. Remove if the API settles on one format.
				var env struct {
					Data CompletionChunk `json:"data"`
				}
				if envErr := json.Unmarshal([]byte(data), &env); envErr != nil || env.Data.Id == "" {
					if err == nil {
						err = fmt.Errorf("chunk has no id: %s", data)
					}
					yield(nil, fmt.Errorf("decode chunk in %s: %w", path, err))
					return
				}
				chunk = env.Data
			}
			if !yield(&chunk, nil) {
				return
			}
		}
		if err := sc.Err(); err != nil {
			yield(nil, err)
		}
	}
}
