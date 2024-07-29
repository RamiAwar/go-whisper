package whisper_test

import (
	"context"
	"testing"

	// Packages
	whisper "github.com/mutablelogic/go-whisper/pkg/whisper"
	task "github.com/mutablelogic/go-whisper/pkg/whisper/task"
	assert "github.com/stretchr/testify/assert"

	// Namespace imports
	. "github.com/djthorpe/go-errors"
)

const MODEL_TINY = "ggml-tiny.en-q5_1.bin"

func Test_whisper_001(t *testing.T) {
	assert := assert.New(t)
	service, err := whisper.New(t.TempDir())
	if !assert.Nil(err) {
		t.SkipNow()
	}
	assert.NotNil(service)
	assert.NoError(service.Close())
}

func Test_whisper_002(t *testing.T) {
	assert := assert.New(t)
	service, err := whisper.New(t.TempDir())
	if !assert.Nil(err) {
		t.SkipNow()
	}
	defer service.Close()

	t.Run("NotFound", func(t *testing.T) {
		// Download a model - not found
		_, err = service.DownloadModel(context.Background(), "notfound.bin", nil)
		assert.ErrorIs(err, ErrNotFound)
	})

	t.Run("Download", func(t *testing.T) {
		// Download a model
		model, err := service.DownloadModel(context.Background(), MODEL_TINY, nil)
		assert.NoError(err)
		t.Log(model)
	})

	t.Run("Exists", func(t *testing.T) {
		// Get the model
		model := service.GetModelById(MODEL_TINY)
		assert.NotNil(model)
	})

	t.Run("Delete", func(t *testing.T) {
		// Delete the model
		err := service.DeleteModelById(MODEL_TINY)
		assert.NoError(err)
	})

	t.Run("NotExists", func(t *testing.T) {
		// Get the model
		model := service.GetModelById(MODEL_TINY)
		assert.Nil(model)
	})
}

func Test_whisper_003(t *testing.T) {
	assert := assert.New(t)
	service, err := whisper.New(t.TempDir())
	if !assert.Nil(err) {
		t.SkipNow()
	}
	defer service.Close()

	t.Run("Download", func(t *testing.T) {
		// Download a model
		model, err := service.DownloadModel(context.Background(), MODEL_TINY, nil)
		assert.NoError(err)
		t.Log(model)
	})

	t.Run("WithModelContext1", func(t *testing.T) {
		model := service.GetModelById(MODEL_TINY)
		assert.NotNil(model)

		// Get the model for the first time
		assert.NoError(service.WithModel(model, func(ctx *task.Context) error {
			assert.NotNil(ctx)
			return nil
		}))
	})

	t.Run("WithModelContext2", func(t *testing.T) {
		model := service.GetModelById(MODEL_TINY)
		assert.NotNil(model)

		// Get the model for the first time
		assert.NoError(service.WithModel(model, func(ctx *task.Context) error {
			assert.NotNil(ctx)
			return nil
		}))
	})

	t.Run("WithModelContext3", func(t *testing.T) {
		model := service.GetModelById(MODEL_TINY)
		assert.NotNil(model)

		// Get the model for the first time
		assert.NoError(service.WithModel(model, func(ctx *task.Context) error {
			assert.NotNil(ctx)
			return nil
		}))
	})
}

func Test_whisper_004(t *testing.T) {
	assert := assert.New(t)
	service, err := whisper.New(t.TempDir(), whisper.OptMaxConcurrent(2))
	if !assert.Nil(err) {
		t.SkipNow()
	}
	defer service.Close()

	t.Run("Download", func(t *testing.T) {
		// Download a model
		model, err := service.DownloadModel(context.Background(), MODEL_TINY, nil)
		assert.NoError(err)
		t.Log(model)
	})

	t.Run("WithModelContext1", func(t *testing.T) {
		t.Parallel()
		model := service.GetModelById(MODEL_TINY)
		assert.NotNil(model)

		// Get the model for the first time
		assert.NoError(service.WithModel(model, func(ctx *task.Context) error {
			assert.NotNil(ctx)
			return nil
		}))
	})

	t.Run("WithModelContext2", func(t *testing.T) {
		t.Parallel()
		model := service.GetModelById(MODEL_TINY)
		assert.NotNil(model)

		// Get the model for the first time
		assert.NoError(service.WithModel(model, func(ctx *task.Context) error {
			assert.NotNil(ctx)
			return nil
		}))
	})

	t.Run("WithModelContext3", func(t *testing.T) {
		t.Parallel()
		model := service.GetModelById(MODEL_TINY)
		assert.NotNil(model)

		// Get the model for the first time
		err := service.WithModel(model, func(ctx *task.Context) error {
			assert.NotNil(ctx)
			return nil
		})
		assert.ErrorIs(err, ErrChannelBlocked)
	})
}
