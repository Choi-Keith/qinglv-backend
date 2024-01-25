package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net"
	"net/http"
	"regexp"
	"strings"
)

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

// CheckPassword
// minLength: Specifies the minimum length of a password
// maxLength：Specifies the maximum length of a password
// minLevel：Specifies the minimum strength level required for passwords
// pwd：Text passwords
func CheckPassword(minLength, maxLength, minLevel int, pwd string) error {
	// First check whether the password length is within the range
	if len(pwd) < minLength {
		return fmt.Errorf("BAD PASSWORD: The password is shorter than %d characters", minLength)
	}
	if len(pwd) > maxLength {
		return fmt.Errorf("BAD PASSWORD: The password is logner than %d characters", maxLength)
	}

	// The password strength level is initialized to D.
	// The regular is used to verify the password strength.
	// If the matching is successful, the password strength increases by 1
	level := levelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}

	// If the final password strength falls below the required minimum strength, return with an error
	if level < minLevel {
		return fmt.Errorf("the password does not satisfy the current policy requirements")
	}
	return nil
}

// write email checker

func CheckEmail(email string) bool {
	pattern := `^[0-9a-z-A-Z_]+@[a-zA-Z\d.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func GetClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""

}

const (
	defaultMaxMemory = 36 << 20 // 64 MB
)

func GetFormFile(r *http.Request, key string) ([]multipart.File, []*multipart.FileHeader, error) {
	if r.MultipartForm == nil {
		err := r.ParseMultipartForm(defaultMaxMemory)
		if err != nil {
			return nil, nil, err
		}
	}
	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		if fhs := r.MultipartForm.File[key]; len(fhs) > 0 {
			files := make([]multipart.File, len(fhs))
			for i, _ := range fhs {
				f, err := fhs[i].Open()
				if err != nil {
					return nil, nil, err
				}
				files = append(files, f)
			}
			return files, fhs, nil

		}
	}
	return nil, nil, errors.New("http: no such file")
}
