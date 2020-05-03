package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	t.Run("unsupported file", func(t *testing.T) {
		from := path.Join("/", "dev", "zero")
		to := "output.txt"
		var offset int64 = 0
		var limit int64 = 0
		err := Copy(from, to, offset, limit)
		assert.Error(t, err)
		assert.Equal(t, err, ErrUnsupportedFile)
	})

	t.Run("offset exids file size", func(t *testing.T) {
		from := path.Join("testdata", "input.txt")
		to := "output.txt"
		var offset int64 = 10000
		var limit int64 = 0
		err := Copy(from, to, offset, limit)
		assert.Error(t, err)
		assert.Equal(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("offset 0 limit 0", func(t *testing.T) {
		from := path.Join("testdata", "input.txt")
		to := "output.txt"
		var offset int64 = 0
		var limit int64 = 0
		err := Copy(from, to, offset, limit)
		assert.Nil(t, err)
		resultFileBytes, err := ioutil.ReadFile(to)
		assert.Nil(t, err)
		corrctFileBytes, err := ioutil.ReadFile(path.Join("testdata", "out_offset0_limit0.txt"))
		assert.Nil(t, err)
		assert.Equal(t, resultFileBytes, corrctFileBytes)
		os.Remove(to)
	})

	t.Run("offset 0 limit 10", func(t *testing.T) {
		from := path.Join("testdata", "input.txt")
		to := "output.txt"
		var offset int64 = 0
		var limit int64 = 10
		err := Copy(from, to, offset, limit)
		assert.Nil(t, err)
		resultFileBytes, err := ioutil.ReadFile(to)
		assert.Nil(t, err)
		corrctFileBytes, err := ioutil.ReadFile(path.Join("testdata", "out_offset0_limit10.txt"))
		assert.Nil(t, err)
		assert.Equal(t, resultFileBytes, corrctFileBytes)
		os.Remove(to)
	})

	t.Run("offset 0 limit 1000", func(t *testing.T) {
		from := path.Join("testdata", "input.txt")
		to := "output.txt"
		var offset int64 = 0
		var limit int64 = 1000
		err := Copy(from, to, offset, limit)
		assert.Nil(t, err)
		resultFileBytes, err := ioutil.ReadFile(to)
		assert.Nil(t, err)
		corrctFileBytes, err := ioutil.ReadFile(path.Join("testdata", "out_offset0_limit1000.txt"))
		assert.Nil(t, err)
		assert.Equal(t, resultFileBytes, corrctFileBytes)
		os.Remove(to)
	})

	t.Run("offset 0 limit 10000", func(t *testing.T) {
		from := path.Join("testdata", "input.txt")
		to := "output.txt"
		var offset int64 = 0
		var limit int64 = 10000
		err := Copy(from, to, offset, limit)
		assert.Nil(t, err)
		resultFileBytes, err := ioutil.ReadFile(to)
		assert.Nil(t, err)
		corrctFileBytes, err := ioutil.ReadFile(path.Join("testdata", "out_offset0_limit10000.txt"))
		assert.Nil(t, err)
		assert.Equal(t, resultFileBytes, corrctFileBytes)
		os.Remove(to)
	})

	t.Run("offset 100 limit 1000", func(t *testing.T) {
		from := path.Join("testdata", "input.txt")
		to := "output.txt"
		var offset int64 = 100
		var limit int64 = 1000
		err := Copy(from, to, offset, limit)
		assert.Nil(t, err)
		resultFileBytes, err := ioutil.ReadFile(to)
		assert.Nil(t, err)
		corrctFileBytes, err := ioutil.ReadFile(path.Join("testdata", "out_offset100_limit1000.txt"))
		assert.Nil(t, err)
		assert.Equal(t, resultFileBytes, corrctFileBytes)
		os.Remove(to)
	})

	t.Run("offset 6000 limit 1000", func(t *testing.T) {
		from := path.Join("testdata", "input.txt")
		to := "output.txt"
		var offset int64 = 6000
		var limit int64 = 1000
		err := Copy(from, to, offset, limit)
		assert.Nil(t, err)
		resultFileBytes, err := ioutil.ReadFile(to)
		assert.Nil(t, err)
		corrctFileBytes, err := ioutil.ReadFile(path.Join("testdata", "out_offset6000_limit1000.txt"))
		assert.Nil(t, err)
		assert.Equal(t, resultFileBytes, corrctFileBytes)
		os.Remove(to)
	})
}
