package dto

type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginResponseOf(access, refresh string) loginResponse {
	return loginResponse{AccessToken: access, RefreshToken: refresh}
}

type refreshResponse struct {
	AccessToken string `json:"access_token"`
}

func RefreshResponseOf(access string) refreshResponse {
	return refreshResponse{AccessToken: access}
}
