package model_test

import (
	"testing"

	"github.com/rockenbach/imersao-fullstack-fullcycle/codepix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {

	name := "Charles"
	email := "charles@bla.com"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser(name, "")
	require.NotNil(t, err)

	_, err = model.NewUser("", email)
	require.NotNil(t, err)
}
