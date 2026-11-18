package middleware

import "github.com/gin-gonic/gin"

// OptionalAuth is a placeholder for JWT/session auth and RBAC.
// In production, validate tokens and enforce roles (particulier, professionnel, salarie, admin).
func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) { c.Next() }
}

func RequireRole(roles ...string) gin.HandlerFunc {
	allowed := map[string]bool{}
	for _, r := range roles {
		allowed[r] = true
	}
	return func(c *gin.Context) { _ = allowed; c.Next() }
}
