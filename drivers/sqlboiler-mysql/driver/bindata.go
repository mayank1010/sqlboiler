// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.tpl
// override/templates_test/singleton/mysql_main_test.tpl
// override/templates_test/singleton/mysql_suites_test.tpl
// override/templates_test/singleton/mysql_upsert.tpl
// override/templates_test/upsert.tpl
// DO NOT EDIT!

package driver

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templates17_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdf\x6f\xdb\x38\xf2\x7f\x96\xfe\x8a\x59\x63\xbf\xbb\xd2\x17\x2e\xdb\x03\x0e\xf7\x90\x83\x1f\x36\x3f\xda\xed\x6d\xd2\xb5\xe3\xf6\x0a\x5c\x51\x14\xb2\x34\x72\x88\xd0\xa4\x96\xa2\xd2\xfa\xb4\xfa\xdf\x0f\x33\x94\x6c\xc9\x76\x12\x77\xb3\x3d\xdc\x53\x22\x0e\x39\x33\xfc\xcc\xcc\x67\x86\xae\xeb\x67\xf0\xbd\x4b\x16\x0a\xdf\x24\x2b\x9c\x4b\xbd\xac\x54\x62\xe1\x64\x02\xe2\x2d\xad\x0a\x5a\x86\xdf\xa1\xec\x24\xbf\x83\x93\x4e\xe1\x59\x52\x22\x3c\x6b\x9a\x90\x15\xdc\x25\xf6\xe8\xe3\x69\xb2\x42\x35\x3c\x5e\xa6\x37\xb8\x4a\xf8\xc0\xfe\x51\x31\xdf\x4a\xf9\x80\xcc\x41\xfc\x94\x65\xaf\x94\x59\x24\x8a\x95\x3c\x7f\x0e\xef\x8a\x12\xad\x7b\x05\x89\x73\xb8\x2a\x5c\x09\x89\x06\xa9\x69\x6d\x0c\x89\xce\x20\x33\xc8\x6b\x55\x91\x25\x0e\xc1\x58\x90\x4b\x6d\x2c\x82\xd1\x90\x1a\x9d\x2b\x99\x3a\x11\xe6\x95\x4e\x21\x32\xf0\xff\x75\xbd\x0f\x4a\xd3\xc4\x9d\x99\x88\xbd\xd0\xc6\x81\x78\x63\xce\x8c\x76\xf8\xc5\x35\x4d\xea\xbe\x90\x2e\xfa\x10\xed\xe2\x18\xea\x1a\x75\x46\x5e\xb6\xa6\xcf\x8c\xaa\x56\xba\x84\x0f\x1f\x4b\x67\xa5\x5e\x8e\xe1\xf3\x8d\x74\xa8\x64\xe9\x40\x08\xe1\x17\x63\x40\x6b\x8d\x85\x3a\x0c\x2c\xba\xca\x6a\x30\xc2\x9b\xf6\x96\xfb\x56\x17\x46\x2a\xf1\x0a\xdd\xf9\x69\x14\xd7\x35\xaa\x12\xd9\x93\x31\x74\x82\x76\x67\x2b\xd7\x59\xd3\x8c\x87\xbe\xf4\x5c\x10\x42\xc4\x61\x13\x86\x1b\xaf\x43\x0f\x38\x41\xd8\x03\x9d\xfe\x9d\x26\x5a\xa6\x3b\xf0\x4f\x9f\x86\x3f\xb0\xce\x92\xd6\xf8\xfe\xc7\x07\x64\xfa\xed\x23\x52\x87\x81\xcc\xc9\x2f\x4a\xd1\xff\x62\x38\xfe\xce\x36\xbf\x9b\x80\x96\x8a\x9c\x08\x0a\x02\x29\x62\x75\xef\x6d\x52\x5c\x58\x1b\xa1\xb5\x71\x1c\x06\xcd\xa1\xd0\xdd\x13\xab\x43\xa1\x82\x8a\xea\x94\xbe\xf1\x0b\xa6\x95\x33\xf6\x6b\x8a\xa7\xa7\xba\xf8\x83\x71\x9c\xee\xe3\x49\x9e\x78\xec\x2e\x5a\x9f\x7a\xa8\xee\x07\x77\xbb\xbd\x5d\xea\x9d\x3a\x80\xf5\x13\x82\x7e\x20\xd7\xfa\xb9\x45\x8e\x7c\xb3\xc0\x6e\xa0\xfe\xd3\x83\x78\x64\xa0\xfe\xd7\xe2\xb4\xa1\x4b\x99\x83\x81\xc9\x16\xd3\x96\x3e\x59\x5e\x8a\x37\xf8\x39\x1a\xd5\xb5\x98\xde\x2e\xe9\x5e\x4d\x73\x02\xda\x40\x5d\xf7\x3a\x4e\xd3\x40\x61\xcd\x9d\xcc\x30\x83\xdc\x58\xa8\xf8\xbe\x23\x0e\x42\x18\x50\xbb\x22\xc0\x15\x41\x38\x72\x72\x85\xa5\x4b\x56\xc5\x27\xbf\xeb\xd3\x0d\xaa\x02\xed\x08\x04\x34\x7e\xf7\x36\x4f\x7e\x36\xe6\xb6\xe4\xe8\x0d\x32\x2a\x33\xa7\x98\x1b\x8b\x1e\x56\xde\x74\x74\x7a\xed\x67\xd0\xf6\xb6\xe4\x2e\x7b\xcb\x68\x86\x61\xa0\xff\x7d\x8e\x79\x52\x29\x57\x92\xe1\xdf\x2a\xb4\x12\x4b\xf1\xc6\xe8\x7f\xa1\x35\xad\x68\x8e\x14\xd8\xdd\x6e\xde\x34\x6d\x10\xde\x4b\x77\xd3\xee\x1c\x83\x89\xc3\x30\x78\xfe\x1c\x4e\x2b\xa9\x32\x48\x93\xf4\x06\xe1\x16\xd7\x20\xf5\x33\x25\x35\x42\xb5\x54\x52\xad\xe1\x19\xac\xd6\xe5\x6f\x0a\xee\x4a\x28\xe8\x6f\x61\xcd\x42\xe1\xaa\x0c\x83\x45\x95\x93\x27\xa5\xb3\xab\x44\x2f\x15\x12\x37\x9e\x56\x79\x8e\x36\x8a\xc3\x80\xa0\xff\x34\x86\x94\xb6\xd8\x44\x2f\x71\x27\x1f\xe8\xb6\x8b\x2a\x17\xef\xad\x74\x38\xe7\x2c\x88\x52\x0e\xd2\x76\xf9\x74\xed\x30\xfa\x51\xfc\x78\x50\xdd\x36\x8d\x9e\xac\xaa\x07\xed\x03\xba\x08\x9c\x93\x09\x90\xb4\x15\xc4\x61\xb0\xbd\xfd\xb4\xea\x6e\xbf\xa8\xf2\x98\x93\x67\x3f\x10\x3e\x4b\xce\x08\xec\xab\xca\x89\xeb\x4b\x93\xde\x92\x1a\x86\x7f\xec\xa3\x90\x91\x95\x47\x0e\x7f\xb8\xc5\xf5\xc7\xe3\x4c\xbc\xd3\xca\x1b\x09\x83\xbb\xc4\x72\xc2\x71\x31\x85\x9c\xc7\xdf\xb5\x26\xe9\xde\x5d\xbb\xb7\xe8\x86\x81\x7d\xcd\x02\x1f\x39\x4a\xb1\x30\x08\x0e\x9a\xee\x68\xf2\x11\x79\x3f\x0d\x8f\xd8\x6a\x2a\xd7\xdf\xbd\x0d\x16\x7f\x6e\xd2\x80\xbe\xe2\x30\x08\x5a\x7a\x1c\x5c\xe0\x5d\x2f\xf5\x9e\x72\x81\xa9\x95\xab\xc4\xae\x7f\xc1\x75\x7f\xe7\xb0\x47\xb0\x1b\x84\x66\x0e\x0a\x75\xe4\x85\x31\x51\xda\x0b\x46\xf9\x71\x46\xab\x34\x0f\xd2\xce\xb4\xdc\xb5\xcb\x6f\x63\x48\x4d\xa5\x32\x66\x98\x05\x57\x6f\x7b\xe7\x94\x5d\x00\x42\x83\xf8\x8e\x09\xcf\xa7\x96\x20\xbe\x58\xc3\xc4\xef\xf7\x29\x32\xa3\xa5\xab\xf5\x7c\x76\x19\x65\x32\x51\x98\xba\x31\x8c\x76\x4c\x8d\x3a\x1a\x1f\xb7\xcd\x29\xde\x68\xb4\xe8\x35\xc0\x04\xf2\x95\x13\xf3\xc2\x4a\xed\x72\x86\x76\x34\xbf\xb8\xbc\x38\x7b\x0b\xff\x57\xc2\xcb\xeb\x5f\xaf\xc8\xff\xcb\x59\xd3\xec\xe8\xae\x6b\x71\x3d\x6b\x1a\x78\xff\xf3\xc5\xf5\x05\xd4\xf5\xe7\x1b\xb4\x78\xa6\x92\xaa\x44\x10\x97\x33\x10\xd7\x33\x78\xd1\xbd\x25\xa6\xbf\xe0\x5a\xb4\x10\x93\x5b\x64\xc6\xf7\x8e\x52\xfc\xc3\x48\x1d\xf5\xb2\x35\x43\xed\x66\x95\x71\x38\x57\x32\xc5\xee\x72\xe2\x72\x36\x86\xee\xff\xeb\x19\xa7\x79\x3c\x86\xd1\x78\x14\x6f\x62\xe6\x6f\x76\x97\xa8\x0a\xaf\x92\xa2\xe0\x76\x45\x05\xb3\xe5\xdb\x53\xa9\xb3\x56\x74\x90\x69\xdf\xae\x0b\x1c\x1f\xac\xde\x8d\xc2\x2d\x90\x6d\x1f\xe9\xf1\xff\xa0\x01\x10\xe9\x74\x69\x44\xce\xd2\xc6\x36\x87\x36\x21\xf8\x76\x6e\x92\x45\x32\x75\xc0\xc9\xa1\x97\xec\x66\xe3\xdb\x2b\x43\xc7\xc4\x8a\x39\x03\xfd\x5a\x67\xd2\x62\xea\xa2\x6e\xe1\x9f\xb4\xe3\xd7\x3c\x32\x34\x16\xdd\x25\x6a\xd0\xcd\x58\x58\xbe\xb4\x66\xd5\x39\xcf\x0a\x5b\x62\x1c\x04\x26\xf6\x74\xe6\x3d\xa1\xf1\x42\x6a\x87\x36\x4f\x52\xac\x7d\x87\x26\xd4\x76\x61\xea\x41\xd8\x1d\xdc\x1a\x9f\x3a\x7b\xbf\xe9\x9e\x0e\x7f\x53\x99\xfb\xe9\xe7\x1c\x17\xd5\xf2\xca\x64\xc8\x5a\xa9\x12\x5e\x72\x25\x28\x1d\x6d\xe5\xdc\x4d\x6c\xa7\x8b\x6b\x31\x7e\x7c\x37\xa1\xb3\x19\x5b\xbe\x4f\x13\x7d\x99\x94\xce\x73\xf1\xeb\xf3\xde\x3b\xfb\x6c\x47\xc2\x33\x0a\xcf\x20\x7b\x87\x58\x14\xec\xcc\x7e\x7e\xd5\x62\xc9\x73\x41\x3b\xd9\xd0\x7c\xc2\x93\x60\xd4\x73\xda\xfb\xc4\x6f\x4c\xd2\x42\xe3\xe1\xc3\x87\x5b\x0b\x11\x0f\x3f\x0f\x28\x6a\xe7\xe2\x81\xce\xc3\x6e\x7e\xea\x52\xfd\xeb\x1c\xdc\x3f\xf6\xf5\xae\x75\xb3\xd8\x81\x92\x18\x72\x3a\x0d\xff\x34\xf9\x7b\x3e\x7d\x88\xd9\x69\x16\xd9\xa5\xdc\x4d\xc8\xef\x0d\x20\x25\xbe\xa2\xd5\x73\x90\xda\xfd\xed\xaf\x03\xe7\x48\x28\x89\x02\x65\x2e\xd1\x9e\x19\xb5\x5b\x1c\x0f\x54\xc7\xa6\x49\x2d\x8d\x33\xc0\x33\x44\x3b\xf8\x3e\xea\x93\xf7\xa7\x43\xd9\xe7\x83\xe8\x6d\xcb\x68\xcc\xb9\x17\xb8\x0b\x6b\xe7\x6b\x9d\xbe\x4c\xa4\xda\x26\xbc\x51\xfc\x0b\xd2\xc9\x04\xa4\xce\xf0\xcb\xa1\x56\x00\x2f\xb6\xc1\xa1\x03\xbd\xaa\xe0\x97\x3a\xf7\xc3\x8d\xa6\xc1\xd6\xb7\xd2\x29\x3f\x6b\x6d\xe4\xbd\x9f\xc6\x68\xa7\x11\xde\x0b\xbf\xb3\x69\x80\xc7\xb2\xd4\x28\x41\xf4\xd9\x34\x91\xbf\xb3\xbf\x57\x1b\x0f\xe6\x97\x1f\x7e\xb8\x1f\xdf\xbf\x90\x74\x57\xf2\xe1\xc5\x47\x92\x3d\xc0\xc7\x1f\x46\x43\x5f\x46\x1f\xef\x8f\xd3\xe0\xd9\xb0\x93\x0a\x93\x61\x32\xd4\x9c\xe3\x7e\x0a\x3e\x04\xaf\x2f\x44\x02\x42\x0c\xd1\x19\x0f\x8b\xe3\x89\x8c\xd8\xcd\x12\x47\x90\xe2\xf0\x3e\xbe\x4c\x37\xa9\xb9\x4b\x16\xbd\x92\x67\xfd\xd7\xe6\x73\x34\xb4\x78\x48\xa1\x98\xa7\x09\x77\x5c\x6a\x10\xde\x42\x9f\x4a\x0e\x28\x3d\xc0\x25\x5f\x6f\xa0\xc3\xf2\x4f\x60\x97\xc2\x14\x15\xbf\x70\x33\x3f\x2a\x03\xf7\xb2\xf2\x01\xba\xe9\xe7\xd0\xc9\xde\xcb\xe0\x88\x77\x46\xf7\x92\x79\x6c\x2f\xbf\x5c\x60\xe2\x51\x3a\x4e\xf5\xe6\x05\x13\x3c\xf0\x26\xdf\xfc\xcc\x9a\x99\x9f\x72\x87\xf6\x0f\xbd\xc7\xdb\xd2\xe9\xb5\x35\x56\xaa\x89\x95\xb6\x45\xd5\x84\xff\x09\x00\x00\xff\xff\x29\x2c\xd8\x2e\x75\x17\x00\x00")

func templates17_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertTpl,
		"templates/17_upsert.tpl",
	)
}

func templates17_upsertTpl() (*asset, error) {
	bytes, err := templates17_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.tpl", size: 6005, mode: os.FileMode(420), modTime: time.Unix(1527689483, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_main_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5b\x6f\xdb\x3a\x12\x7e\x96\x7e\xc5\xac\x81\x9c\x95\x72\x14\xa6\x40\x81\x7d\x48\x61\x04\x8d\xe3\x14\x41\x9b\x4b\xed\xec\x16\x8b\xa6\xd8\x32\xd2\x38\x21\x2a\x91\x0a\x49\xc5\xf1\x06\xf9\xef\x8b\x21\x75\x73\x62\x19\x2d\xd0\x7d\x3b\x4f\xad\x38\x1f\xe7\xf2\x0d\xe7\xe2\x3c\x70\x0d\xfa\xf6\xf1\x6c\x35\xff\xfc\xe9\x07\xae\x60\x0c\x1a\x6f\xf1\xb1\x64\x67\x95\xb1\x13\x55\x94\x22\xc7\xe8\x7b\x74\x58\xc4\x51\x94\x5c\xcb\xf8\xf0\xda\xfc\x39\xb9\x38\x9f\x5f\xcd\xde\x9f\x9e\x5f\xb1\xdd\xc3\x93\x8b\xd9\xf4\xf4\xc3\x39\x7c\x9c\xfe\x9b\xed\x1e\x5e\xcb\xf8\xcf\xef\x71\x18\xda\x55\x89\x50\xac\xcc\x7d\x7e\x85\xc6\xa2\x06\x63\x75\x95\x5a\x78\x0a\x83\xec\x66\xa2\xa4\x84\x5d\x73\x9f\xb3\xe3\xa3\x90\x0e\xce\x79\x81\x81\xb1\x5a\xc8\xdb\x30\xb8\x53\xc6\xb6\x1f\x95\x41\xdd\x7e\x94\xdc\x98\xf6\xc3\x98\xbc\x50\x59\x77\xad\x54\xda\x06\x42\xda\x30\x0c\x54\x69\x85\x92\x27\x22\x47\xa8\xa5\x61\x60\xd1\xd8\xe3\x23\x32\xd4\x9c\x3d\x87\xe1\xa2\x92\x29\x08\x29\x6c\x14\x7b\xcf\xce\xb8\x90\x30\x86\x3f\x7a\x9e\x3f\x3d\xb7\xc8\xa8\x80\xdd\x9e\x24\x06\x83\xb6\x2a\xa3\x18\x50\x6b\xa5\x49\x03\xb1\x89\x5a\xfb\x83\x30\x0c\x1e\x44\x89\x9a\xcd\xd1\x1e\xe3\x82\x57\xb9\x8d\x46\xee\x3e\xab\x9d\x1f\x25\x30\xb2\xba\xc2\x51\x3c\x0c\xa5\xb8\x46\x09\xbc\x7d\xfb\xe6\x1f\x71\x18\x06\x05\xf3\x74\xc1\x18\xfc\x8d\x0f\x68\xe7\x2e\xa0\xe6\x42\x76\x23\x79\xe1\x54\x16\x8c\xb8\x1c\x46\x92\xd4\xe3\x88\xe6\x61\x1c\x49\x3d\x8e\x32\x30\x8c\x23\x69\x8d\x53\x7a\xcd\xee\xa9\x5c\x8f\xc7\x81\x6a\x12\x86\xf5\x35\x2c\x51\xdc\xc4\xea\x18\x1e\x78\xce\xd9\x11\xde\x0a\xf9\x2f\x9e\x8b\x8c\x53\x9e\xa3\x98\xd5\x1f\x18\x85\x41\xe0\x20\x5e\xcf\xb9\xb2\xd3\xa2\xb4\xab\xc8\x07\x98\xc0\x5a\x3c\xc9\x20\x98\x78\x69\xc1\x9e\xa4\x16\x7c\xae\x6c\xe4\xfe\x33\xbd\xaf\x78\x6e\x22\x1f\x6b\x02\x6f\xda\x0b\x3e\xc0\x2d\xea\x7d\x02\x5b\x7c\x93\xaf\xe1\x0b\x35\x0f\xed\x8d\x96\x97\x24\x0c\x62\x36\xb9\xc3\xf4\x47\x44\x1c\x89\x85\x7b\x7c\x7f\x1b\x83\x14\x39\x3d\xc7\x40\xa3\xad\xb4\xa4\xd3\x30\x78\x0e\xc3\x60\x7f\x1f\x26\x1a\xb9\x45\xe0\xa0\xb9\xcc\x54\x21\xfe\x8b\x19\x64\x37\x40\x3e\x30\xca\x4a\xaf\x50\xc6\x1d\x86\xcd\x2d\xbf\xc9\xd1\x0b\xda\x18\x7a\x46\xc7\x50\xb0\x82\xff\xc0\x8b\xb6\xf6\xa2\xf8\xdd\xb0\x3b\x4a\x1b\xf6\x45\xf3\x32\x42\x4d\x79\x49\x55\x95\x67\xf2\xef\x16\x48\x05\xf8\xfa\x85\x85\xc8\xdd\x33\x7e\x5e\xb7\x92\x69\x55\x5e\x39\x27\xb7\x5a\xa0\x7b\xfd\x6b\xa9\x8b\xfb\x27\x2f\x86\x41\x56\x15\xe5\xa4\xc8\xe0\x60\x0c\xf8\x88\x29\x9b\xa8\xa2\xe0\x32\xab\x9f\x26\x49\x47\x09\x39\xe3\x8b\xd5\xf8\x80\x13\x18\xed\xed\x49\xb5\x97\x71\xcb\xbd\xb8\xa6\x29\xf0\xd6\x87\x15\x0e\x29\x23\x4d\x37\xdc\xa0\x93\x77\xa9\x21\xe2\x75\x02\x4b\xd2\x26\x14\xbb\x14\x25\x46\x71\xeb\x34\x9b\xdb\x4c\x55\x54\x80\xcb\x9e\x65\x3a\x75\xdd\x4d\xe2\xf2\xe4\x23\xae\x8e\xd1\x58\xad\x56\xa8\xa3\x5e\xff\x4f\x40\xaf\x25\xb5\xd3\xc8\xb5\xfd\x95\x8c\x2e\xb8\xc8\x31\x03\xab\xc0\xd0\x55\x68\x69\x83\xd4\x07\xee\x33\xdb\x59\xea\xfb\xf9\x3b\x6c\xad\xdb\xd9\x10\xd2\x17\x2e\x36\x59\x59\x14\x96\x5d\x6a\x21\x6d\x2e\x49\x7d\xfc\x73\x86\x97\x5c\x58\x58\x28\x3d\x14\x67\x18\x2c\xd9\x24\x57\x06\xa3\x18\xf6\xf7\xe1\xfd\x82\x66\x62\xf3\xc6\x84\x81\x4c\x49\x4c\x20\x25\x04\xd8\x3b\x84\xa5\x16\x16\x01\x65\x06\x6a\xe1\x0e\x4a\x51\x62\xb8\x91\xae\xff\x63\x1c\x2f\x63\xa8\x15\x48\x91\x6f\x99\x89\x26\x3f\x53\x19\x46\xae\xb5\xfb\x51\x1b\xd7\xff\x92\x5f\x66\x29\x6c\x7a\x07\x4e\xfa\x14\x06\x29\x37\x58\xcf\xc0\x83\xce\xc3\xd1\x6c\xfa\xf9\x9f\xa7\xb3\xe9\xf1\xa8\x41\x2c\x78\x6e\xd6\x21\xc7\xa7\xf3\xf7\x47\x9f\x1c\xa4\xae\x9b\xbe\xf4\x72\x36\x3d\x99\xce\xbc\x86\x2d\x03\x7c\xbd\xe2\x7a\x6e\xd6\x7a\x88\xc4\x79\x49\x2c\x2e\x22\xaa\xc6\x1a\xbe\x47\xad\x69\xbc\x63\x5c\x55\x76\xdb\x46\x3c\x6c\xe8\x65\x6b\xec\x56\x06\x5b\x94\x89\xcb\x9d\x2b\xe5\xca\x8a\x9c\x5d\x61\x51\x3a\xd8\x88\x16\x04\xaf\xbf\x69\x86\xdb\x7a\xfc\x60\x56\xfd\x63\xd9\xd8\x57\xcd\xd5\xe4\x92\x4c\x3b\x82\xc3\xe0\x3f\x49\xfd\xbe\x94\xa1\x3a\xb4\xf5\x34\xf4\x86\x95\x61\xa7\x86\xe6\xd2\xa3\x30\xd6\x3d\x2a\xe7\x80\xd7\x31\x06\xca\x62\x18\x3c\x03\xe6\x06\xe1\x17\xfc\x74\xcd\x1f\xa4\xb2\x54\xbe\x16\x8a\x76\x2f\x21\x07\x29\x03\x27\x65\xfd\x8e\x1d\x57\xa3\xaf\x69\x2e\x50\xda\x6f\x04\xe9\xc4\x8b\x5a\x4a\x97\xc7\x3b\xe6\x5a\xba\xe4\xd4\xce\xbf\x86\xd1\x94\x1e\xef\x64\x35\x8c\xbe\x36\xc2\x68\x55\xe8\xb4\xd1\xd7\x66\x6d\xdc\x98\xa5\xd2\x59\x07\xa5\x93\x8d\x50\x63\xf2\x3d\x7a\xfd\x1d\xb4\xad\x98\x66\xb0\xc7\x9e\x6e\xcf\x6b\x53\xcb\x2f\x48\x28\xb5\xb2\x2a\x55\xf9\xd8\xa6\xe5\x36\xae\xda\x7e\xf4\x17\x5d\x2f\xe9\xea\x97\x2e\x3d\xdf\xa2\x64\x6e\xbb\x89\xbb\x4e\x47\x67\x75\xdb\x1e\xae\xed\xf5\xcd\xa2\xab\x6c\x6a\xa0\x54\x59\xfd\x1e\x52\x57\x62\x33\xd6\x61\xc7\xbc\x7b\x35\xda\x1b\xe3\x05\xd3\x95\x9c\x14\x59\x64\xee\xf3\x66\xf7\x1b\x6d\xf1\xa3\xbf\x18\x6d\xf7\x82\x90\x9d\x0f\x54\xaa\x54\xd1\xe6\xb7\x7a\x63\x91\xeb\x4c\x2d\x65\xdf\x17\xb1\x70\x3b\x91\xfb\xfd\xd7\xeb\x0c\xcd\x59\x4b\x75\x7f\x60\x1f\xfc\xe2\xce\xd7\x3a\xac\x0c\x9b\x61\xa1\x1e\xe8\xa5\xfc\x54\x8f\x6e\xe2\xa3\x35\x29\x69\x46\x5f\x3d\x13\x12\xe0\xfa\xd6\x00\x63\xac\x19\x69\x6d\x50\x4e\x30\x06\x5e\x96\x28\xb3\xe8\xeb\x37\x0f\x78\x7a\xb9\xce\x3d\x7b\x15\x8c\x31\x7a\x5f\xe9\x86\x4d\xb0\xb6\xd8\xc3\x11\xac\x5d\xdb\xbc\x5e\xc3\xce\x71\x39\x43\x9e\xa1\xf6\x9e\x92\x36\xe3\xf7\xbd\x83\x31\xfc\x71\xb3\xb2\x68\xd8\x51\xb5\x58\xb8\x9f\xae\x24\xaa\x59\x7c\x25\x4a\xfb\x9b\xa2\x57\xd1\x1e\xfa\x21\xe0\x2f\xf7\x53\x41\xe2\x59\x25\x5f\x67\xa1\xbf\x6e\x34\x93\x47\x57\x52\x0a\x79\x7b\x30\x6a\xd9\xf4\xb1\xc5\xeb\x70\x6f\xba\xfe\xd9\x13\xc5\xaf\xa5\xa8\x75\x5f\x3a\x90\xef\xad\x5b\x49\xaa\x24\xbd\xc4\xa8\xfe\xab\x43\xe2\xd3\x17\x0f\x3f\xca\xf6\xd5\x7b\x49\xe2\xd4\x3b\x73\xeb\xbf\xf1\x83\x0e\x51\x73\x76\x9f\xb3\x8b\x12\x65\xb7\xd8\x67\x5a\x3c\xa0\x66\x6e\xc5\x3e\xaa\x44\x9e\x7d\xae\x50\xaf\xea\x80\x9a\x9f\xa6\xbe\x05\xae\x17\x5f\xd3\x91\x9b\x96\x5b\x77\xbf\x5e\xcf\x5b\xcf\x41\x47\x44\xf2\x8a\x9d\xf5\x40\x9e\xc3\xff\x05\x00\x00\xff\xff\xaa\xfc\xbe\x6b\xf8\x11\x00\x00")

func templates_testSingletonMysql_main_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_main_testTpl,
		"templates_test/singleton/mysql_main_test.tpl",
	)
}

func templates_testSingletonMysql_main_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_main_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_main_test.tpl", size: 4600, mode: os.FileMode(420), modTime: time.Unix(1527187183, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_suites_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\xb1\x8e\x83\x30\x10\x44\x7b\xbe\x62\x84\x28\xe0\x04\xfe\x80\x93\xae\xba\xea\xae\x48\x11\x91\x0f\x70\xc2\x82\x2c\x39\x1b\x84\x17\x29\x92\xf1\xbf\x47\x18\x8b\x90\xce\xe3\x99\xb7\x3b\xdb\xcf\x7c\x43\x4b\x4e\x2e\xa3\xa3\x49\x4a\xc1\x97\x90\x13\xc3\x83\x6a\x2b\xf8\x0c\xf0\xbe\xc1\xa4\x79\x20\x14\x86\x3b\x7a\xd6\x28\x44\x5f\x2d\xe1\xfb\x07\xaa\x5d\x5f\x2e\x84\x94\x33\x7d\x32\xd5\x9f\xfb\x7f\x18\x8e\x36\x9a\xdd\x27\xeb\x8e\x72\xcb\x9e\xf4\x3d\x0e\x4b\x64\x94\x0b\x46\x3b\x4f\xda\x62\x81\x18\xb1\xf4\xab\x77\x50\xd4\x79\xe6\x32\xf7\xfe\x4d\x87\x90\xd7\x58\x6b\x7f\x7e\x6e\x27\x55\x71\x19\x71\x77\xec\x91\x54\xc8\x5e\x01\x00\x00\xff\xff\x2f\xea\xf2\xb5\x00\x01\x00\x00")

func templates_testSingletonMysql_suites_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_suites_testTpl,
		"templates_test/singleton/mysql_suites_test.tpl",
	)
}

func templates_testSingletonMysql_suites_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_suites_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_suites_test.tpl", size: 256, mode: os.FileMode(420), modTime: time.Unix(1527187000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xdf\x8f\xd3\x30\x0c\xc7\x9f\x93\xbf\xc2\x54\x3a\x5d\x2b\x45\x3d\xee\x15\x69\x0f\x77\x6c\x9c\x0a\x63\xbf\x07\x42\x88\x87\x6c\x71\xb6\x48\x5d\x3a\x12\x67\x30\xa1\xfd\xef\x28\x5d\xb7\xf5\x8e\x21\xf1\xc0\x4b\xeb\xc4\xf6\x37\xfe\xd8\xbe\xbb\x83\x45\x30\xa5\x9a\x6f\x3d\x3a\x1a\x07\x74\xfb\x8f\xfb\xe9\xb8\x7f\xbc\xf5\x20\x21\x1e\x3c\x49\xc2\x0d\x5a\x02\x4f\xce\xd8\x15\x04\x1f\xbf\xb4\x46\x08\x75\x62\x57\x92\x84\xad\xab\x76\x46\xa1\xca\xb9\x0e\x76\x79\x5d\x37\x55\x46\x82\x72\x66\x87\xce\xe7\x5d\x23\x4b\x5c\x92\x00\x92\x8b\x12\x07\x72\x83\x8d\xbe\x80\xb0\x55\x92\x50\xc0\x8f\xb5\x21\x2c\x8d\x27\xf8\xfa\xed\xe8\xcb\x4e\x35\xfc\xe2\xec\xe2\xed\xc4\xdb\x8d\xb4\xab\x12\xf3\x42\xa1\xa5\x71\xa8\x08\xa7\xa5\x59\x62\x7c\x32\xef\x8f\x05\xc4\xff\x64\xdc\xd2\xcc\x38\x67\x8b\xa0\xe1\x4d\x3b\xfb\x09\xe9\x31\x68\x8d\x2e\xcd\x38\x53\xa8\xd1\xb5\x9c\xa3\x70\x72\x2e\x82\x8e\xe9\x3b\xe9\x60\x59\x95\x61\x63\x7d\x53\x17\x67\x46\x43\x89\x36\xbd\x3c\x03\xaf\x3a\xf0\x3a\xd6\xcb\x4e\xa1\x9d\x26\xd8\xe7\xef\x2b\xd3\x0a\x15\x90\x08\x48\x32\xce\x0e\xfc\xac\x73\x6c\x45\x06\x9d\x93\x88\xde\x50\xfe\x6e\xeb\x8c\x25\x9d\x72\xc6\x22\x82\x88\xff\xa4\x18\x4c\x7b\x93\x19\x14\x4f\x83\xe1\xa4\x07\xc5\x60\x36\x84\x1b\x0f\xe9\x8d\xcf\xe0\xd3\x43\x7f\xde\x9b\xd6\x76\x52\x07\x9f\x5b\x5e\x9f\x9a\xba\x6a\xbb\x45\x5b\xca\x25\xae\xab\x52\xa1\xf3\x75\x17\xe7\x1e\x0b\xab\xf0\x67\xdb\x21\x5e\xc0\x0a\xb8\x17\x70\x9f\x45\xa9\x8c\x33\xe6\x90\x82\xb3\xb0\x08\x3a\x9f\xd6\xc8\x69\x43\xf7\x82\xa2\x81\x38\x33\xfc\xa5\x78\x18\x0e\xa0\x3b\x1f\xf5\x8b\xb7\x0f\xb3\x1e\x7c\xe8\x7d\x81\xf9\xa8\x1b\xcd\x9a\xea\x19\x54\x8b\xe9\xbf\x21\xc5\x91\xeb\xca\x81\x11\xb0\x8b\x6b\xe3\xa4\x5d\x61\xb3\xac\xf5\x6c\x8c\x06\x73\x19\x77\xa4\xca\x3f\x3b\x43\xf8\xb8\x27\x4c\x6f\xc5\x6d\x6c\xc9\x81\x33\xf6\x3d\xae\xa7\x7a\xbe\x79\x97\xbd\xfd\x63\x65\x77\x19\x6f\x89\x35\x8d\x3c\x6a\x5c\xf3\x24\xd0\x69\x9a\x96\x26\xff\x98\x79\x2c\x30\xbb\x6d\xa6\x73\x6d\x6c\x07\xfe\x3b\x00\x00\xff\xff\xa9\x3a\x4a\xd3\x2e\x04\x00\x00")

func templates_testSingletonMysql_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_upsertTpl,
		"templates_test/singleton/mysql_upsert.tpl",
	)
}

func templates_testSingletonMysql_upsertTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_upsert.tpl", size: 1070, mode: os.FileMode(420), modTime: time.Unix(1527189392, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4d\x6f\xda\x40\x10\x3d\x7b\x7f\xc5\x14\xb5\x95\x5d\x39\x1b\xf5\x9a\x8a\x43\x42\x72\xa8\xaa\x22\x14\xcc\x0f\xd8\xd8\x63\xb2\x62\xd9\xb5\x76\xc7\x09\x74\xb3\xff\xbd\x5a\x1b\x12\xc2\x87\x84\x54\x55\x6a\x0f\x20\x31\xbc\xf7\xe6\xcd\xf3\x8c\xbd\xbf\x80\x8f\x24\x1e\x14\x8e\xc5\x12\xa7\x52\xcf\x5b\x25\x2c\x5c\x0d\x81\x17\xb1\xca\x63\x19\x5e\xc0\x6d\xff\x79\x01\x92\xa4\x70\x24\x1c\xc2\x45\x08\xec\xbd\xc0\x44\xb5\x56\xa8\x43\x7a\xd3\xd7\x8f\x92\x9f\x84\x3d\x8b\x5a\x8a\x25\xaa\xa3\xd4\xb3\x6c\xbf\xa7\xd7\xad\x2e\x81\xd0\x91\xf7\xfb\xee\x43\x98\x35\x0e\x2d\xa5\x04\x5f\x22\x42\xea\x39\x2f\x32\xf0\x2c\x21\x3e\x11\x56\x28\x85\x2a\xcd\x18\x4b\x64\x0d\x0a\x75\xea\xfd\xbe\x8f\x10\x46\x46\xb5\x4b\xed\x32\x18\x0e\x4f\x62\x26\x56\x2e\x85\x5d\xff\xc0\xf5\x2b\xda\xb3\x24\x21\x3e\x5d\xc8\x26\x1d\xc4\xef\x46\xea\x39\x74\xf6\xe0\x59\xd2\x23\x18\xad\xd6\xd0\xf4\x3c\x58\xe0\x1a\xca\x9e\x39\xc8\x58\x12\x18\x4b\x1c\x62\x15\x43\xb0\x42\x57\x66\x29\x7f\x21\x1f\xe3\xf3\x14\xb1\x4a\x33\x96\x3c\x09\x0b\x68\xbb\x8f\xb1\x2c\xb9\xbc\x84\x6b\x22\x5c\x36\x04\xf4\x88\xf0\x7d\x3c\xbd\xbb\x2f\xc0\xc9\x0a\xc1\xd4\x20\x34\xcc\x26\xb1\xc2\x12\x13\x15\x77\x63\x7a\x1b\xc1\x87\x2e\x85\xa8\xba\xdb\x74\x4a\xb6\x2d\x29\x8d\x6e\x72\xf8\x6c\x72\x38\x36\xfe\xed\x4d\xb1\x6e\xd0\xe5\x40\xb6\xc5\xec\x5b\x27\xf2\x61\x08\x5a\xaa\x4d\x0c\x77\xd1\x67\x9d\x0e\x66\xba\x0b\x80\xcc\x5b\x87\x13\x76\xc0\x75\x8d\xaf\xe0\x93\x1b\xe4\x51\x70\x13\x8b\xf7\xb2\x06\x6d\x08\xf8\xd8\x8c\x8c\x26\x5c\x51\x08\x25\xad\xe2\x60\x65\xff\x9b\xdf\x88\x72\x31\xb7\xa6\xd5\x55\x9a\x79\x8f\xba\x0a\x81\x25\x3d\xe4\x67\xeb\xa8\x58\xa5\x9d\xca\xae\xc2\x41\xe1\xc1\x48\xc5\x6f\x70\x2e\x75\xa7\xa1\x1c\xee\xd6\x8a\x55\x5a\xd2\x2a\x8f\x13\x6e\x3b\x9c\x05\xca\x58\x52\x61\x8d\x16\x68\xc5\xef\x8d\x52\x0f\xa2\x5c\xc4\xe7\xf9\x1a\xbc\xe1\x9b\x85\x3d\x35\x67\x7c\x00\xa8\xab\xb8\xf8\xb0\x15\x3f\x2b\xf0\xb6\xd3\x3d\x91\xf6\x41\xcc\xa5\x69\x35\x75\x85\xfd\x8d\xd9\x1e\x56\x9a\xf1\x51\x04\x9d\xe9\xf4\x6d\xc8\x43\x9f\xe9\xb6\x6f\x84\x74\x9d\x23\xe8\xeb\x3b\xc8\xe0\x59\x68\x02\xa3\x11\x2c\x96\xc6\x56\x39\xcc\x0d\x5d\x0d\xf2\x1e\xbf\x71\xbd\x77\x07\xb3\xc9\xed\x75\x71\x77\xec\x0e\xfe\x78\xd1\x6b\xa1\x1c\x1e\xc7\x1c\xbc\x0b\x38\xe7\x7f\xf7\x26\xfe\xe5\xed\xf9\x5f\x96\x27\xb0\xdf\x01\x00\x00\xff\xff\x59\x1c\xa8\x05\x3f\x07\x00\x00")

func templates_testUpsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertTpl,
		"templates_test/upsert.tpl",
	)
}

func templates_testUpsertTpl() (*asset, error) {
	bytes, err := templates_testUpsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.tpl", size: 1855, mode: os.FileMode(420), modTime: time.Unix(1527865928, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/17_upsert.tpl": templates17_upsertTpl,
	"templates_test/singleton/mysql_main_test.tpl": templates_testSingletonMysql_main_testTpl,
	"templates_test/singleton/mysql_suites_test.tpl": templates_testSingletonMysql_suites_testTpl,
	"templates_test/singleton/mysql_upsert.tpl": templates_testSingletonMysql_upsertTpl,
	"templates_test/upsert.tpl": templates_testUpsertTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.tpl": &bintree{templates17_upsertTpl, map[string]*bintree{}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_main_test.tpl": &bintree{templates_testSingletonMysql_main_testTpl, map[string]*bintree{}},
			"mysql_suites_test.tpl": &bintree{templates_testSingletonMysql_suites_testTpl, map[string]*bintree{}},
			"mysql_upsert.tpl": &bintree{templates_testSingletonMysql_upsertTpl, map[string]*bintree{}},
		}},
		"upsert.tpl": &bintree{templates_testUpsertTpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

