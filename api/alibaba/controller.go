package alibaba

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/helper/alibaba"
	"tdp-cloud/module/dborm/vendor"
)

func apiProxy(c *gin.Context) {

	userId := c.GetUint("UserId")
	vendorId := cast.ToUint(c.Param("id"))

	vendor, err := vendor.Fetch(vendorId, userId)

	if err != nil || vendor.Id == 0 {
		c.Set("Error", "厂商不存在")
		return
	}

	// 构造参数

	params := &alibaba.Params{
		SecretId:  vendor.SecretId,
		SecretKey: vendor.SecretKey,
	}

	if err = c.ShouldBindJSON(params); err != nil {
		c.Set("Error", err)
		return
	}

	// 发起请求

	if res, err := alibaba.Request(params); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
