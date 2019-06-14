# bulrush-addition
Provides cross-module references.  

### logger
```
logDir := path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string))
transports := []*logger.Transport{
	// only for error
	&logger.Transport{
		Dirname: path.Join(logDir, "error"),
		Level:   logger.ERRORLevel,
		Maxsize: logger.Maxsize,
	},
	// combined all level
	&logger.Transport{
		Dirname: path.Join(logDir, "combined"),
		Level:   logger.INFOLevel,
		Maxsize: logger.Maxsize,
	},
	// console level
	&logger.Transport{
		Level: logger.INFOLevel,
	},
}
logger := logger.CreateLogger(
	logger.INFOLevel,
	nil,
	transports,
)
logger.Info("after")
```
### mongo
```go
mongo := mongo.New(conf.Cfg)
func AddUsers(users []interface{}) {
	User:= mongo.Model("user")
	err := User.Insert(users...)
	if err != nil {
		panic(err)
	}
}
// RegisterUser genrate user routers
func RegisterUser(r *gin.RouterGroup) {
	mongo.API.List(r, "user").Pre(func(c *gin.Context) {
		fmt.Println("before")
	}).Post(func(c *gin.Context) {
		fmt.Println("after")
	})
	mongo.API.One(r, "user")
	mongo.API.Create(r, "user")
	mongo.API.Update(r, "user")
	mongo.API.Delete(r, "user")
}
app.Use(Model, Route)
// Open model autoHook
app.Use(mongo.AutoHook)
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