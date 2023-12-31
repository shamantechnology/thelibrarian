package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func TestOllama(t *testing.T) {
	llm, err := ollama.New(ollama.WithModel("mistral"))
	require.NoError(t, err)

	ctx := context.Background()

	completion, err := llm.Call(
		ctx,
		"Human: Who was the first man to walk on the moon?\nAssistant:",
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	require.NoError(t, err)
	require.NotNil(t, completion)
}
