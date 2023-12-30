package password

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(LoginPass, UserPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(UserPass), []byte(LoginPass))
	return err == nil
}

// encryptPassword
// The password does irreversible encryption.
func EncryptPassword(Pass string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(Pass), bcrypt.DefaultCost)
	// This encrypted string can be saved to the database and can be used as password matching verification
	return string(hashPwd), err
}
