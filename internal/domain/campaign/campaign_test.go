package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campanha 1"
	content  = "Conteudo"
	contacts = []string{"fabio@email.com", "maria@email.com"}
)

func Test_Campaign_New(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.Greater(campaign.CreatedOn, time.Now().Add(-time.Minute))
}

func Test_Campaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("Name is required with min 5", err.Error())
}

func Test_Campaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("Content is required with min 5", err.Error())
}

func Test_Campaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("Contacts is required with min 1", err.Error())
}
