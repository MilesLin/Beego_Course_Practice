package controllers

import (
	"fmt"
	"lastchance/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// BankingController operations for Banking
type BankingController struct {
	beego.Controller
}

// @router /banking [get]
func (c *BankingController) ShowAccounts() {
	c.TplName = "banking.tpl"
}

func (c *BankingController) URLMapping() {
	c.Mapping("ShowAccounts", c.ShowAccounts)
}

// @router /api/transfer [post]
func (c *BankingController) Transfer() {
	var transfer models.Transfer
	json.Unmarshal(c.Ctx.Input.RequestBody, &transfer)
	fmt.Println(transfer)
	c.Ctx.WriteString("success")
}

// // Post ...
// // @Title Create
// // @Description create Banking
// // @Param	body		body 	models.Banking	true		"body for Banking content"
// // @Success 201 {object} models.Banking
// // @Failure 403 body is empty
// // @router / [post]
// func (c *BankingController) Post() {

// }

// // GetOne ...
// // @Title GetOne
// // @Description get Banking by id
// // @Param	id		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.Banking
// // @Failure 403 :id is empty
// // @router /:id [get]
// func (c *BankingController) GetOne() {

// }

// // GetAll ...
// // @Title GetAll
// // @Description get Banking
// // @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// // @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// // @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// // @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// // @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// // @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// // @Success 200 {object} models.Banking
// // @Failure 403
// // @router / [get]
// func (c *BankingController) GetAll() {

// }

// // Put ...
// // @Title Put
// // @Description update the Banking
// // @Param	id		path 	string	true		"The id you want to update"
// // @Param	body		body 	models.Banking	true		"body for Banking content"
// // @Success 200 {object} models.Banking
// // @Failure 403 :id is not int
// // @router /:id [put]
// func (c *BankingController) Put() {

// }

// // Delete ...
// // @Title Delete
// // @Description delete the Banking
// // @Param	id		path 	string	true		"The id you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 id is empty
// // @router /:id [delete]
// func (c *BankingController) Delete() {

// }
