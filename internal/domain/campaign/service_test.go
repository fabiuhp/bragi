package campaign

import (
	"bragi/internal/contract"
	"bragi/internal/internalerrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	validCampaign = contract.NewCampaign{
		Name:    "Teste A",
		Content: "Corpo",
		Emails:  []string{"fabio@email.com"},
	}
	invalidCampaignName = contract.NewCampaign{
		Name:    "",
		Content: "Corpo",
		Emails:  []string{"fabio@email.com"},
	}
)

func setup() (*repositoryMock, Service) {
	repositoryMock := new(repositoryMock)
	service := Service{Repository: repositoryMock}
	return repositoryMock, service
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock, service := setup()
	repositoryMock.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(validCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock, service := setup()
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == validCampaign.Name &&
			campaign.Content == validCampaign.Content &&
			len(campaign.Contacts) == len(validCampaign.Emails)
	})).Return(nil)

	service.Create(validCampaign)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	_, service := setup()
	_, err := service.Create(invalidCampaignName)

	assert.NotNil(err)
	assert.Equal("nome é obrigatório", err.Error())
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock, service := setup()
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	_, err := service.Create(validCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}
