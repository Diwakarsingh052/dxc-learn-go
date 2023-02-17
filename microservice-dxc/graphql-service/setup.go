package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"graphql-service/auth"
)

const privatePem = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDK7OxpEQBg+vIx
U9zm874g8LwatccbVRo/sP+ldZK17dikHzihFpQTcY7Nct9iTdNimP7gIS5oGX6O
d902MMfeM/hJH3RGb/peVLBW2VpiVM1ww/R79RgSlXEQsgsCBIweFUvRI5preonJ
RJmzy85J9XmHxfBazNkQCZ1kfyW1Fz5CGpiz+xE0UKFbdxooTabUHgxf9+EzKYfV
GsipT6kfVcJiUGwYHMvKX7qZTvYxckg7z0xrAYuXuavjC8+qmC+d5eo529DdjXhi
U3KEXjDSL9JSGT1cIEYZhdgB++4NoL+8WT/y8HdW202+iZN+N4xQe2QEBtrWPwt9
irzDv+VlAgMBAAECggEAUVKRi6mclUD8Pnh1VrjnwRu4xpuCp5l7Y3rzeMOdf/tJ
jrxUXXpG7WPc7sjSbPrzc9gMjJW/m0hcR4wRwt3Zu1robaWFW5UUqSkXYjbI2g9S
JZbiK6sVcp0hKqevcaeQ3515PN8fy2LYRSiQ0FUk3s7ZXWgd4sWlV6ACy3uJCQEH
2R34dvnwK6Z6ydAY7F/pYYMck2dZJeZSAvRDONJ/X4T+hD6ieuo1Dm+Sa8fu7uhI
ze8niFI5O1ZUT2JwTyqDIxCkz34klN3NJZ7mBk3SdMAJyjn/q2OJGIPSs9BzviJB
Wb0Lt+IOFe72t51qV2IPZoFICUx9w8xG2DLgPMlfIQKBgQD9+fgvCNobIaphN58t
CvR2uNkubvH5TVPwTm9e1m+tud0VchjZ+TQnFgAROIiXlQf+ASgtXwj8JoKEHRf7
rlGM0KxOtNuOJUag19lrobkTXyWcY6D83QnhMQmfrD2EolP1wKhBlC4uZEX6O75P
XemPRWNTaGSXitXtqAipVDecDwKBgQDMitOP6D1adA5xkyrKIxxPNKVX/Z3+2ZjZ
K6/MQnufhO6ExJPUntL3RsUuEB+MoQYDhVsKO102zj8ww1A7/eLNZqXXZ6cXuZRB
AtWdaLqNKn8c9Oepyp3MHF7Ki9RT6UDJFiQw7fJ+rJabQ5eMSVsZu8jGqQB4ASf6
x0yoOE8DSwKBgDohNITPadw74FtP98a/cySZOlw+WHPPFx8xVJxISFdi6w81hiqd
tyobBvjqD5liKuogKan4zn2n8sd6QTsBsvif6lA5ZOcr8PZvcPwJj+q3JEyfIQG/
NbutFZ7ONWZeIQlmhxw7ZjaIDNwxQGW6APMh1pIImr10sXIru4/GDtRRAoGBAMrK
mRaEWs/jYK6aCkJ+rQKaX5Ptj0es1S37ckBBae+uDAz2GNsk3GTtdXFF6wGyZBP5
k1mBodlEgsIF4vaXsNWUf7ggvDl5dNM/zCpUOyDakUxUQ5rDl89WEWUfXVQUXnb/
zMH2KAYPdwJY8VfVIs1QsK7FwAE6pDiugAIIUc4hAoGBAIJ7Kb2K0vW+thqpIUij
MX14SQ3Pnc1a0haLnF2HCpsN8oBMZnaqtHr7p2mjUV+IEXF6cWXO9yijZnQ4jaqD
Zk0U/azY0BakZt/Ej6CEDu89Nqsa42xU9vIaLfqntviPxM5axNwGqDck7T6uzstP
0eJtz5/p1nE+jzaAVqOyu89R
-----END PRIVATE KEY-----
`

const publicKeyPem = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyuzsaREAYPryMVPc5vO+
IPC8GrXHG1UaP7D/pXWSte3YpB84oRaUE3GOzXLfYk3TYpj+4CEuaBl+jnfdNjDH
3jP4SR90Rm/6XlSwVtlaYlTNcMP0e/UYEpVxELILAgSMHhVL0SOaa3qJyUSZs8vO
SfV5h8XwWszZEAmdZH8ltRc+QhqYs/sRNFChW3caKE2m1B4MX/fhMymH1RrIqU+p
H1XCYlBsGBzLyl+6mU72MXJIO89MawGLl7mr4wvPqpgvneXqOdvQ3Y14YlNyhF4w
0i/SUhk9XCBGGYXYAfvuDaC/vFk/8vB3VttNvomTfjeMUHtkBAba1j8LfYq8w7/l
ZQIDAQAB
-----END PUBLIC KEY-----`

func runAuth() (*auth.Auth, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatePem))
	if err != nil {
		return nil, fmt.Errorf("parsing auth private key %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyPem))
	if err != nil {
		return nil, fmt.Errorf("parsing auth public key %w", err)
	}

	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("constructing auth %w", err)
	}

	return a, nil
}
