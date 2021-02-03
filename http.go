package dingtalk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/xiaojiaoyu100/cast"
)

type postResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (g *Group) post(cast *cast.Cast, message interface{}) error {
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, g.secretKey)
	sign := hmacSha256(stringToSign, g.secretKey)

	var param struct {
		Timestamp int64  `url:"timestamp"`
		Sign      string `url:"sign"`
	}
	param.Timestamp = timestamp
	param.Sign = sign

	request := cast.NewRequest().
		WithQueryParam(&param).WithJSONBody(message).
		WithTimeout(5 * time.Second).
		Post()
	response, err := cast.Do(context.TODO(), request)
	if err != nil {
		return castRequestFailed
	}

	// 尝试序列化
	var postResp postResp
	err = json.Unmarshal(response.Body(), &postResp)
	if err == nil && postResp.ErrCode != 0 {
		// 如果序列化成功，识别错误码
		return errors.New(postResp.ErrMsg)
	}

	// 如果序列化失败，说明成功
	return nil
}

func hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
