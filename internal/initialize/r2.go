package initialize

import (
	"auth_service/global"

	"github.com/thanvuc/go-core-lib/storage"
)

func initR2() {
	r2Configs := global.Config.R2
	r2Clients, err := storage.NewClient(storage.Config{
		AccountID: r2Configs.AccountID,
		Endpoint:  r2Configs.Endpoint,
		AccessKey: r2Configs.AccessKeyID,
		SecretKey: r2Configs.SecrecAccessKey,
		Bucket:    r2Configs.BucketName,
		UseSSL:    r2Configs.UseSSL,
		PublicURL: r2Configs.PublicURL,
	})
	if err != nil {
		global.Logger.Error("Failed to initialize R2 client", "")
		panic(err)
	}
	global.R2Client = r2Clients
}
