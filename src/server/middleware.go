package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/esc-chula/gearfest-backend/src/utils"
	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

func Validation(supabase *supa.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "" && strings.HasPrefix(token, "Bearer ") {
			token := strings.TrimPrefix(token, "Bearer ")
			user, err := supabase.Auth.User(context.Background(), token)
			if err != nil {
				utils.HandleErrorResponse(ctx, http.StatusForbidden, "Failed to verify user.")
				return
			}
			ctx.Set("user_id", user.ID)
			if name, ok := user.UserMetadata["name"]; ok {
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
