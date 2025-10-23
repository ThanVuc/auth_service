package constant

import (
	"time"

	"github.com/thanvuc/go-core-lib/storage"
)

func UserAvatar() storage.PresignOptions {
	return storage.PresignOptions{
		KeyPrefix:   "user_avatars/",
		ContentType: "image/webp",
		Expiry:      2 * time.Minute, // 15 minutes
	}
}

func UpdateUserAvatar(objectKey *string) storage.PresignOptions {
	return storage.PresignOptions{
		ContentType: "image/webp",
		Expiry:      2 * time.Minute, // 15 minutes
		ObjectKey:   objectKey,
	}
}
