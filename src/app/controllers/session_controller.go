
func (_ AuthController) CreateSession(c *gin.Context) {
}

func (_ AuthController) DestroySession(c *gin.Context) {
	c.Redirect(302, "/users")
}

func (_ AuthController) SessionCheck(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("UserId")

	if uid == nil {
		c.Redirect(302, "/sing_in")
		c.Abort()
	} else {
		c.Set("UserId", uid)
		c.Next()
	}
}
