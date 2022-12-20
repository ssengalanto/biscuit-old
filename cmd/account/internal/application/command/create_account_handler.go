package command

import (
	"context"
	"fmt"

	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
)

type CreateAccountCommandHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
}

func NewCreateAccountCommandHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		log:               logger,
		accountRepository: accountRepository,
	}
}

func (c *CreateAccountCommandHandler) Topic() string {
	return CommandCreateAccount
}

func (c *CreateAccountCommandHandler) Handle(
	ctx context.Context,
	request mediatr.Request,
) (any, error) {
	entity := account.Entity{}

	command, ok := request.(*CreateAccountCommand)
	if !ok {
		c.log.Error("invalid command", map[string]any{"command": command})
		return nil, fmt.Errorf("%w: command %s", errors.ErrInvalid, command.Topic())
	}

	acct := account.New(command.Email, command.Password, command.Active)
	err := acct.IsValid()
	if err != nil {
		c.log.Error("account is invalid", map[string]any{"account": acct, "error": err})
		return nil, err
	}

	pers := person.New(acct.ID, command.FirstName, command.LastName, command.Email, command.Phone, command.DateOfBirth)
	err = pers.IsValid()
	if err != nil {
		c.log.Error("person is invalid", map[string]any{"person": pers, "error": err})
		return nil, err
	}

	entity = acct
	entity.Person = &pers

	err = c.accountRepository.Create(ctx, entity)
	if err != nil {
		c.log.Error("account creation failed", map[string]any{"payload": entity, "error": err})
		return nil, err
	}

	return acct.ID.String(), err
}
