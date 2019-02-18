package addition

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

// mgoHooks -
type mgoHooks struct {
	One  func(name string) func(c *gin.Context)
	List func(name string) func(c *gin.Context)
}

// list -
func list(mgo *Mgo) func(string) func(c *gin.Context) {
	return func(name string) func(c *gin.Context) {
		return func(c *gin.Context) {
			var match map[string]interface{}
			Model, manifest := mgo.Model(name)
			target := LeftOkV(manifest["reflector"])
			list := createSlice(target)

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

// one -
func one(mgo *Mgo) func(string) func(c *gin.Context) {
	return func(name string) func(c *gin.Context) {
		return func(c *gin.Context) {
			id := c.Param("id")
			Model, manifest := mgo.Model(name)
			target := LeftOkV(manifest["reflector"])
			one := createObject(target)
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

// create
func create(mgo *Mgo) func(string) func(c *gin.Context) {
	return func(name string) func(c *gin.Context) {
		return func(c *gin.Context) {

		}
	}
}
