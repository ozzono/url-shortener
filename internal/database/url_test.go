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

	testURL := &models.URL{
		Source: fmt.Sprintf("https://%s.%s", utils.RString(5, 7), utils.RString(2, 3)),
	}

	client, err := NewClient()
	require.NoError(t, err, "NewClient")

	testURL, err = client.AddURL(testURL, false)
	require.NoError(t, err, "failed to store test url")

	url, found, err := client.FindURLBySource(testURL, false)
	require.NoError(t, err, "client.FindURLBySource")
	require.True(t, found, "source url not found")

	_, found, err = client.FindURLByShortened(testURL, false)
	require.NoError(t, err, "client.FindURLByShortened")
	require.True(t, found, "shortened url not found")

	foundURL := new(models.URL)
	require.Condition(t, func() (success bool) {
		return reflect.DeepEqual(foundURL, testURL)
	}, "found URL is not equal to test URL")

	incURL, err := client.IncrementURL(url, false)
	incu := new(models.URL)
	require.NoError(t, err, "client.IncrementURL")
	require.Condition(t, func() (success bool) {
		return url.Count+1 == incURL.Count
	}, "failed to increment url counter")

	err = client.DelURL(foundURL)
	require.NoError(t, err, "client.DelURL")

	_, found, err = client.FindURLBySource(foundURL, false)
	require.NoError(t, err, "client.FindURLBySource")
	require.Condition(t, func() (success bool) {
		return !found
	}, "found URL that should not exist")
	incu.Log("found url", true)
}
