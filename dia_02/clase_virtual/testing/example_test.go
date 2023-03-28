package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {

	//Arrange
	var usr = "Agus"
	var usersTest = []string {
		"Vincent",
		"Sol",
		"Nico",
		"Mati",
		"Agus",
	}

	//Act
	addUser(usr)

	//Assert - se espera que el slice tenga tamaña 5
	assert.Equal(t, 5, len(users), "El tamaño debería ser 5")
	assert.Equal(t, usersTest, users, "Error en el test")

}
