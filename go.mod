module github.com/2637309949/bulrush-addition

go 1.12

require (
	github.com/2637309949/bulrush v0.0.0-20190615094031-919971fe3950
	github.com/gin-gonic/gin v1.4.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/jinzhu/gorm v1.9.9
	github.com/stretchr/testify v1.3.0
	github.com/thoas/go-funk v0.4.0
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/go-playground/validator.v9 v9.29.0
)

// ## just for dev
replace github.com/2637309949/bulrush => ../bulrush

// ## end

replace cloud.google.com/go => github.com/googleapis/google-cloud-go v0.40.0

replace google.golang.org/api => github.com/googleapis/google-api-go-client v0.6.0
