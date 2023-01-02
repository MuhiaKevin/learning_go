package newsfeed

import (
	"newsfeeder/platform/newsfeed"
	"testing"
)

func TestAdd(t *testing.T) {
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{})
	if len(feed.Items) != 1 {
		t.Error("Item was not added")
	}

}
func TestGetAll(t *testing.T) {

	feed := newsfeed.New()
	feed.Add(newsfeed.Item{})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Error("Item was not added")
	}
}
