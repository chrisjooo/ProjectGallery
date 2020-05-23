package helpers

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	AccessToken   string
	AccessUuid    string
	AccessExpires int64
}

type AccessDetails struct {
	AccessUuid string
	Username   string
}

func CreateToken(username string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AccessExpires = time.Now().Add(time.Hour * 1).Unix()
	td.AccessUuid = uuid.NewV4().String()

	secret := beego.AppConfig.String("ACCESS_SECRET")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["username"] = username
	atClaims["exp"] = td.AccessExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		log.Printf("error CreateToken: %v\n", err)
		return nil, err
	}
	td.AccessToken = token
	return td, nil
}

func CreateAuth(username string, td *TokenDetails) error {
	conn := NewPool().Get()
	defer conn.Close()

	temp := time.Now().Unix()
	exp := td.AccessExpires - temp

	_, err := conn.Do("HSET", username, td.AccessUuid, td.AccessToken)
	if err != nil {
		log.Printf("error inserting token to redis: %v", err)
		return err
	}
	_, err = conn.Do("EXPIRE", username, exp)
	if err != nil {
		log.Printf("error inserting TTL to redis: %v", err)
		return err
	}

	return nil
}

func ExtractToken(ctx *context.Context) string {
	bearToken := ctx.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//Verify only
func VerifyToken(ctx *context.Context) (*jwt.Token, error) {
	secret := beego.AppConfig.String("ACCESS_SECRET")
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errMessage := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(errMessage)
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(ctx *context.Context) error {
	token, err := VerifyToken(ctx)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(ctx *context.Context) (*AccessDetails, error) {
	token, err := VerifyToken(ctx)
	if err != nil {
		return nil, errors.New("Unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, errors.New("Error extracting authorization data")
		}
		username := fmt.Sprintf("%s", claims["username"])
		if err != nil {
			return nil, errors.New("Error extracting authorization data")
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			Username:   username,
		}, nil
	}
	return nil, errors.New("Unauthorized")
}

func FetchAuth(authD *AccessDetails) error {
	conn := NewPool().Get()
	defer conn.Close()

	result, err := conn.Do("HGET", authD.Username, authD.AccessUuid)
	if err != nil || result == nil {
		return errors.New("Unauthorized")
	}
	log.Printf("result getting from redis2: %v\n", result)

	return nil
}

func DeleteAuth(username, givenUuid string) error {
	conn := NewPool().Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", username, givenUuid)
	if err != nil {
		log.Printf("error deleting redis: %v", err)
		return err
	}
	return nil
}
