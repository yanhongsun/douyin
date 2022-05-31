package pack

import (
	"douyin/cmd/comment/pack/configdata"
	"encoding/base64"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tms/v20201229"
)

func CommentModeration(content string) (string, error) {
	var res string
	credential := common.NewCredential(
		configdata.TencentCloudConfig.SecretId,
		configdata.TencentCloudConfig.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tms.tencentcloudapi.com"
	client, _ := tms.NewClient(credential, "ap-beijing", cpf)

	request := tms.NewTextModerationRequest()

	request.Content = common.StringPtr(base64.StdEncoding.EncodeToString([]byte(content)))

	response, err := client.TextModeration(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return res, err
	}
	if err != nil {
		panic(err)
	}

	res = *response.Response.Label

	return res, nil
}
