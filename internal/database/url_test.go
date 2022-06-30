package database

import (
	"fmt"
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
			Source:    fmt.Sprintf("https://%s.%s", utils.RString(5, 7), utils.RString(2, 3)),
			Shortened: utils.RString(5, 7),
			Debug:     false,
		},
	}

	client, err := NewClient()
	require.NoError(t, err, "NewClient")

	testURL.URL, err = client.AddURL(testURL.URL)
	require.NoError(t, err, "failed to store test url")

	url, found, err := client.FindURL(testURL.URL)
	require.NoError(t, err, "client.FindURL")
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

	err = client.DelURL(foundURL.URL)
	require.NoError(t, err, "client.DelURL")

	_, _, err = client.FindURL(foundURL.URL)
	require.NoError(t, err, "client.FindURL")
	require.Condition(t, func() (success bool) {
		return found
	}, "found URL that should not exist")
	foundURL.Log("found url")
}

func (t testingURL) Log(header string) {
	if len(header) > 0 {
		t.Logf("%s - testURL\n", header)
	}
	t.Logf("testURL.ID --------- %s", t.URL.ID.String())
	t.Logf("testURL.Source ----- %s", t.URL.Source)
	t.Logf("testURL.Shortened -- %s", t.URL.Shortened)
	t.Logf("testURL.Count ------ %d", t.URL.Count)
}
