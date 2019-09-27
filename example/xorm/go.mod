module github.com/2637309949/bulrush-template

go 1.12

// ## just for dev
replace github.com/2637309949/bulrush => ../../../bulrush

replace github.com/2637309949/bulrush-addition => ../../../bulrush-addition

// ## end

require (
	github.com/2637309949/bulrush v0.0.0-20190923101754-9016de9ddc56
	github.com/2637309949/bulrush-addition v0.0.0-20190923102210-0242705a86a6
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.7
	github.com/kataras/go-events v0.0.2
	github.com/leodido/go-urn v1.1.0 // indirect
	gopkg.in/guregu/null.v3 v3.4.0
)
