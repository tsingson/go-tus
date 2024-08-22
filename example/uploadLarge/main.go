package main

import (
	"fmt"
	"os"

	"github.com/eventials/go-tus"
	"github.com/eventials/go-tus/memorystore"
)

const (
	photoUriExporter   = "X-PhotoUri-Header"
	fileInfoIdExporter = "X-FileInfoId-Header"
	metaTypeExporter   = "X-MetaType-Header"
	metaIdExporter     = "X-MetaID-Header"
	ownerIdExporter    = "X-OwnerId-Header"
	ownerIdMeta        = "ownerid"
	metaTypeMeta       = "metatype"
	// MetaIdMeta   = "metaid"
	fileTypeMeta = "filetype"
	fileNameMeta = "filename"
)

func main() {
	fp := "/Users/qinshen/go/src/github.com/tsingson/ultrastream/vodfront/3.mp4"

	f, err := os.Open(fp)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		os.Exit(1)
	}

	url := "http://127.0.0.1:8083/files/"

	s, _ := memorystore.NewMemoryStore()

	cfg := &tus.Config{
		ChunkSize:           1 * 1024 * 1024,
		Resume:              true,
		OverridePatchMethod: false,
		Store:               s,
		Header: map[string][]string{
			"X-Extra-Header": {"somevalue"},
		},
	}

	client, err := tus.NewClient(url, cfg)
	if err != nil {
		fmt.Println("new client error:", err)
		os.Exit(1)
	}

	upload, err := tus.NewUploadFromFile(f)
	if err != nil {
		fmt.Println("new upload error:", err)
		os.Exit(1)
	}
	upload.Metadata[metaTypeMeta] = "0"
	upload.Metadata[fileTypeMeta] = "1"
	upload.Metadata[ownerIdMeta] = "0"
	upload.Metadata[fileNameMeta] = "3.mp4"
	upload.Metadata[fileTypeMeta] = "video/mp4"

	uploader, err := client.CreateUpload(upload)
	if err != nil {
		fmt.Println("create upload error:", err)

		os.Exit(1)
	}

	err = uploader.Upload()
	if err != nil {
		fmt.Println("upload error:", err)
		os.Exit(1)
	}
}
