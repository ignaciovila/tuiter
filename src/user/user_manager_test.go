package userManager_test

import (
	"testing"

	"github.com/ignaciovila/tuiter/src/user"
)

func TestAddUserToRegistry(t *testing.T) {
	userManager.AddUser(userManager.NewUser("nacho vila", "nacho@ml.com", "nacho", "pass"))

	if len(userManager.GetUsers()) != 1 {
		t.Error("El usuario no se ha agregado correctamente")
	}

	if userManager.GetUsers()[0].Name != "nacho vila" {
		t.Error("El nombre del usuario es incorrecto")
	}
}
