package signUtil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/juju/errors"
	"hash"
	"net/url"
	"sort"
)

func ALipayNotifyRsaSignCreate(params_urlvalues url.Values, rsa_key string, if_rsa2 string) (if_pass bool, err error) {
	var h hash.Hash
	var hType crypto.Hash
	sign_get := params_urlvalues.Get("sign")
	switch if_rsa2 {
	case "yes", "":
		h = sha256.New()
		hType = crypto.SHA256
	case "no":
		h = sha1.New()
		hType = crypto.SHA1
	}

	sign_submit_str := ""
	var params_keys []string
	for k, _ := range params_urlvalues {
		if k == "sign" || k == "sign_type" {
			continue
		}
		params_keys = append(params_keys, k)
	}
	sort.Strings(params_keys)
	for _, k := range params_keys {
		if params_urlvalues.Get(k) != "" {
			sign_submit_str += "&" + k + "=" + params_urlvalues.Get(k)
		}
	}
	sign_submit_str = sign_submit_str[1:]

	block2, _ := pem.Decode([]byte(rsa_key))
	if block2 == nil {
		return false, errors.New("PublicKeyFormatError")
	}
	public_obj, err := x509.ParsePKIXPublicKey(block2.Bytes)
	if err != nil {
		return false, errors.New("CanNotRestorePublicKey:" + err.Error())
	}
	rsaPub, _ := public_obj.(*rsa.PublicKey)
	data, _ := base64.StdEncoding.DecodeString(sign_get)

	_, err = h.Write([]byte(sign_submit_str))
	if err != nil {
		return false, err
	}
	hashaa := h.Sum(nil)

	err = rsa.VerifyPKCS1v15(rsaPub, hType, hashaa[:], data)
	if err != nil {
		return false, err
	}
	return true, nil
}
func AlipayRsaSignCreate(params_urlvalues url.Values, rsa_key string, if_rsa2 string) (sign_new string, err error) {
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

	sign_submit_str := ""
	var params_keys []string
	for k, _ := range params_urlvalues {
		params_keys = append(params_keys, k)
	}
	sort.Strings(params_keys)
	for i, k := range params_keys {
		if params_urlvalues.Get(k) != "" && k != "sign" {
			if i == 0 {
				sign_submit_str = k + "=" + params_urlvalues.Get(k)
			} else {
				sign_submit_str += "&" + k + "=" + params_urlvalues.Get(k)
			}
		}
	}
	block2, _ := pem.Decode([]byte(rsa_key))
	if block2 == nil {
		return "", errors.New("PrivateKeyFormatError")
	}
	priv_obj, err := x509.ParsePKCS1PrivateKey(block2.Bytes)
	if err != nil {
		return "", errors.New("CanNotRestorePrivateKey:" + err.Error())
	}
	h.Write([]byte(sign_submit_str))
	hashed := h.Sum(nil)
	signature2, err := rsa.SignPKCS1v15(rand.Reader, priv_obj, hType, hashed)
	if err != nil {
		return "", errors.New("SignCreateError:" + err.Error())
	}
	sign_new = base64.StdEncoding.EncodeToString(signature2)
	return sign_new, nil
}
