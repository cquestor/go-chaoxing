package chaoxing

import (
	"errors"
	"net/http"
	"strings"
)

// 登录方法获取Cookie
func Login(uname, password string) (map[string]string, error) {
	password, err := EncryptDesECB([]byte(password), []byte(LOGIN_KEY))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	payload := "fid=1283&uname=" + uname + "&password=" + password + "&refer=http%253A%252F%252Fmooc1-1.chaoxing.com%252Fvisit%252Finteraction%253Fs%253D52246a6dac0400616184f6ca8cddac96&t=true&forbidotherlogin=0&validate=&doubleFactorLogin=0&independentId=0"
	request, _ := http.NewRequest("GET", URL_LOGIN, strings.NewReader(payload))
	return nil, nil
}
