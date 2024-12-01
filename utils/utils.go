package utils

import (
	"crypto/md5"
	cryptoRand "crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"general_spider_controll_panel/types"
	mathRand "math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/joho/godotenv"

	"golang.org/x/crypto/bcrypt"
)

type Env struct {
	value map[string]string
	mu    sync.Mutex
}

var env *Env

func init() {
	env = &Env{value: map[string]string{}}
}

func ClientIP(request *http.Request) string {
	ip := request.Header.Get("Cf-Connecting-IP")
	if ip != "" {
		return ip
	}
	ip = request.Header.Get("X-Real-IP")
	if ip == "" {
		ip = request.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = request.RemoteAddr
		}
	}

	if strings.Contains(ip, ",") {
		ips := strings.Split(ip, ",")
		ip = strings.TrimSpace(ips[0])
	}

	if strings.Contains(ip, ":") {
		ips := strings.Split(ip, ":")
		ip = ips[0]
	}

	return ip
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasNumber    bool
		hasUppercase bool
	)

	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUppercase = true
		}
	}

	return hasNumber && hasUppercase
}

func ConvertFileSize[T types.Number](size T) string {
	sizeInBytes := int64(size)

	if sizeInBytes < 1024 {
		return fmt.Sprintf("%d B", sizeInBytes)
	} else if sizeInBytes < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(sizeInBytes)/1024)
	} else if sizeInBytes < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(sizeInBytes)/(1024*1024))
	} else {
		return fmt.Sprintf("%.2f GB", float64(sizeInBytes)/(1024*1024*1024))
	}
}

func ConvertTIme(timestampMillis int64) string {
	t := time.UnixMilli(timestampMillis)

	now := time.Now()
	duration := now.Sub(t)

	if duration < time.Minute {
		return fmt.Sprintf("%d seconds ago", int(duration.Seconds()))
	} else if duration < time.Hour {
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	} else if duration < time.Hour*24 {
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	} else if duration < time.Hour*24*30 {
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	} else if duration < time.Hour*24*365 {
		return fmt.Sprintf("%d months ago", int(duration.Hours()/(24*30)))
	} else {
		return fmt.Sprintf("%d years ago", int(duration.Hours()/(24*365)))
	}
}

func TimeUntil(ms int64) string {
	if ms < 0 {
		return "Time has already passed!"
	}

	seconds := ms / 1000

	months := seconds / (30 * 24 * 60 * 60)
	seconds %= 30 * 24 * 60 * 60
	days := seconds / (24 * 60 * 60)
	seconds %= 24 * 60 * 60
	hours := seconds / (60 * 60)
	seconds %= 60 * 60
	minutes := seconds / 60
	seconds %= 60

	result := ""
	if months > 0 {
		result += fmt.Sprintf("%d months ", months)
	}
	if days > 0 {
		result += fmt.Sprintf("%d days ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%d hours ", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%d minutes ", minutes)
	}
	if seconds > 0 || result == "" {
		result += fmt.Sprintf("%d seconds", seconds)
	}

	return "in " + result
}

func Getenv(key string) string {
	env.mu.Lock()
	defer env.mu.Unlock()
	if val, ok := env.value[key]; ok {
		return val
	}

	if os.Getenv("HOSTNAME") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file: " + err.Error())
		}
	}

	val := os.Getenv(key)
	env.value[key] = val

	if val == "" {
		panic("Asking for env: " + key + " but got nothing, please set your environment first")
	}

	return val
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	seededRand := mathRand.New(mathRand.NewSource(time.Now().UnixNano() + int64(mathRand.Intn(9999))))
	var result strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := seededRand.Intn(len(charset))
		result.WriteString(string(charset[randomIndex]))
	}
	return result.String()
}

func GenerateCSRFToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := cryptoRand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write(tokenBytes)
	hashedToken := hash.Sum(nil)

	csrfToken := base64.URLEncoding.EncodeToString(hashedToken)

	return csrfToken, nil
}

func ParseUserAgent(userAgent string) (map[string]string, map[string]string) {
	browserInfo := make(map[string]string)
	osInfo := make(map[string]string)
	if strings.Contains(userAgent, "Firefox") {
		browserInfo["browser"] = "Firefox"
		parts := strings.Split(userAgent, "Firefox/")
		if len(parts) > 1 {
			version := strings.Split(parts[1], " ")[0]
			browserInfo["version"] = version
		}
	} else if strings.Contains(userAgent, "Chrome") {
		browserInfo["browser"] = "Chrome"
		parts := strings.Split(userAgent, "Chrome/")
		if len(parts) > 1 {
			version := strings.Split(parts[1], " ")[0]
			browserInfo["version"] = version
		}
	} else {
		browserInfo["browser"] = "Unknown"
		browserInfo["version"] = "Unknown"
	}

	if strings.Contains(userAgent, "Windows") {
		osInfo["os"] = "Windows"
		parts := strings.Split(userAgent, "Windows ")
		if len(parts) > 1 {
			version := strings.Split(parts[1], ";")[0]
			osInfo["version"] = version
		}
	} else if strings.Contains(userAgent, "Macintosh") {
		osInfo["os"] = "Mac OS"
		parts := strings.Split(userAgent, "Mac OS X ")
		if len(parts) > 1 {
			version := strings.Split(parts[1], ";")[0]
			osInfo["version"] = version
		}
	} else if strings.Contains(userAgent, "Linux") {
		osInfo["os"] = "Linux"
		osInfo["version"] = "Unknown"
	} else if strings.Contains(userAgent, "Android") {
		osInfo["os"] = "Android"
		parts := strings.Split(userAgent, "Android ")
		if len(parts) > 1 {
			version := strings.Split(parts[1], ";")[0]
			osInfo["version"] = version
		}
	} else if strings.Contains(userAgent, "iPhone") || strings.Contains(userAgent, "iPad") || strings.Contains(userAgent, "iPod") {
		osInfo["os"] = "iOS"
		parts := strings.Split(userAgent, "OS ")
		if len(parts) > 1 {
			version := strings.Split(parts[1], " ")[0]
			osInfo["version"] = version
		}
	} else {
		osInfo["os"] = "Unknown"
		osInfo["version"] = "Unknown"
	}

	return browserInfo, osInfo
}

func IntToString[T types.Number](number T) string {
	return fmt.Sprintf("%d", number)
}

func FloatToString[T types.Float](number T) string {
	return fmt.Sprintf("%.2f", number)
}
