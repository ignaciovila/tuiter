package user_test

import (
	"testing"

	"github.com/ignaciovila/tuiter/src/user"
)

func TestAddUserToRegistry(t *testing.T) {
	user.AddUser(user.NewUser("nacho vila", "nacho@ml.com", "nacho", "pass"))

	if len(user.GetUsers()) != 1 {
		t.Error("El usuario no se ha agregado correctamente")
	}

	if user.GetUsers()[0].Name != "nacho vila" {
		t.Error("El nombre del usuario es incorrecto")
	}
}

func TestExistsUser(t *testing.T) {
	if !user.ExistsUser("nacho") {
		t.Error("El usuario no existe")
	}
}