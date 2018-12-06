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

	if user.GetUsers()[0].Nombre != "nacho vila" {
		t.Error("EL nombre del usuario es incorrecto")
	}
}
