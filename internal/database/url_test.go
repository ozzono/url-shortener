package database

import (
	"reflect"
	"testing"

	"url-shortener/internal/models"
	"url-shortener/utils"

	"github.com/stretchr/testify/require"
)

type testingURL struct {
	*testing.T
	*models.URL
}

func TestURL(t *testing.T) {

	testURL := testingURL{
		t,
		&models.URL{
			FullPath:  utils.RString(5, 7),
			ShortPath: utils.RString(5, 7),
			Debug:     false,
		},
	}

	client, err := NewClient()
	require.NoError(t, err, "NewClient")

	testURL.URL, err = client.AddURL(testURL.URL)
	require.NoError(t, err, "failed to store test url")

	url, found, err := client.FindURL(testURL.URL)
	require.Condition(t, func() (success bool) {
		if !found {
			testURL.Log("test url not found")
		}
		return found
	})

	foundURL := testingURL{t, url}
	require.Condition(t, func() (success bool) {
		eq := reflect.DeepEqual(foundURL.URL, testURL.URL)
		if !eq {
			testURL.Log("test  url")
			foundURL.Log("found url")
		}
		return eq
	}, "found URL is not equal to test URL")

	require.NoError(t, err, "failed to add url")
}

func (t testingURL) Log(header string) {
	if len(header) > 0 {
		t.Logf("%s - testURL\n", header)
	}
	t.Logf("testURL.ID --------- %s", t.URL.ID.String())
	t.Logf("testURL.FullPath --- %s", t.URL.FullPath)
	t.Logf("testURL.ShortPath -- %s", t.URL.ShortPath)
	t.Logf("testURL.Count ------ %d", t.URL.Count)
}
