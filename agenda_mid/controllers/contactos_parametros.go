package controllers

import (
	"github.com/udistrital/agenda_mid/helpers"

	"github.com/astaxie/beego"
)

type ContactosParametrosController struct {
	beego.Controller
}

func (c *ContactosParametrosController) URLMapping() {
}

func (c *ContactosParametrosController) GetAll() {
	defer helpers.ErrorControllers(c.Controller, "ContactosParametrosController")
	beego.Debug("Entró al método GetAll en ContactosParametrosController")

	if v, err := helpers.ListarContactosParametros(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "Listado consultado con exito", "Data": v}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

