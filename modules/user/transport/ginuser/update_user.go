package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-curd/component"
	"simple-rest-curd/modules/user/biz"
	"simple-rest-curd/modules/user/model"
	"simple-rest-curd/modules/user/storage"
	"strconv"
)

func UpdateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		var data model.UserUpdate
		store := storage.NewSqlStore(appCtx.GetMainDBConnection())
		updateBiz := biz.NewUpdateUserBiz(store)

		updateBiz.UpdateUser(context.Request.Context(), id, &data)

	}
}
