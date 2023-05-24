package wrappers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	clientTimeout = 10 * time.Second
)

func Do[Req, Res any](ctx context.Context, req *Req, method, path string) (*Res, error) {
	rawData, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("encode request: %w", err)
	}

	ctx, fnCancel := context.WithTimeout(ctx, clientTimeout)
	defer fnCancel()

	httpRequest, err := http.NewRequestWithContext(ctx, method, path, bytes.NewBuffer(rawData))
	if err != nil {
		return nil, fmt.Errorf("prepare request: %w", err)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", httpResponse.StatusCode)
	}

	var res Res
	if err = json.NewDecoder(httpResponse.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &res, nil
}
