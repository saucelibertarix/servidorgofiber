package services

import (
	"errors"
	database "github.com/saucelibertarix/servidorgofiber/database"
	"github.com/saucelibertarix/servidorgofiber/models"
	"github.com/saucelibertarix/servidorgofiber/utils"
)

func SignUpClient(signupClientRequest *models.SignupClientRequest) error {
	client := new(models.Client)

	// buscamos si existe en la db
	database.GormDB.Where("email = ?", signupClientRequest.Email).Find(&client)

	// si tiene id > 0 significa que existe
	if client.ID > 0 {
		return errors.New("Client ya existe")
	}

	//
	client.Name = signupClientRequest.Name
	client.Email = signupClientRequest.Email
	client.Password = utils.HashPassword(signupClientRequest.Password)

	// insertamos el objeto la base de datos
	createResult := database.GormDB.Create(&client)

	// devolvemos el resultando de la operación de insert
	return createResult.Error
}

func LoginClient(loginClientRequest *models.LoginClientRequest, ) (*models.LoginClientResponse, error) {
	clientBD := new(models.Client)

	database.GormDB.Where("email = ?", loginClientRequest.Email).Find(&clientBD)

	// si tiene id > 0 significa que existe
	if clientBD.ID < 1 {
		return nil, errors.New("Cliente no existe")
	}

	if !utils.ComparePasswords(clientBD.Password, loginClientRequest.Password) {
		return nil, errors.New("Contraseña incorrecta")
	}

	token, err := utils.GenerateClientToken(clientBD.Email, clientBD.ID, clientBD.Role)
	if err != nil {
		return nil, err
	}

	clientLoginResponse := &models.LoginClientResponse{
		Name:  clientBD.Name,
		Email: clientBD.Email,
		Token: token,
	}

	return clientLoginResponse, nil
}
