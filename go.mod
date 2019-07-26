module github.com/2637309949/bulrush-addition

go 1.12

require (
	github.com/2637309949/bulrush v0.0.0-20190725143958-5a43e012d374
	github.com/2637309949/bulrush-utils v0.0.0-20190719014903-7f23f85d8694
	github.com/gin-gonic/gin v1.4.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/jinzhu/gorm v1.9.9
	github.com/jinzhu/inflection v1.0.0
	github.com/kataras/go-events v0.0.2
	github.com/mattn/go-isatty v0.0.8 // indirect
	github.com/mojocn/base64Captcha v0.0.0-20190716153509-e5e80f1b3816 // indirect
	github.com/thoas/go-funk v0.4.0
	golang.org/x/image v0.0.0-20190523035834-f03afa92d3ff // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7
	gopkg.in/go-playground/validator.v9 v9.29.0
)

replace cloud.google.com/go => github.com/googleapis/google-cloud-go v0.40.0

replace google.golang.org/api => github.com/googleapis/google-api-go-client v0.6.0
