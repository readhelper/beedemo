package rsa

import (
	"runtime"
	"path/filepath"
	"github.com/astaxie/beego"
	"testing"
	"github.com/bmizerany/assert"
	"fmt"
	"os"
	"crypto/rand"
	"encoding/pem"
	"crypto/x509"
	"strings"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	os.Chdir(apppath)
	//apppath = filepath.Join(apppath, "conf", "app_win.conf")
	//beego.BeeApp.LoadAppConfig("ini",apppath)
	beego.TestBeegoInit(apppath)
	fmt.Println("------------------------------")
}

func TestNewCertificate(t *testing.T) {
	var certinfo = CertInformation{
		Country            :[]string{"cn"},
		Organization       :[]string{"huawei"},
		OrganizationalUnit :[]string{"paas"},
		EmailAddress       :[]string{"hao@huawei.com"},
		Province           :[]string{"shanxin"},
		Locality           :[]string{"xian"},
		CommonName         :"psm",
		CrtName:"psm.crt",
		KeyName  :"psm.pem",
		IsCA              :false,
		//Names              []pkix.AttributeTypeAndValue
	}
	err := CreateCRT(nil, nil, certinfo)
	assert.Equal(t, nil, err)
	beego.BeeLogger.Info("hello")
}

func encryptPEMBlock(key string, passphrase string) ([]byte, error) {
	p, _ := pem.Decode([]byte(key))
	_, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: p.Bytes,
	}

	block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(passphrase), x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(block), nil
}
func decryptPEMBlock(key string, passphrase string) ([]byte, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, fmt.Errorf("no valid private key found")
	}
	var (
		privateKeyBytes []byte
		err error
	)

	if x509.IsEncryptedPEMBlock(block) {
		beego.BeeLogger.Debug("x509.IsEncryptedPEMBlock(block)=%s", true)
		privateKeyBytes, err = x509.DecryptPEMBlock(block, []byte(passphrase))
		if err != nil {
			return nil, fmt.Errorf("couldn't decrypt private key")
		}
	} else {
		beego.BeeLogger.Debug("x509.IsEncryptedPEMBlock(block)=%s", false)
		privateKeyBytes = block.Bytes
	}
	beego.BeeLogger.Debug("block.Type=%s", block.Type)
	switch block.Type {
	case "RSA PRIVATE KEY":
		_, err = x509.ParsePKCS1PrivateKey(privateKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse der encoded key: %v", privateKeyBytes)
		}
	default:
		return nil, fmt.Errorf("unsupported key type %q", block.Type)
	}
	block = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	return pem.EncodeToMemory(block), nil
}

func TestEncryptPEMBlock(t *testing.T) {
	data, err := encryptPEMBlock(SERVER_KEY, SERVER_PWD)
	assert.Equal(t, nil, err)
	beego.BeeLogger.Info(string(data))
}
func TestDecryptPEMBlock(t *testing.T) {
	data, err := decryptPEMBlock(SERVER_KEY_ENCRYPT, SERVER_PWD)
	assert.Equal(t, nil, err)
	beego.BeeLogger.Info(string(data))
}
func TestAllPEMBlock(t *testing.T) {
	data, err := encryptPEMBlock(SERVER_KEY, SERVER_PWD)
	assert.Equal(t, nil, err)

	SERVER_KEY_ENCRYPT = string(data)
	data, err = decryptPEMBlock(SERVER_KEY_ENCRYPT, SERVER_PWD)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, strings.Compare(SERVER_KEY, string(data)))

}

var SERVER_PWD = "1234"
var SERVER_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAuLk8jHQk8DKcXu1NVMjXjGUtvlThn3ArLzwk1XeuQ52hmH1Z
4c5rdHgbGHM+SRGeGGWLfkOq4hvgVek+Q9KIwHxq2vfwYtvXkQJH4G/mhvTW2SB+
ylLl9EWAeA5lmgWcAl50Bxstm3LGJph4sNGFNs5lvV5Gh0I9VXcliwlu66sCPQ0F
EYvRmrRPRvBYghBSdd5w1yBfhkSz9LjnylbE6JSlXCX2occjUZcGJVj9uF/CDMvU
jzeWuJmY/w457RMcW7Q9TdpDaJJ9bCvZAzrBAZA0P7RaAlrM6UWlqKY6a6SxpkGQ
j9ZHcLHGyYLT2sw1E5rWdv4Jxf3riheRv2vyowIDAQABAoIBAB43LpBK7z/bB3j/
mAamU5vDcRgRClbqIiY30E6apQaqYiRvXKRy/2CtxMDbGPUazKFw+sBMkUcrCCcF
YAn4BiZ2M4RdyhRhoYE2vksYAr1Qj1Q03Ih7xuGN/NWmhTgMcthwWspOx8cFnyPx
DMzfeMreOAYhxaeaTi4Mrzdu85XYu7Rg9fYezOkeI7s4iUaYIDmfzrwCfDVv9Ami
xEK4qCWa/p7f8Sq/fjy02jD6WtAIfmDQ8GoVTMhkopwoJMfPr1AXs4O8odpzgcJz
jnmxNs0AFrwYuf/2InaojiAYKYnO5xWaxVCgrzAMKrERFkLu9PpXT2l0oDRBFV+5
yN6YXgECgYEAxnLMLkeTChaZH+jzLjXAZNSNI6nCvDdAdo3lvMsueGw1FSuKHHYu
tDygGACtZaZEPsq6jWntd2C+e07zW2YSnkYOFLtYzPmx6hb1vkxmcC565M4C5HXW
Du3hRnEuzzY/GWUMV14PGOtOzxpOAhQkoVUpn/EDvnBy24W9/CkGS6sCgYEA7kt6
q3GYa6Ad71WsDgtS0nXCM0uFj6IFiryh7bs4FkVHAjLvh2I37spiIffbdKug5mEI
deEkIbDZUylaRhMfRYpmpT4Pd2kElRncSoqDrtx4kH2NA9xfPUteJO64JNtne99Z
QttIVRg4TJMIUnG0wuameP0X3dgv5TQhmA/fPOkCgYAfyELrJotmEjhz49sOi41E
mMYB8C4/9plcY11n3yKSJsJZoZ9873CLbSo2reUXYomLhOxbuZtOgy/+j6Zp/O7+
ajnXGCogqdzUoNi6oYHGdas4+cV80WJ3AaISpg3ZIdb5OjW/yFCirQzyHkRgENnM
/NatxuvSlC/Q+Jp0yB06AQKBgQC5QYjenHOIyHQ9Svd1+6up3s7Zng75uVZho4Co
F3LlLXqk5QB+2gbt5/0et1ivfXabZDh1nFDAuIJcOrvp2pfnMk/Wi0bu6IAdgntW
oBT77n70pgbpR3vrZRqNz+DtFn1/OgZlCaFUNn4eWfiQT8Hd+/7T9L+HQrkJkdvE
IiLFcQKBgQC6bjN1MHvpLHpCDDzIFk2yEbg7YbR2HrRbj7J5p3J34FWdzrvha9t6
xR8na2j7zlHKYWhrCw13l73YaM3S/kosyBDBi/17xz1zP2j3dqsXFxUGUBqYF3OH
ygf6Uor/Bvw/dk+vDOSASmkR4YCdHn7821vWinzkVQzCqV4thzp3rg==
-----END RSA PRIVATE KEY-----
`

var SERVER_KEY_ENCRYPT = `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,ef29f583ae6d94c4b63933a85c53596a

mihMuTbUYWnOrW8d77dPAma5R9oIPnfQ5sQEkNrOST32cY/wP6B30fNPUD3SehAb
9f26pexqIKbwp9hIBL5Ak7tO+b3/FqwgExkOapTpzwb5SEsJ2kMHV39J0qB4nWVX
kW7YClaIl8Thp4vJZmPg8wQZvVdpU5jby+zFcXEht/R+zetUNLA+LO3HeMg9xMil
XlQhkAbLjkh4a8eNxKQ7jOgVFwIC6aWmhCMR5CNEQvb/z8uWCJ0TsvvuQhBMbYsx
hVskmal4+FaONr+hLFD7YbA40L2gx+k+7FO4C4UWRQsKYF90RqBcI8h3x+X79fWy
FwPBUPi9gOWTipKD9KG4WlXjxw5ZhR3EP2LARNW3iSp2Edj/A5FcL4i2Bj8UfuMO
M35FGw2az1olx3z55r4F+QULAevlcHZTj8pWaDDc/1PtFAueuMg1pbRqiTj6tpgp
LEaSkiwzdRPCxs5nDMh7kG42oz3EFB6EMPv0inRtPRhawSkpH4kRlNmDcqSVNg6C
aSTsdjdYIuhNwaKL13ndS3zoIUS3CzZWDMoHK9Zsmvxf7ZVvjiP1hkeYv1udrDx9
YeddnfLPiFE6dg7l0fUYiCRMiR4ciSI8Gj6gWe3b/C8lZWAk/SUXFZkdJfxF/Z6D
+P3g3+xtr9tt/cUNtx0qMFefD6E7Za6wMG7eirxT/pynpcDygBgEMY8dZChLdVq7
vugPjs7KIft86QfdaYz1aiABdxry7MGJ/aqIJzlwGn1qlcPkAd6rtqAG5baznCGI
o6FIwJMVQ8NL+eZGxxRPXH1dSSMwoyPSEySxHdK4pMh8O500CTe3chTBF1tKeNbz
gTsJKJQ1XNtitrGNSY3EDbn1W+bXrTdAwtH5mySArtuSDMQDIKzhBIuHOgLWynZH
xS8jHa2EDjrUw2eO3HBBtdz0fCpuaHihlF1cgxf5Awe35rdRY+pZO771E0j2UuMc
QG8Q8etlIkUXFRRwAjTXqtruAp03wuJn4TFG/q520bfMpwb4N7xdn8vstYmUkrGK
p9RUF/CHayeRCIH8qeUeuUc/GSl6lQQ8CLk+y1I+gnMKXVspqj6h8xknPdB1CXav
xr3l53eIEf4xnrJqwU+EOWb5DWOGElaOpHrp2qgqLLc5eHyABletHZF/3XCUA19+
R2XsAfzaJLqHG5E3nIJvjxXOvVf0bSUzvC340iiqaNWF9c44uJLjdZdUUoxOau22
0pIFev7XgzmqNF/utAehg5c0dBpCjCGxI6rQtY6semfcuJLWyOacZYRP4dmFqgyS
pcdBLl+VzTRS9YCmr3HmCYWiYoVgiUeQ2i7q529z0YpsGhQD6a/pDqqk8Riubzt3
hSAFZCWsuuvTxQHKypPZzMc2Xulk42t5VTR87AlphvwFLe+T3ouDS9LaSyyydoGq
Aut63FEcCmh9i8y8ETLtin+aFANpEDIIxJb0Acl6mCfS2oY69pXeFWy7734gkLl+
7arHSOzH7EeZ0X/VccQcYjT/41I33h1uKfsX6BlEWIDhLk8F3qsk2elm7oxNOQFf
YisbUryJF/BYSayTzAy8TaJPgndG+ZTTi+FhsvzuOed5Iy6CPkUqWWQqqiPeoG3U
-----END RSA PRIVATE KEY-----`