package Routes

import (
	"github.com/gin-gonic/gin"
	"main/Tools"
	"net/http"
)

func CreateInI(data *gin.Context) {
	var inifile Tools.InI
	if err := data.ShouldBind(&inifile); err != nil {
		data.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": 101, "msg": err.Error()})
	}
	res, err := Tools.CreateIni(inifile.Inidata)
	if err != nil {
		data.JSON(http.StatusOK, gin.H{"code": 101, "msg": "配置生成失败"})
	} else {
		data.JSON(http.StatusOK, gin.H{"code": 100, "msg": "OK", "data": res})
	}
}

func AnalyzeIni(data *gin.Context) {
	inifile, _, _ := data.Request.FormFile("inifile")
	res, err := Tools.AnalyzeInifile(inifile)
	if err != nil {
		data.JSON(http.StatusOK, gin.H{"code": 101, "msg": err.Error()})
	} else {
		data.JSON(http.StatusOK, gin.H{"code": 100, "msg": "OK", "data": res})
	}
}
