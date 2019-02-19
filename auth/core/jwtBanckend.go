package core

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/edumar111/fastpv-auth/auth/model"
	"github.com/edumar111/fastpv-auth/auth/redis"
	"github.com/edumar111/fastpv-auth/tools"
	"log"

	"github.com/edumar111/fastpv-auth/settings"
	"os"
	"time"
)


type JWTAuthenticationBackend struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)


var authBackendInstance *JWTAuthenticationBackend = nil

//InitJWTAuthenticationBackend load  keys Private and public
func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	log.Println("InitJWTAuthenticationBackend")
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			PrivateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return authBackendInstance
}

//GenerateToken generate token
func (backend *JWTAuthenticationBackend) GenerateToken(username string) (string, error) {
	log.Println("GenerateToken")
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix(),
		"iat": time.Now().Unix(),
		"sub": username,
	}
	tokenString, err := token.SignedString(backend.PrivateKey)
	if err != nil {
		panic(err)
		return "", err
	}

	return tokenString, nil
}


//Authenticate validate usernama and password
func (backend *JWTAuthenticationBackend) Authenticate(user *model.UserLogin) bool {

	hashedPassword, _ := tools.HashPassword("123456")

	testUser := model.UserLogin{
		ID:     0,
		Username: "edumar111",
		Password: string(hashedPassword),
	}

	//return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
	return user.Username == testUser.Username && tools.CheckHashAndPassword(testUser.Password,user.Password)

	}

func (backend *JWTAuthenticationBackend) IsInBlacklist(token string) bool {

	redisConn := redis.Connect()
	redisToken, error := redisConn.GetValue(token)
	if error != nil{
		log.Println(error)
	}
	log.Println(redisToken)
	if redisToken == nil {
		return false
	}

	return true
}

func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {

	redisConn := redis.Connect()

	return redisConn.SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"]))

}
func (backend *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}
//getPrivateKey get private key RSA
func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}
//getPublicKey get public key RSA
func getPublicKey() *rsa.PublicKey {

	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
