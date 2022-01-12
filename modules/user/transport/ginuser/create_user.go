package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-curd/common"
	"simple-rest-curd/component"
	"simple-rest-curd/modules/user/biz"
	"simple-rest-curd/modules/user/model"
	"simple-rest-curd/modules/user/storage"
)

func CreateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data model.UserCreate

		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}

		store := storage.NewSqlStore(appCtx.GetMainDBConnection())
		userBiz := biz.NewCreateUserBiz(store)

		if err := userBiz.CreateUser(context.Request.Context(), &data); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, common.DataSuccessResponse(data))
	}
}
