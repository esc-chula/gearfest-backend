package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/esc-chula/gearfest-backend/src/config"
	"github.com/esc-chula/gearfest-backend/src/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

func Validation(cfg config.GoogleConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "" && strings.HasPrefix(token, "Bearer ") {
			token := strings.TrimPrefix(token, "Bearer ")
			payload, err := idtoken.Validate(context.Background(), token, cfg.ClientID)
			if err != nil || cfg.ClientID != payload.Audience {
				utils.HandleErrorResponse(ctx, http.StatusUnauthorized, "Invalid token.")
				return
			}
			ctx.Set("user_id", payload.Subject)
			if name, ok := payload.Claims["name"]; ok {
				ctx.Set("user_google_name", name)
			} else {
				utils.HandleErrorResponse(ctx, http.StatusForbidden, "Name not found.")
				return
			}
			ctx.Next()
		} else {
			utils.HandleErrorResponse(ctx, http.StatusUnauthorized, "Invalid token.")
		}
	}
}
