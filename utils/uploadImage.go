package utils

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// GetGCS : authenticate
func GetGCS(ctx context.Context, credPath string) (*storage.Client, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credPath))
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Upload : upload object to google storage
func Upload(ctx context.Context, r io.Reader, bucketName, credPath, projectID, filename string) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	client, err := GetGCS(ctx, credPath)
	if err != nil {
		return nil, nil, err
	}

	bh := client.Bucket(bucketName)
	// check if bucket exists
	if bkt, _ := bh.Attrs(ctx); bkt == nil {
		if err := bh.Create(ctx, projectID, nil); err != nil {
			return nil, nil, err
		}
	}

	obj := bh.Object(filename)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}

	// close file after writing
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	// Set object access to public
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return nil, nil, err
	}

	attr, err := obj.Attrs(ctx)
	return obj, attr, err
}

// Get object URL
func ObjectURL(attr *storage.ObjectAttrs) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", attr.Bucket, attr.Name)
}
