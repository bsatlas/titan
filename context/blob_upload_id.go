package context

import (
	"context"
)

// BlobUploadIDKey is a context upload id for OCI blobs.
var BlobUploadIDKey Key = "blob-upload-id-key"

// BlobUploadIDFromContext returns a blob upload id defined in a context.Context.
func BlobUploadIDFromContext(ctx context.Context) string {
	uploadID := ctx.Value(BlobUploadIDKey).(string)
	return uploadID
}

// WithBlobUploadID returns a new context with the provided blob upload id.
func WithBlobUploadID(ctx context.Context, uploadID string) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, BlobUploadIDKey, uploadID)
}
