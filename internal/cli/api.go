package cli

import (
	"context"
	"errors"
	"fmt"
	"github.com/tidwall/resp"
	"time"
)

func (c *CliClient) TrySendCommandWithRetries(command string, args ...string) (string, error) {
	retries := 3
	waitInterval := time.Second
	increaseInterval := func(t time.Duration) time.Duration {
		return t * 2
	}

	var (
		err      error
		response string
	)

	ctx := context.Background()
	response, err = c.trySendCommand(ctx, command, args...)
	if err == nil {
		return response, nil
	}

	//returns all internal errors not related to timeout
	if !errors.Is(err, context.DeadlineExceeded) {
		return "", err
	}

	for range retries {
		fmt.Println("retrying")
		response, err = c.trySendCommand(ctx, command, args...)
		if err == nil {
			return response, nil
		}
		time.Sleep(waitInterval)
		waitInterval = increaseInterval(waitInterval)
	}
	//returns deadline exceeded
	return "", err
}

func (c *CliClient) trySendCommand(ctx context.Context, command string, args ...string) (string, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Second)
	defer cancel()

	resultCh := make(chan struct {
		response string
		err      error
	}, 1)

	go func() {
		res, err := c.sendCommand(command, args...)
		resultCh <- struct {
			response string
			err      error
		}{res, err}
		close(resultCh)
	}()

	select {
	case res := <-resultCh:
		if res.err != nil {
			return "", res.err
		}
		return res.response, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (c *CliClient) sendCommand(command string, args ...string) (string, error) {
	values := []resp.Value{resp.StringValue(command)}
	for _, arg := range args {
		values = append(values, resp.StringValue(arg))
	}

	var stringRequest string
	for _, val := range values {
		stringRequest += val.String()
	}

	req, err := resp.ArrayValue(values).MarshalRESP()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to Marshal RESP request: "), err)
		return "", err
	}

	_, err = c.conn.Write(req)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to send data to server: "), err)
		return "", err
	}

	//TODO: обернуть в отдельную функцию и распарсить RESP
	var n int
	buf := make([]byte, 4096)
	n, err = c.conn.Read(buf)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read response from server: "), err)
		return "", err
	}

	return string(buf[:n]), nil
}
