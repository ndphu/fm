package model

type FacebookValidateResponse struct {
	FacebookAccessToken FacebookAccessToken `json:"data"`
}

type FacebookAccessToken struct {
	AppId       string                      `json:"app_id"`
	Type        string                      `json:"type"`
	Application string                      `json:"application"`
	IsValid     bool                        `json:"is_valid"`
	ExpiresAt   int64                       `json:"expiresAt"`
	IssuedAt    int64                       `json:"issued_at"`
	UserId      string                      `json:"user_id"`
	Scopes      []string                    `json:"scopes"`
	MetaData    FacebookAccessTokenMetaData `json:"metadata"`
}

type FacebookAccessTokenMetaData struct {
	SSO string `json:"sso"`
}

type CodeValidationResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}
