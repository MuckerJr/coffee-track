package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Login Handler
func Login(c *gin.Context) {
	domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")
	redirectURI := os.Getenv("AUTH0_CALLBACK_URL")

	authURL := "https://" + domain + "/authorize?response_type=code&client_id=" + clientID + "&redirect_uri=" + redirectURI + "&scope=openid profile email"

	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

// Callback Handler
func Callback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth code not provided"})
		return
	}

	// Exchange the code for a token
	token, err := exchangeCodeForToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// exchangeCodeForToken exchanges the code for a token
func exchangeCodeForToken(code string) (string, error) {

	// Look at the Auth0 documentation to see how to exchange the code for a token
	return "", nil
}
