package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-curd/common"
	"simple-rest-curd/component"
	biz2 "simple-rest-curd/modules/user/biz"
	"simple-rest-curd/modules/user/model"
	"simple-rest-curd/modules/user/storage"
)

func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var filter model.Filter
		if err := context.ShouldBind(&filter); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		var paging common.Paging
		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fulfill()
		store := storage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := biz2.NewListUserBiz(store)

		results, err := biz.ListUser(context.Request.Context(), &paging, &filter)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, common.SuccessResponse(results, paging, filter))
	}
}
