package controllers

import (
	"fmt"
	"lastchance/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/context"
)

// BankingController operations for Banking
type BankingController struct {
	beego.Controller
}

// @router /banking [get]
func (c *BankingController) ShowAccounts() {
	c.Data["accounts"] = []models.Account{
		models.Account{
			ID:		1,
			Name:	"Checking",
			Number:	"8888",
			Amount:	642.27,
		},
		models.Account{
			ID:		2,
			Name:	"Savings",
			Number:	"33449",
			Amount:	10000,
		},
	}
	c.TplName = "banking.tpl"
}

func (c *BankingController) URLMapping() {
	c.Mapping("ShowAccounts", c.ShowAccounts)
}

// @router /api/transfer [post]
func (c *BankingController) Transfer() {
	var transfer models.Transfer
	json.Unmarshal(c.Ctx.Input.RequestBody, &transfer)

	valid := validation.Validation{}
	isValid, _ := valid.Valid(&transfer)
	fmt.Println(transfer)
	fmt.Println(valid.ErrorMap())

	response := models.TransferResponse{
		Transfer: transfer,
	}

	if isValid {
		response.Status = "success"
	} else {
		response.Status = "failure"
	}
	c.Data["json"] = response;
	c.ServeJSON()
}

// @router /lifecycle [get]
func (c *BankingController) ShowLifecycle() {
	fmt.Println("Action Execution")
}

func (c *BankingController) Init(ctx *context.Context, controllerName,
actionName string, app interface{}) {
	fmt.Printf("Initialization: %s.$s\n", controllerName, actionName)
	c.Controller.Init(ctx, controllerName, actionName, app)
}

func (c *BankingController) Prepare() {
	fmt.Println("Prepare Controller")
	c.Controller.Prepare()
}

func (c *BankingController) Render() error{
	fmt.Println("Render result")
	c.Ctx.WriteString("result")
	
	return nil
}

func (c *BankingController) Finish() {
	fmt.Println("Finsih Controller")
	c.Controller.Finish()
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
