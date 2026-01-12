package fs

import (
	"io/fs"
	"os"
	"time"

	"github.com/ancientlore/cachefs"
	"github.com/golang/groupcache"
)

func NewFileSystem(dirPath string) *fs.FS {
	groupcache.RegisterPeerPicker(func() groupcache.PeerPicker { return groupcache.NoPeers{} })

	cachedFS := cachefs.New(os.DirFS(dirPath), &cachefs.Config{
		GroupName:   "groupName",
		SizeInBytes: 10 * 1024 * 1024,
		Duration:    10 * time.Second,
	})

	return &cachedFS
}
