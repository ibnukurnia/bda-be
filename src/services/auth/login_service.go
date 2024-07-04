package auth

import (
	"bytes"
	"encoding/json"
	"msbda/pkg/e"
	"msbda/src/dtos/requests"
	"msbda/src/dtos/responses"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Login(pernr, password string) (*responses.Auth, error) {
	if pernr != "admin@admin.com" || password != "password123" {
		return nil, &e.SchemaError{
			Status:  400,
			Message: "email or password not matched",
		}
	}

	// res, err := reqLdap(requests.LoginRequest{Pernr: pernr, Password: password})

	// if err != nil {
	// 	return nil, &e.SchemaError{
	// 		Status:  res.StatusCode,
	// 		Message: err.Error(),
	// 	}
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pernr": pernr,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return nil, &e.SchemaError{
			Status:  500,
			Message: "failed to sign token",
		}
	}

	return &responses.Auth{
		Token: tokenStr,
	}, nil
}

func reqLdap(req requests.LoginRequest) (*responses.LdapResponse, error) {
	resp := responses.LdapResponse{
		StatusCode: 500,
	}

	client := &http.Client{}

	ldapUrl := os.Getenv("LDAP_ENDPOINT")

	jsonReq, err := json.Marshal(req)
	if err != nil {
		return &resp, err
	}

	request, err := http.NewRequest("POST", ldapUrl, bytes.NewBuffer(jsonReq))
	if err != nil {
		return &resp, err
	}

	request.SetBasicAuth("username", "password")
	response, err := client.Do(request)

	if err != nil {
		return &resp, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return &resp, err
	}

	if resp.ResponseCode != "00" {
		resp.StatusCode = 400
	} else {
		resp.StatusCode = 200
	}

	return &resp, nil
}
