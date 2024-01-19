package helpers

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/agenda_mid/models"
)

func ListarContactosParametros() (contactosParametros models.ContactosParametros, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "ListarContactosParametros", "err": err, "Status": "500"}
			panic(outputError)
		}
	}()

	var contactos []models.Contactos
	var parametros []models.Parametros

	url := "contactos?query=IdCargo__IdCargo:1&limit=0"
	if err := GetRequestNew("UrlCrudContactos", url, &contactos); err != nil {
		logs.Error(err)
		panic(err.Error())
	}
	fmt.Println("Contactos ", contactos)

	urlParametros := "parametros?limit=0"
	if err := GetRequestNew("UrlCrudParametros", urlParametros, &parametros); err != nil {
		logs.Error(err)
		panic(err.Error())
	}
	fmt.Println("Parametros ", parametros)

	contactosParametros.Contactos = contactos
	contactosParametros.Parametros = parametros

	fmt.Println(contactosParametros)

	return contactosParametros, outputError
}
