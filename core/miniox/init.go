package miniox

import (
	"context"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func (c *Cfg) NewClient() *Client {

	conn, err := minio.New(c.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Username, c.Password, ""),
		Secure: c.TLS,
	})

	if err != nil {
		log.Fatalln(err)
	}
	if c.Bucket == "" {
		panic("no bucket name")
	}

	timeout := 10 * time.Second
	if c.Timeout != 0 {
		timeout = time.Duration(c.Timeout) * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err = conn.MakeBucket(ctx, c.Bucket, minio.MakeBucketOptions{Region: c.Location})
	if err != nil {
		// 检查我们是否已经拥有这个桶
		_, err = conn.BucketExists(ctx, c.Bucket)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return &Client{
		conn:     conn,
		Bucket:   c.Bucket,
		Location: c.Location,
		Timeout:  timeout,
	}
}
