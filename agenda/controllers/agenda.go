package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/agenda/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

// AgendaController operations for Agenda
type AgendaController struct {
	beego.Controller
}

// URLMapping ...
func (c *AgendaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Agenda
// @Param	body		body 	models.Agenda	true		"body for Agenda content"
// @Success 201 {int} models.Agenda
// @Failure 403 body is empty
// @router / [post]
func (c *AgendaController) Post() {
	var v models.Agenda
	v.FechaEvento = time_bogota.TiempoCorreccionFormato(v.FechaEvento)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddAgenda(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registro completado", "Data": v}
		} else {
			logs.Error(err)
			c.Data["Message"] = "Error en el servicio de POST: la respuesta contiene datos o parametros incorrectos"
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Data["Message"] = "Error en el servicio de POST: la respuesta contiene datos o parametros incorrectos"
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Agenda by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Agenda
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AgendaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAgendaById(id)
	if err != nil {
		logs.Error(err)
		c.Data["Message"] = "Error en el servicio de GetOne: la respuesta contiene datos o parametros incorrectos"
		c.Abort("400")
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registro completado", "Data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Agenda
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Agenda
// @Failure 403
// @router / [get]
func (c *AgendaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAgenda(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["Message"] = "Error en el servicio de GetAll: la respuesta contiene datos o parametros incorrectos"
		c.Abort("400")
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registro completado", "Data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Agenda
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Agenda	true		"body for Agenda content"
// @Success 200 {object} models.Agenda
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AgendaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Agenda{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaEvento = time_bogota.TiempoCorreccionFormato(v.FechaEvento)
		if err := models.UpdateAgendaById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registro completado", "Data": v}
		} else {
			logs.Error(err)
			c.Data["Message"] = "Error en el servicio de Put: la respuesta contiene datos o parametros incorrectos"
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Data["Message"] = "Error en el servicio de Put: la respuesta contiene datos o parametros incorrectos"
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Agenda
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AgendaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAgenda(id); err == nil {
		d := map[string]interface{}{"Id": id}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Eliminacion completada", "Data": d}
	} else {
		logs.Error(err)
		c.Data["Message"] = "Error en el servicio de Delete: la respuesta contiene datos o parametros incorrectos"
		c.Abort("400")
	}
	c.ServeJSON()
}
