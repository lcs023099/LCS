package router

import (
	"github.com/astaxie/beego"
	"secKill/SecAdmin/controller/product"
	"secKill/SecAdmin/controller/activity"
)

func init() {
	beego.Router("/product/list", &product.ProductController{}, "*:ListProduct")
	beego.Router("/", &product.ProductController{}, "*:ListProduct")
	beego.Router("/product/create", &product.ProductController{}, "*:CreateProduct")
	beego.Router("/product/submit", &product.ProductController{}, "*:SubmitProduct")

	beego.Router("/activity/create", &activity.ActivityController{}, "*:CreateActivity")
	beego.Router("/activity/list", &activity.ActivityController{}, "*:ListActivity")
	beego.Router("/activity/submit", &activity.ActivityController{}, "*:SubmitActivity")
}
