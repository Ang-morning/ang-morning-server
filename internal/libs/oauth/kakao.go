package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"angmorning.com/internal/config"
)

type KakaoOauthClient struct {
	httpClient *http.Client
}

func newKakaoClient() *KakaoOauthClient {
	return &KakaoOauthClient{
		httpClient: http.DefaultClient,
	}
}

func (kakao *KakaoOauthClient) GetToken(code string) string {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", config.Oauth.Kakao.ClientId)
	data.Set("client_secret", config.Oauth.Kakao.ClientSecret)
	data.Set("redirect_uri", config.Oauth.Kakao.RedirectUri)
	data.Set("code", code)

	res, err := kakao.httpClient.Post("https://kauth.kakao.com/oauth/token", "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)

		fmt.Println(string(body))
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	json.NewDecoder(res.Body).Decode(&result)

	return result.AccessToken
}

func (kakao *KakaoOauthClient) GetUserInfo(accessToken string) *OauthUserInfo {
	req, err := http.NewRequest("GET", "https://kapi.kakao.com/v2/user/me", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	res, err := kakao.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		fmt.Println("!!!", string(body))
	}

	return kakao.parseUserInfo(body)
}

func (kakao *KakaoOauthClient) parseUserInfo(body []byte) *OauthUserInfo {
	var result struct {
		KakaoAccount struct {
			Profile struct {
				Nickname        string `json:"nickname"`
				ProfileImageUrl string `json:"profile_image_url"`
			} `json:"profile"`
			Email string `json:"email"`
		} `json:"kakao_account"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil
	}

	userInfo := OauthUserInfo{
		Email:           result.KakaoAccount.Email,
		Nickname:        result.KakaoAccount.Profile.Nickname,
		ProfileImageUrl: result.KakaoAccount.Profile.ProfileImageUrl,
	}

	return &userInfo
}
