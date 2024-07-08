package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate: "email"`
}

type Campaign struct {
	ID        string    `validate: "required"`
	Name      string    `validate: "min=5,max=24"`
	CreatedOn time.Time `validate: "required"`
	Content   string    `validate: "min=5,max=1024"`
	Contacts  []Contact `validate: "min=1"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("nome é obrigatório")
	}

	if content == "" {
		return nil, errors.New("conteudo é obrigatório")
	}

	if len(emails) == 0 {
		return nil, errors.New("contatos são obrigatórios")
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}, nil
}
