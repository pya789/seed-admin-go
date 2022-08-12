package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	useRecovery(r)
	useLogger(r)
	useCors(r)
}
