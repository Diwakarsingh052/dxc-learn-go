package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"graphql-service/auth"
)

const privatePEM = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDCZ68KZEihchk4
TmCbaiUYd5sJtiN3kRGvhxpM0x9/bjrmtGczlPvn71C0JukpduQOxWmh9lUUmlNx
qTXLyZShpf3SdpZTQlT/PTbuC6R18gZmf0gZ+/OT5dS+9A+c6ihpcll1QGI/k4I2
7KTwa4fe0XY0mnlPhhlASwEewxIsoCcQcheunNpMgkpgOVsX4GFIYDPLedmvVXSO
qmZbn+m4Svg2B1mYaAobpR64Fx/j0zuo1u567PbOgJAmRRtCTMopaulyUUCtDfTP
FEe2GOJrIz0u5/YP7eLtwTQYlghbtBEoFz2BJgqATJkFvBnp8wdiu2EzL9pt2p+3
mK5DzFylAgMBAAECggEAaSUXATHYLqm9hjyt96DjS2Z1Tj1a1XJ69ErIKMEPiiae
aOUt5DIyRPGk3qsk9K4/DtYrGdRXk/scIp94Xy4522wh6lEFYSbBPgNn0DwvyTML
zygMYTFqPpXSCS2LvDzReqbER6x49xXbGcXIN9iQ9iyoGC9saVyID8QBaRfsnoFW
yocKqsx8vF6QCW1yk6TBXhWit0l76oZrJBXn3RsaM1KCy6G3q8bWv1nByoD/uF0U
P15L4phGjBylsuIfAgueBvPFgW9v4WT3rCN/5GrNUjyCdBiBPKZvT5mXl9WlrYFi
IDQgJPUtMmeGWOdFwdj4QwIg5OxMPbN0HmuJtZY9AQKBgQDpwTwvRSwkIhOWBhYZ
U+0814S4hzAvLgxMT7VBBl4g3/nI8AlrY7IQ9GtjwtlSd0rCebRMgPgExGI3UuXk
aOMg3TRmEWDbjNXINlvw4j0ZelgtfC+rkFk8AmZ85Eo7YcIntb/D4sHXJ065/c2L
/SejEX6QHPOQIMeKmrtbr6jKRQKBgQDU5822BCHy9NziX1O32xg+NCsx6S9PgSr0
xslDh3k6QpvvI4YLzmzz1hZKKK/iL3EGiOsdCYMsGoiHqHzMIc3ZD8OeR7KbOp4s
cbJuX2qV0GGGV0d79mOc+zU5OTLB503Ly4jjIMW+Q3MgVN21E3FYEYyFsYFeslyT
0iVFDOye4QKBgQDGZ5d339SJjsrlCaF0OgIhJCSoo0ZIpWMW2ojT+l2mtbTD0smL
9wDK10rXUCk5j7tKuhZI4SailTVwE9LBPi2aVWcGQzXc4/sIhuse2EhX+boiUUf8
7PZwsvDejL5dDKrJHcD4uN0ii2CNCVmPun4MMOsl0w0AhnIXuSHRvpUbqQKBgBlh
rSeG5Jk3UeH25V8k0oYP6BpiJ06+ImXeEP5o9y3X5QkkXgWoTVrgafXbVeSMLVhP
GGB00tt+Kkqp7n7ThNvcwkBrYcKZwWOhBlmcLHPBzO6cFxyTKhr748N3qzJspdym
3iHdtVVSazYuh+PfdoK+TNdfawHkF59TzTenK8phAoGAanrmCaqqOMdAzfKXNhNF
SD3m3Gm/cYubC/cA2ZWsLvJdJpSfP/fKdAk1c6Ykw3uw5lgqKaAd4VKPe96g/Fkk
2RcJMYnOzuhXTisqKcQRq636Pg1rMG74tbVmO7CxhGNhimbLJIMD5qt1JGf08n+q
qHiBJszclVssmZjYeLe0VDI=
-----END PRIVATE KEY-----`

func runAuth() (*auth.Auth, error) {

	//ParseRSAPrivateKeyFromPEM parses a PEM encoded PKCS1 or PKCS8 private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatePEM))
	if err != nil {

		return nil, fmt.Errorf("parsing auth private key %w", err)
	}

	//creating instance of the Auth struct from auth package
	a, err := auth.NewAuth(privateKey)

	if err != nil {

		return nil, fmt.Errorf("constructing auth %w", err)
	}
	return a, nil
}
