package test

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"testing"
	. "hellogo/rsas"
	"flag"
)
var cfg string

func init() {
	cfg = "c:/psm/"
	if flag.Lookup("test.v") != nil {
		//flag不为空,则说明存在test所拥有的参数,是在 go test 模式
		cfg = "c:/psm/"
	}
}
func Test_Crt(t *testing.T) {
	baseinfo := CertInformation{
		Country: []string{"CN"},
		Organization: []string{"WS"},
		IsCA: true,
		OrganizationalUnit: []string{"work-stacks"},
		EmailAddress: []string{"czxichen@163.com"},
		Locality: []string{"SuZhou"},
		Province: []string{"JiangSu"},
		CommonName: "Work-Stacks",
		CrtName: cfg+"test_root.crt",
		KeyName: cfg+"test_root.key",
	}
	err := CreateCRT(nil, nil, baseinfo)

	if err != nil {
		t.Log("Create crt error,Error info:", err)
		return
	}
	crtinfo := baseinfo
	crtinfo.IsCA = false
	crtinfo.CrtName = cfg+"test_server.crt"
	crtinfo.KeyName = cfg+"test_server.key"
	crtinfo.Names = []pkix.AttributeTypeAndValue{
		{
			asn1.ObjectIdentifier{2, 1, 3},
			"MAC_ADDR",
		},
	}
	//添加扩展字段用来做自定义使用
	crt, pri, err := Parse(baseinfo.CrtName, baseinfo.KeyName)
	if err != nil {
		t.Log("Parse crt error,Error info:", err)
		return
	}
	err = CreateCRT(crt, pri, crtinfo)
	if err != nil {
		t.Log("Create crt error,Error info:", err)
	}
	//os.Remove(baseinfo.CrtName)
	//os.Remove(baseinfo.KeyName)
	//os.Remove(crtinfo.CrtName)
	//os.Remove(crtinfo.KeyName)
}
