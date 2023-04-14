package webflow

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	testClient = NewClient(
		testToken,
	)
}

var (
	testClient *Client
)

const (
	testToken        = ""
	testEmail        = ""
	testSiteId       = ""
	testDomain       = ""
	testCollectionId = ""
	testItemId       = ""
)

func TestAuthorizedUser(t *testing.T) {
	user, err := testClient.AuthorizedUser()
	spew.Dump(user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testEmail, user.Email)
}

func TestAuthorizedInfo(t *testing.T) {
	info, err := testClient.AuthorizedInfo()
	spew.Dump(info)
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

func TestListSites(t *testing.T) {
	sites, err := testClient.ListSites()
	spew.Dump(sites)
	assert.NoError(t, err)
	assert.Greater(t, len(sites), 0)
}

func TestGetSite(t *testing.T) {
	site, err := testClient.GetSite(testSiteId)
	assert.NoError(t, err)
	assert.NotNil(t, site)
	assert.Equal(t, testSiteId, site.Id)
}

func TestPublishSite(t *testing.T) {
	queue, err := testClient.PublishSite(testSiteId, []string{testDomain})
	spew.Dump(queue)
	assert.NoError(t, err)
	assert.Equal(t, queue.Queued, true)
}

func TestListDomains(t *testing.T) {
	domains, err := testClient.ListDomains(testSiteId)
	spew.Dump(domains)
	assert.NoError(t, err)
	assert.Greater(t, len(domains), 0)
}

func TestListCollections(t *testing.T) {
	collections, err := testClient.ListCollections(testSiteId)
	spew.Dump(collections)
	assert.NoError(t, err)
	assert.Greater(t, len(collections), 0)
}

func TestGetCollection(t *testing.T) {
	collection, err := testClient.GetCollection(testCollectionId)
	spew.Dump(collection)
	assert.NoError(t, err)
	assert.NotNil(t, collection)
	assert.Equal(t, testCollectionId, collection.Id)
}

func TestListItems(t *testing.T) {
	resp, err := testClient.ListItems(testCollectionId, 0, 10)
	spew.Dump(resp)
	assert.NoError(t, err)
	assert.Greater(t, len(resp.Items), 0)
}

func TestGetItem(t *testing.T) {
	item, err := testClient.GetItem(testCollectionId, testItemId)
	spew.Dump(item)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, testItemId, item.Id())
}

func TestCreateItem(t *testing.T) {
	req := &RequestItem{
		Fields: &Item{
			"name":      "Test",
			"slug":      "Create",
			"post-body": "This is a create test",
			"_archived": false,
			"_draft":    true,
		},
	}
	item, err := testClient.CreateItem(testCollectionId, true, req)
	spew.Dump(item)
	assert.NoError(t, err)
	assert.NotNil(t, item)
}

func TestUpdateItem(t *testing.T) {
	req := &RequestItem{
		Fields: &Item{
			"name":      "Test",
			"slug":      "Update",
			"post-body": "This is an update test",
			"_archived": false,
			"_draft":    true,
		},
	}
	item, err := testClient.UpdateItem(testCollectionId, testItemId, true, req)
	spew.Dump(item)
	assert.NoError(t, err)
	assert.NotNil(t, item)
}

func TestRemoveItem(t *testing.T) {
	resp, err := testClient.RemoveItem(testCollectionId, testItemId, false)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, resp.Deleted)
}
