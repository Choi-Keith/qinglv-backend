package password

import (
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(UserPass, LoginPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(UserPass), []byte(LoginPass))
	logx.Debugf("err:%v\n", err)
	return err == nil
}

// encryptPassword
// The password does irreversible encryption.
func EncryptPassword(Pass string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(Pass), bcrypt.DefaultCost)
	// This encrypted string can be saved to the database and can be used as password matching verification
	return string(hashPwd), err
}
