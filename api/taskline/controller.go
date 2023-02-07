package taskline

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	taskline "tdp-cloud/module/dborm/taskline"
)

// 记录列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if lst, err := taskline.FetchAll(userId); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取记录

func detail(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if res, err := taskline.Fetch(id, userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加记录

func create(c *gin.Context) {

	var rq *taskline.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := taskline.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
	} else {
		c.Set("Error", err)
	}

}

// 修改记录

func update(c *gin.Context) {

	var rq *taskline.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := taskline.Update(rq); err == nil {
		c.Set("Message", "更新成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除记录

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if err := taskline.Delete(id, userId); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}