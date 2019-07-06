## Directory

<!-- TOC -->

- [Directory](#directory)
- [bulrush-addition](#bulrush-addition)
    - [mgo](#mgo)
    - [gorm](#gorm)
    - [logger](#logger)
    - [redis](#redis)
    - [apidoc](#apidoc)
        - [Install apidoc](#install-apidoc)
        - [Add ignore to .igonre file](#add-ignore-to-igonre-file)
        - [Generate apidoc](#generate-apidoc)
        - [Use apidoc plugin](#use-apidoc-plugin)
- [MIT License](#mit-license)

<!-- /TOC -->

## bulrush-addition
	Provides the ability to expose default interfaces based on database-driven wrappers

### mgo

**create mgoext**
```go
var MGOExt = mgoext.New(conf.Cfg)
```

**use as a bulrush plugin**
```go
app.PostUse(addition.MGOExt)
```

**defined model and custom your own config**
```go
type User struct {
	Base     `bson:",inline"`
	Name     string          `bson:"name" form:"name" json:"name" xml:"name"`
	Password string          `bson:"password" form:"password" json:"password" xml:"password" `
	Age      int             `bson:"age" form:"age" json:"age" xml:"age"`
	Roles    []bson.ObjectId `ref:"role" bson:"roles" form:"roles" json:"roles" xml:"roles" `
}

var _ = addition.MGOExt.Register(&mgoext.Profile{
	DB:        "test",
	Name:      "user",
	Reflector: &User{},
	BanHook:   true,
})

// RegisterUser inject function
func RegisterUser(r *gin.RouterGroup) {
	addition.MGOExt.API.List(r, "user").Pre(func(c *gin.Context) {
		addition.Logger.Info("before")
	}).Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	})
	addition.MGOExt.API.Feature("feature").List(r, "user")
	addition.MGOExt.API.One(r, "user")
	addition.MGOExt.API.Create(r, "user")
	addition.MGOExt.API.Update(r, "user")
	addition.MGOExt.API.Delete(r, "user")
}
```

### gorm

**create gormext**
```go
var GORMExt = gormext.New(gormConf)
var _ = GORMExt.Init(func(ext *gormext.GORM) {
	ext.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	ext.API.Opts.Prefix = "/template/gorm"
	ext.API.Opts.RouteHooks = &gormext.RouteHooks{
		List: &gormext.ListHook{
			Pre: func(c *gin.Context) {
				Logger.Info("all gormext before")
			},
		},
	}
})
```
**use as a bulrush plugin**
```go
app.PostUse(addition.GORMExt)
```

```go
type User struct {
	Base
	Name string `form:"name" json:"name" xml:"name"`
	Age  uint   `form:"age" json:"age" xml:"age"`
}

var _ = addition.GORMExt.Register(&gormext.Profile{
	DB:        "test",
	Name:      "user",
	Reflector: &User{},
	BanHook:   true,
})

**defined model and custom your own config**
// RegisterUser inject function
func RegisterUser(r *gin.RouterGroup) {
	addition.GORMExt.API.List(r, "user").Pre(func(c *gin.Context) {
		addition.Logger.Info("before")
	}).Post(func(c *gin.Context) {
		addition.Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	})
	addition.GORMExt.API.Feature("subUser").List(r, "user")
	addition.GORMExt.API.One(r, "user")
	addition.GORMExt.API.Create(r, "user")
	addition.GORMExt.API.Update(r, "user")
	addition.GORMExt.API.Delete(r, "user")
}
```
### logger
```
// Logger application logger
var Logger = addition.RushLogger.AppendTransports([]*logger.Transport{
	&logger.Transport{
		Dirname: path.Join(path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string)), "error"),
		Level:   logger.ERROR,
		Maxsize: logger.Maxsize,
	},
	&logger.Transport{
		Dirname: path.Join(path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string)), "combined"),
		Level:   logger.SILLY,
		Maxsize: logger.Maxsize,
	},
	// console level
	// &logger.Transport{
	// 	Level: logger.SILLY,
	// },
}...)
Logger.Info("after")
```

### redis
```go
redis := redis.New(conf.Cfg)
rules := []limit.Rule{
	limit.Rule{
		Methods: []string{"GET"},
		Match:   "/api/v1/user*",
		Rate:    1,
	},
	limit.Rule{
		Methods: []string{"GET"},
		Match:   "/api/v1/role*",
		Rate:    2,
	},
}
app.Use(&limit.Limit{
	Frequency: &limit.Frequency{
		Passages: []string{},
		Rules: rules,
		Model: &limit.RedisModel{
			Redis: redis,
		},
	},
})
```

### apidoc

#### Install apidoc
```shell
npm install apidoc -g
```
#### Add ignore to .igonre file
```txt
/doc/*
!/doc/api_data.js
!/doc/api_project.js
```
#### Generate apidoc 
```shell
apidoc
```
	apidoc will generate doc dir and some files in doc dir

#### Use apidoc plugin

```so
// APIDoc defined http rest api
// APIDoc defined http rest api
var APIDoc = apidoc.New()
var _ = APIDoc.
	Config(path.Join("", "doc")).
	Init(func(ctx *apidoc.APIDoc) {
		ctx.Prefix = "/docs"
	})
app.Use(APIDoc)
```

## MIT License

Copyright (c) 2018-2020 Double

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.