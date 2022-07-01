package database

import (
	"fmt"
	"reflect"
	"testing"

	"url-shortener/internal/models"
	"url-shortener/utils"

	"github.com/stretchr/testify/require"
)

func TestURL(t *testing.T) {

	testURL := &models.TestURL{
		t,
		&models.URL{
			Source: fmt.Sprintf("https://%s.%s", utils.RString(5, 7), utils.RString(2, 3)),
		},
	}

	client, err := NewClient()
	require.NoError(t, err, "NewClient")

	testURL.URL, err = client.AddURL(testURL.URL, false)
	require.NoError(t, err, "failed to store test url")

	url, found, err := client.FindURLBySource(testURL.URL, false)
	require.NoError(t, err, "client.FindURLBySource")
	require.True(t, found, "source url not found")

	_, found, err = client.FindURLByShortened(testURL.URL, false)
	require.NoError(t, err, "client.FindURLByShortened")
	require.True(t, found, "shortened url not found")

	foundURL := &models.TestURL{t, url}
	require.Condition(t, func() (success bool) {
		eq := reflect.DeepEqual(foundURL.URL, testURL.URL)
		if !eq {
			testURL.Log("test  url")
			foundURL.Log("found url")
		}
		return eq
	}, "found URL is not equal to test URL")

	incURL, err := client.IncrementURL(url, false)
	incu := &models.TestURL{t, incURL}
	require.NoError(t, err, "client.IncrementURL")
	require.Condition(t, func() (success bool) {
		success = url.Count+1 == incURL.Count
		if !success {
			u := &models.TestURL{t, url}
			incu.Log("   incremented")
			u.Log("not incremented")
		}
		return
	}, "failed to increment url counter")

	err = client.DelURL(foundURL.URL)
	require.NoError(t, err, "client.DelURL")

	_, found, err = client.FindURLBySource(foundURL.URL, false)
	require.NoError(t, err, "client.FindURLBySource")
	require.Condition(t, func() (success bool) {
		return !found
	}, "found URL that should not exist")
	incu.Log("found url")
}
