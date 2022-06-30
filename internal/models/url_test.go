package models

import (
	"testing"
	"url-shortener/utils"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {

	url := &URL{
		FullPath:  utils.RString(5, 7),
		ShortPath: utils.RString(5, 7),
	}

	_, err := url.Add()
	require.NoError(t, err, "failed to add url")
}
