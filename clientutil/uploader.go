package clientutil

import (
	gignore "github.com/sabhiram/go-git-ignore"
	"github.com/tsileo/blobstash/client/interface"
)

var (
	uploader    = 25 // concurrent upload uploaders
	dirUploader = 12 // concurrent directory uploaders
)

type Uploader struct {
	bs  client.BlobStorer
	kvs client.KvStorer

	uploader    chan struct{}
	dirUploader chan struct{}

	Ignorer *gignore.GitIgnore
	Root    string
}

func NewUploader(bs client.BlobStorer, kvs client.KvStorer) *Uploader {
	return &Uploader{
		bs:          bs,
		kvs:         kvs,
		uploader:    make(chan struct{}, uploader),
		dirUploader: make(chan struct{}, dirUploader),
	}
}

// Block until the client can start the upload, thus limiting the number of file descriptor used.
func (up *Uploader) StartUpload() {
	up.uploader <- struct{}{}
}

// Read from the channel to let another upload start
func (up *Uploader) UploadDone() {
	select {
	case <-up.uploader:
	default:
		panic("No upload to wait for")
	}
}

// Block until the client can start the upload, thus limiting the number of file descriptor used.
func (up *Uploader) StartDirUpload() {
	up.dirUploader <- struct{}{}
}

// Read from the channel to let another upload start
func (up *Uploader) DirUploadDone() {
	select {
	case <-up.dirUploader:
	default:
		panic("No upload to wait for")
	}
}
