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

	url, found, err := client.FindURLBySource(testURL.URL)
	require.NoError(t, err, "client.FindURLBySource")
	require.Condition(t, func() (success bool) {
		if !found {
			testURL.Log("source url not found")
		}
		return found
	})

	_, found, err = client.FindURLByShortened(testURL.URL)
	require.NoError(t, err, "client.FindURLByShortened")
	require.Condition(t, func() (success bool) {
		if !found {
			testURL.Log("shortened url not found")
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

	incURL, err := client.IncrementURL(url)
	incu := testingURL{t, incURL}
	require.NoError(t, err, "client.IncrementURL")
	require.Condition(t, func() (success bool) {
		success = url.Count+1 == incURL.Count
		if !success {
			u := testingURL{t, url}
			incu.Log("   incremented")
			u.Log("not incremented")
		}
		return
	}, "failed to increment url counter")

	err = client.DelURL(foundURL.URL)
	require.NoError(t, err, "client.DelURL")

	_, found, err = client.FindURLBySource(foundURL.URL)
	require.NoError(t, err, "client.FindURLBySource")
	require.Condition(t, func() (success bool) {
		return !found
	}, "found URL that should not exist")
	incu.Log("found url")
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
