package mongo

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

// Hook type defined
type Hook struct {
	One    func(name string) func(c *gin.Context)
	List   func(name string) func(c *gin.Context)
	Create func(name string) func(c *gin.Context)
	Update func(name string) func(c *gin.Context)
}

// PUFormat Post Update Format
type PUFormat struct {
	Cond map[string]interface{} `bson:"cond" form:"cond" json:"cond" xml:"cond"`
	Muti bool                   `bson:"muti" form:"muti" json:"muti" xml:"muti"`
	Doc  interface{}            `bson:"doc" form:"doc" json:"doc" xml:"doc" `
}

// list doc
func list(mgo *Mongo) func(string) func(c *gin.Context) {
	return func(name string) func(c *gin.Context) {
		return func(c *gin.Context) {
			var match map[string]interface{}
			Model, error := mgo.Model(name)
			if error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				return
			}
			list, error := mgo.Vars(name)
			if error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				return
			}
			cond := c.DefaultQuery("cond", "%7B%7D")
			page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
			size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
			_range := c.DefaultQuery("range", "PAGE")
			unescapeCond, err := url.QueryUnescape(cond)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			err = json.Unmarshal([]byte(unescapeCond), &match)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			// return mapping body， not json in db
			query := Model.Find(match)
			totalrecords, _ := query.Count()
			if _range != "ALL" {
				query = query.Skip((page - 1) * size).Limit(size)
			}
			err = query.All(&list)
			totalpages := math.Ceil(float64(totalrecords) / float64(size))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"range":        _range,
				"page":         page,
				"totalpages":   totalpages,
				"size":         size,
				"totalrecords": totalrecords,
				"cond":         match,
				"list":         list,
			})
		}
	}
}

// one doc
func one(mgo *Mongo) func(string) func(c *gin.Context) {
	return func(name string) func(c *gin.Context) {
		return func(c *gin.Context) {
			id := c.Param("id")
			Model, error := mgo.Model(name)
			if error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				return
			}
			one, error := mgo.Var(name)
			if error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				return
			}
			isOj := bson.IsObjectIdHex(id)
			if !isOj {
				c.JSON(http.StatusNotAcceptable, gin.H{
					"message": "not a valid id",
				})
				return
			}
			err := Model.FindId(bson.ObjectIdHex(id)).One(one)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, one)
		}
	}
}

// create doc
func create(mgo *Mongo) func(string) func(c *gin.Context) {
	return func(name string) func(c *gin.Context) {
		return func(c *gin.Context) {
			Model, error := mgo.Model(name)
			if error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				fmt.Printf("error: %s", error.Error())
				return
			}

			binds, error := mgo.Var(name)
			if error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				fmt.Printf("error: %s", error.Error())
				return
			}
			if error := c.ShouldBind(&binds); error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				fmt.Printf("error: %s", error.Error())
				return
			}
			// to array
			switch binds.(type) {
			case []interface{}:
			case interface{}:
				binds = []interface{}{binds}
			}
			if error := Model.Insert(binds.([]interface{})...); error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				fmt.Printf("error: %s", error.Error())
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
			return
		}
	}
}

// update doc
func update(mgo *Mongo) func(string) func(c *gin.Context) {
	return func(name string) func(c *gin.Context) {
		return func(c *gin.Context) {
			Model, error := mgo.Model(name)
			if error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				return
			}

			var puDate PUFormat
			if error := c.ShouldBind(&puDate); error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				fmt.Printf("error: %s", error.Error())
				return
			}

			bind, error := mgo.Vars(name)

			if error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": error.Error(),
				})
				fmt.Printf("error: %s", error.Error())
				return
			}

			jsonUpData, _ := json.Marshal(puDate.Doc)
			json.Unmarshal(jsonUpData, &bind)

			if puDate.Muti {
				if _, error := Model.UpdateAll(puDate.Cond, &bind); error != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": error.Error(),
					})
					fmt.Printf("error: %s", error.Error())
					return
				}
			} else {
				if error := Model.Update(puDate.Cond, &bind); error != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": error.Error(),
					})
					fmt.Printf("error: %s", error.Error())
					return
				}
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
			return
		}
	}
}
