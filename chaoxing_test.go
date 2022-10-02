package chaoxing_test

import (
	"fmt"
	"testing"

	"github.com/cquestor/go-chaoxing"
)

func TestMain(m *testing.M) {
	result, err := chaoxing.EncryptDesECB([]byte("123456"), []byte(chaoxing.LOGIN_KEY))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
	chaoxing.Login("13914308903", "app5896302")
}
