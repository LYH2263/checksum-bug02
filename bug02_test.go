package checksum

import (
	"context"
	"testing"
)

func TestBug02_ServeDigestHonorsCancel(t *testing.T) {
	p := New(Options{ChunkSize: 4096})
	body := make([]byte, 256*1024)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := ServeDigest(ctx, p, body)
	if err == nil {
		t.Fatal("expected cancel error")
	}
}
