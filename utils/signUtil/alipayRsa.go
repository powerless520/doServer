package signUtil

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func AlipaySignGetFromKeyFile(mReq map[string]string, priKeyPath string, if_rsa2 string) (sign string, err error) {
	var h hash.Hash
	var hType crypto.Hash
	switch if_rsa2 {
	case "yes", "":
		h = sha256.New()
		hType = crypto.SHA256
	case "no":
		h = sha1.New()
		hType = crypto.SHA1
	}

	_, err = os.Stat(priKeyPath)
	if err != nil && os.IsNotExist(err) {
		return "", errors.New("Alipay priKey file not exists!")
	}
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string
	for i, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			if i != (len(sorted_keys) - 1) {
				signStrings = signStrings + k + "=" + value + "&"
			} else {
				signStrings = signStrings + k + "=" + value
			}
		}
	}
	priKey, err := ioutil.ReadFile(priKeyPath)
	if err != nil {
		return "", errors.New("Error when read prikey file:" + err.Error() + "|" + priKeyPath)
	}
	block, _ := pem.Decode(priKey)
	if block == nil {
		return "", errors.New("rsaSign private_key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	h.Write([]byte(signStrings))
	digest := h.Sum(nil)
	signature_get, err := rsa.SignPKCS1v15(nil, privateKey, hType, digest)
	if err != nil {
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(signature_get)
	return data, nil
}
func AlipayNotifyRSAVerify(sortQueryBuildStr, sign string, pubKeyPath string) (pass bool, err error) {
	_, err = os.Stat(pubKeyPath)
	if err != nil && os.IsNotExist(err) {
		return false, errors.New("Alipay priKey file not exists!")
	}
	pubKey, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return false, errors.New("Error when read pubkey file:" + err.Error() + "|" + pubKeyPath)
	}
	block, _ := pem.Decode(pubKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, errors.New("FailedToParseRSAPublicKey:" + err.Error())
	}
	rsaPub, _ := pub.(*rsa.PublicKey)
	t := sha1.New()
	io.WriteString(t, sortQueryBuildStr)
	digest := t.Sum(nil)
	sign = strings.Replace(sign, " ", "", -1)
	data, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, errors.New("ErrorWhenDecodeReceivedSign:" + err.Error())
	}
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
	if err != nil {
		return false, errors.New("VerifySigError,Reason:" + err.Error())
	}
	return true, nil
}

func GenAlipaySignString(mapBody map[string]string) (sign string) {
	sorted_keys := make([]string, 0)
	for k, _ := range mapBody {
		if k == "sign" || k == "sign_type" {
			continue
		}
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string

	index := 0
	for _, k := range sorted_keys {
		value := mapBody[k]
		if value != "" {
			signStrings = signStrings + k + "=" + value
		}
		if index < len(sorted_keys)-1 {
			signStrings = signStrings + "&"
		}
		index++
	}
	return signStrings
}
