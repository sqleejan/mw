package chaotic

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _app_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\x4d\x4e\xc4\x30\x0c\x85\xf7\x3d\x85\x25\xb6\x34\x2a\x7f\x9b\xcc\x69\x92\xda\x4d\x2d\x52\xbb\x4a\x3c\x9a\x01\xc4\xdd\x49\xa1\xa0\x01\xa4\xc9\x2a\x7a\x79\xef\xd3\xa7\xb8\x99\x02\x52\x81\xb7\x0e\xda\x31\x3a\x5b\x1f\x32\x27\xf1\x30\x92\x18\x95\xc3\x67\xbe\x06\x44\x96\xd4\x47\x35\xd3\xc5\xc3\x3d\x2d\x87\xee\xbd\xeb\x26\x2d\x0b\xb8\x78\x6c\xa9\xd4\x2b\x8c\xef\x6a\x0e\x91\xf2\x2d\x7c\xdd\x29\x91\xe0\x3e\x9a\x54\xac\x3f\x11\xa7\xd9\x3c\x3c\x0c\xc3\xc5\xe4\x5f\xad\xf2\x2b\x35\x05\xf7\x58\x36\x8b\x2d\x8e\x61\x7c\x4e\x45\x8f\x82\xfd\xa8\x59\x8b\x87\xd3\xcc\x46\xbf\xdc\x3d\x0c\x30\xb8\xa7\x4b\xf1\x89\x29\x63\x25\xdb\xe1\x51\x4b\xfb\x09\x0f\x77\xeb\x19\xaa\x66\x46\xb8\x41\xc4\xbf\x90\x1f\x84\x9b\x19\x91\x64\x1f\x23\xd7\x35\x87\x17\x0f\xa2\x42\xdb\xfb\x47\x00\x00\x00\xff\xff\x01\x1f\x45\x6e\x59\x01\x00\x00")

func app_css_bytes() ([]byte, error) {
	return bindata_read(
		_app_css,
		"app.css",
	)
}

func app_css() (*asset, error) {
	bytes, err := app_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "app.css", size: 345, mode: os.FileMode(420), modTime: time.Unix(1422683328, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _app_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x3d\x6f\xdb\x30\x10\x9d\xa3\x5f\x41\xb8\x01\x48\x21\x82\xb2\xc7\xc8\x50\x34\xf5\xd0\xa1\x35\x90\xd1\xf0\xc0\x58\xe7\x88\x00\x4d\x0a\xe4\x49\x85\x5b\xf8\xbf\xf7\x8e\xfa\xb0\x2c\x64\x08\x9a\x21\xa2\x4f\xf7\xee\x3d\xbe\x77\x76\xa7\x83\xf8\x56\x6b\x8f\xe6\x20\x9e\xc5\xdf\x4c\xd0\x1f\x38\xfd\x66\x61\xe3\xc3\xe9\x49\x1c\x5b\x77\x40\xe3\x9d\xca\xe9\xe5\xdd\xbd\x92\x47\x2a\x0b\xe3\x9a\x16\x65\x5e\x06\x38\xf9\x0e\xbe\x22\x06\x25\x2b\x13\x19\x56\xc9\x7c\x9d\x1a\xbf\xc4\xc6\x38\xea\xa9\x4d\x05\x6a\xaa\xe9\x0e\xa8\x16\x6b\xff\x9b\x6b\xcc\x76\x29\xb2\xf4\x1c\xf0\x9f\xa1\xd5\xb7\x84\x85\xc0\xd0\xc2\x82\xe2\x96\xb6\x97\xf2\x11\x6d\xe3\x9b\xd6\x6a\x84\x19\x67\xe3\xad\x39\x9c\x13\xf3\x6e\xf5\x02\x56\x9f\x57\x85\xe8\x0f\x5b\x3e\x6d\xb4\xb1\x6d\x80\xed\x6a\x5f\x92\xac\xef\xfa\x50\xab\x09\x1b\x13\x8c\x07\x8f\xa2\x77\x4e\x9f\xe0\x59\x3e\xc4\x07\xb9\x27\x0d\x9d\xb6\x03\xc1\x2e\xee\x59\xdf\x85\xff\x61\x6d\x62\x79\xf5\x7d\x29\xd2\x7a\x5d\x2d\x4d\xe1\xe4\xb0\xd6\x48\xb1\x31\x9a\x6f\x5a\xbe\x03\xfe\x78\xfd\xf5\x53\xad\x7a\x86\x55\x91\xdd\x91\x96\x09\x57\x69\xd4\x84\x4d\xb0\x72\xbc\x78\x5f\x5d\x8b\xcb\x82\x93\x7d\x5c\x72\x3e\x3e\x8a\x57\x08\x46\x5b\xf3\xc7\xb8\x77\x9a\x03\x62\x08\x06\xbd\x60\x66\xc1\xc3\x7a\x6d\x1d\x09\x9b\xd0\xec\x01\x33\x07\xc0\x36\x38\x71\x13\xe8\x68\x10\x3f\xae\x1e\xe5\xe2\xb2\xee\x07\xf1\xc8\xb4\x9b\xc9\xd6\x14\xc3\x93\xe8\x94\x4c\x27\x99\x17\xb3\xfa\xf6\x49\x34\x3a\x44\xd8\x90\x61\xa8\xc6\x9e\xad\xcc\xc7\xae\x31\xba\x65\xdf\x58\xa7\x4e\x4a\x24\x1b\x02\x99\xad\x64\x5a\xa5\x8f\x2c\x6f\x7c\xc4\xab\xdf\xc9\x84\x32\x62\x20\x7b\xcc\xf1\xdc\x7b\x3b\x70\x8b\xff\x4d\xc2\x9a\x88\xe0\x3e\x91\x7f\xef\x2a\x19\x48\x3d\x32\xb6\x6f\x27\x83\xf4\xe5\x98\x60\xd0\x81\xc3\x69\x3f\xd3\xa7\xb2\x09\xe9\xf9\x02\x47\xdd\x5a\x4c\xb7\xe4\x97\x49\x17\x6f\x80\x1a\x57\x74\x2e\xc8\x38\x83\x4b\x39\xc9\x30\xde\x53\x35\xed\x73\xaf\x7b\xda\xe5\x8c\x02\xcd\xee\xd5\x1c\x36\xfe\xf4\x94\x3c\x51\x0d\x57\xff\x17\x00\x00\xff\xff\x1d\xa2\x77\xba\x93\x04\x00\x00")

func app_js_bytes() ([]byte, error) {
	return bindata_read(
		_app_js,
		"app.js",
	)
}

func app_js() (*asset, error) {
	bytes, err := app_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "app.js", size: 1171, mode: os.FileMode(420), modTime: time.Unix(1422684316, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\x3d\x6f\xdb\x30\x10\x9d\xed\x5f\xc1\xb2\x4b\x32\x98\x8c\x81\x4e\x05\xad\x25\x69\xd1\x2d\x06\xe2\xa5\xe3\x49\x3c\x5b\xe7\x50\xa4\x40\x52\x4e\x9c\x5f\x5f\x4a\x94\x3f\x14\xa0\x40\x90\xa2\x93\x79\x1f\xef\x9e\xef\xdd\xe9\xd4\x97\x87\xc7\xfb\xcd\xef\xf5\x0f\x56\xc7\xc6\x14\x73\x95\x7f\x18\x53\x35\x82\xee\x1f\xe9\x19\x29\x1a\x2c\xee\x6b\x70\x91\x2a\xf6\x6b\xb3\x59\x2b\x99\x7d\x39\x6e\xc8\x3e\xb3\x78\x6c\x71\xc5\x23\xbe\x46\x59\x85\xc0\x99\x47\xb3\xe2\x21\x1e\x0d\x86\x1a\x31\x72\x56\x7b\xdc\xae\xb8\x75\xbe\x01\x43\x6f\x28\x86\x2c\xf9\x89\x12\xe1\x19\x0d\x46\x67\x3f\x5f\x01\xda\xf6\x0a\xac\xe4\xa9\x57\x55\x3a\x7d\x1c\xeb\x69\x3a\xb0\xca\x40\x08\x2b\x5e\x39\x1b\x81\x2c\x7a\x9e\x63\xd3\x68\x0f\x1e\x42\x33\x55\x2f\xdf\xa9\x94\x1c\x27\x84\x4c\x90\x62\x7e\xb2\xb6\x49\x06\x06\x55\x24\x67\x57\x5c\x56\x19\xc4\x59\x83\xb1\x76\x7a\xc5\xd7\x8f\x4f\x9b\xa1\xe2\x15\x8f\x77\x2f\xc9\x35\x9f\x4d\xd9\x03\xbd\xb2\xca\x99\xae\xb1\xa1\x07\xe4\xe2\x84\x46\x07\x8c\xa3\xdd\x8b\x83\x3b\xb4\xba\x78\x40\x03\x47\x25\x47\xeb\x12\x85\x12\x4d\xd1\x7a\x57\x42\x49\x86\xe2\x91\xdd\xdc\x89\xbb\xc5\x52\xdc\xdd\xa6\xe4\x21\x78\xce\x25\xdb\x76\xf1\x44\xde\x2d\xb6\x9d\x31\x8b\x17\xd2\xb1\xe6\xcc\x42\x93\xc4\x1f\x38\xd6\xfc\x6a\x14\x9c\x1d\xc0\x74\xc9\xe0\x4c\x53\x80\xd2\xa0\xee\x85\x9f\xb2\xcf\x67\x33\xdd\x79\xe8\x05\x61\x37\x64\x99\x82\x71\x56\x75\x8c\xed\x77\x29\x77\xce\x80\xdd\x09\xe7\x77\xb2\x7d\xde\xa5\xf5\x6b\x50\x7e\x5d\x83\x0f\xf8\x30\xc2\x78\xd1\x3b\xc5\xc4\xa7\x24\x14\xac\xd7\x1a\xe2\xed\x99\xf0\x73\x2d\x7d\xb4\x23\x25\x27\xf2\x9f\xe6\xfe\x2f\x63\xfb\x09\x64\x3a\x8f\xff\x79\x70\x23\xcb\x87\x47\xf7\xb7\x46\x2f\x8f\xe9\xf2\xb2\xb2\x8b\xe9\xab\x1d\xfb\x55\xd9\x62\x94\xb6\x3d\xb4\x64\xf9\xf9\x73\x22\xad\xd1\x8e\xc9\x17\xda\x42\x08\xa1\x64\xf6\xe6\x02\xb9\xa1\x01\x0f\x07\x3c\xe3\x73\xca\xa2\xf5\xd4\x80\x3f\x4f\x2d\x74\x65\x43\x97\x76\x9e\x06\x44\xdf\xc9\xf8\x6f\x4f\xab\xd1\xef\xca\x78\x00\xae\x3e\x58\x15\x2a\x4f\x6d\xbc\x3e\x2f\x7b\x38\x40\xf6\x72\x16\x7c\x95\xf7\x34\xa4\x45\xad\xb4\xdd\x07\x51\x19\xd7\xe9\xad\x01\x9f\xee\x9c\x6b\x24\xec\xe1\x55\x1a\x2a\x83\x7c\xc3\x36\x3a\xb9\x14\x4b\xf1\x2d\xbf\x45\x43\x56\xec\x93\x2a\x4a\xe6\x7a\xc5\x87\x29\xfb\x33\xf6\x1e\x99\x34\x1a\x6e\x58\xba\x3d\xc3\x25\xff\x13\x00\x00\xff\xff\xa3\xca\xc6\x36\xe1\x05\x00\x00")

func index_html_bytes() ([]byte, error) {
	return bindata_read(
		_index_html,
		"index.html",
	)
}

func index_html() (*asset, error) {
	bytes, err := index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "index.html", size: 1505, mode: os.FileMode(420), modTime: time.Unix(1422684450, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _normalize_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x59\x61\x8f\xdb\x36\x12\xfd\xee\x5f\xc1\xa6\x28\x92\xec\xd9\x5e\xef\xa6\x69\x7b\xde\xeb\x87\x20\x4d\xae\x41\xd3\xe6\xd0\xe4\x70\x07\x04\x0b\x88\x92\x68\x9b\xb7\x92\x28\x88\x94\x77\x9d\xbb\xfb\xef\xf7\x66\x48\xca\x92\x2d\x6f\x52\xe0\x36\x01\x76\xbd\x12\x39\x1c\x0e\xe7\xbd\x79\x43\x9f\x9f\x7d\x25\x2a\xd3\x94\xb2\xd0\x9f\xd4\x3c\xb3\x56\x6c\x9f\xcd\x17\xf3\x4b\xf1\x1f\xf1\xeb\x9b\x0f\xe2\xad\xce\x54\x65\x15\xfe\x5a\x6b\x37\xd7\xe6\xbc\x1b\x2b\xce\xce\x27\x93\xf3\xb3\xb3\x89\x38\x13\x17\x73\xf1\x5e\x39\x91\xab\x95\x6c\x0b\x27\x56\xa6\xc2\x0f\x59\xea\x62\x27\x9c\x11\x56\x56\x76\x66\x55\xa3\x57\x73\x1a\x7c\x39\x17\x7f\x6b\xd4\x56\x61\x8c\x7e\xf7\x5e\x38\x75\xe7\x84\x25\x83\x32\xff\x57\x6b\x9d\x90\x2b\xa7\x1a\x61\x1a\x8d\x11\xd2\x69\x53\x89\x6c\x23\xab\xb5\x9a\x8a\x5b\xed\x36\xa6\xc5\x3a\xda\xca\xb4\xd0\xd5\x9a\xcc\xe1\x5f\x0b\xe3\xe2\x93\x31\x25\xd9\x87\x57\x1b\x57\x16\xe2\xdf\x13\xc1\x8e\xcc\xbc\x23\xcb\x9e\x1b\x57\xe2\x1c\x3e\xd3\x50\x21\x66\xa5\x9d\x91\x0b\x33\x72\x61\xe6\x5d\x58\x8a\x8b\xc5\xe2\x1b\x1e\x75\x19\x46\xdd\xaa\xf4\x46\xbb\xcf\x8e\xfc\x6f\x17\x92\xdf\x55\x69\xb6\xaa\x0b\x49\x29\x9b\xb5\xae\x82\x7f\xa9\xc9\x77\xec\x9f\x7f\xba\x14\x8b\x2b\x3f\x53\xfc\xfc\xe1\xd7\xb7\xcf\x69\x7f\x75\x21\x77\x34\x59\x57\x9a\x42\x60\x31\x58\xfc\xf8\x7f\xfb\xd7\x3f\xbb\x97\xa6\x69\x54\xe6\x44\x92\x16\x26\xbb\x49\xba\xc5\x2b\xe3\xbc\x03\x2a\x47\x1c\x1b\x21\xab\x5d\x70\x4f\x15\xaa\xe4\xd3\xab\xc4\x9b\x57\xe2\x87\xf3\x3f\xcf\xff\x88\x9d\x24\x57\x4e\xea\xc2\x26\x82\xfe\xb0\x6d\x89\x20\xec\x92\x60\xec\x62\x71\x7e\x71\x41\xd6\x64\x95\x8b\xd7\xba\x51\x2b\x73\xf7\xc7\xac\x97\x52\x57\x9d\xb5\x8b\x10\x70\xd9\x38\x9d\x15\x6a\x3a\x91\x56\xe7\xf8\x15\x5c\x98\x4e\x56\x7a\x9d\xc9\x9a\x22\xcc\x9f\xdb\x06\x2f\x57\xc6\x20\x01\xa7\x93\x8d\x92\x39\xff\x5e\x37\xa6\xad\xa7\x13\xb2\x8c\x9f\xaa\x6a\xa7\x93\x4a\x6e\xa7\x13\x0b\x87\x78\x66\xd8\x04\x1f\x69\x70\x6c\x29\xd8\xd1\xab\x5e\x46\x00\x24\xdd\x2e\x74\x85\xf4\x55\xb3\x7b\x36\x33\x0c\x2e\x30\xf3\x5b\x07\xbc\xad\xa2\xed\xc8\x42\xe0\xcf\x75\xc5\x47\x61\x56\x22\xa9\x1b\xb3\x6e\x94\xb5\xbc\xfb\x97\x9b\xc6\x94\x80\x4c\x88\xe1\x94\x03\xfa\xae\x56\x8d\x8c\x21\x69\x73\x6d\xa6\x93\x4c\x56\x5b\x89\x40\xc4\xc9\xd3\xc9\x16\x11\x32\xc3\xad\xf4\xbd\xed\x63\x27\x3a\x32\x63\x47\xb0\x65\x69\x15\x8d\x1c\xc3\x43\x84\x7c\x69\x10\xd4\x4a\xa4\x8d\xb9\x05\x16\xad\x58\xc1\xcf\xb8\x12\x10\x2d\x12\x76\x2c\xe9\xb0\x9e\x01\xc2\x8d\x29\xec\xbc\x07\x2a\x75\x97\xc1\x55\xb1\x51\x7a\xbd\xe1\x34\x24\x1e\x01\x6c\xd4\x16\x6c\x65\xfb\x1b\x5c\x22\xa4\x4f\x3e\x46\x1b\xd7\x4f\x87\xfb\xaa\x0c\x7c\xc5\x03\x6f\x68\x0f\x43\x76\xf8\x45\x9e\x53\x40\x44\xf2\x71\xa3\xf3\x5c\x55\xd7\x89\xb0\x6e\x47\xac\xc3\xe7\x54\xe3\xe5\x00\x04\xe7\x17\x0b\xf6\xf1\x67\xc4\x4f\xb8\x8d\x12\x89\x53\x25\xd6\x71\x2a\x39\x46\x0c\xb2\x7c\x2a\xde\xcb\x95\x6c\xf4\xb4\x9f\xeb\xe2\x2f\xe2\xf2\x32\x6c\x20\xae\x3b\x9d\x44\x43\x63\xde\x7b\xde\x78\xab\xab\x9b\x07\xe4\x88\x10\x77\xda\xd5\xba\x41\xa2\xa6\x32\xbb\x21\x54\xc0\xf1\xcc\x14\x80\x1d\x9f\xa2\x04\x1c\x30\xaa\x20\x5f\x3a\x3c\xc7\xd3\x60\xdf\xf7\xf3\x66\x3c\x6f\x29\x5c\x03\x56\xae\x65\x83\xe8\xf4\x83\xff\xa6\x44\x42\xc2\x56\x03\x0c\xca\x54\x17\xda\xed\xc4\xed\x46\x55\xc0\x78\x06\xb6\xcf\x39\x66\xb2\xb0\x06\xf9\x84\xbf\xc5\x06\x83\x1b\x8f\x1a\x59\x14\x5d\x76\xc5\xc5\x97\xde\x35\xe0\x7f\xc9\x23\xd9\x19\x64\x17\x25\x6b\x8f\x7d\x3f\x10\xbb\x17\xc8\xd3\x42\x58\x55\xca\x0a\xc9\xfd\x80\x41\x8d\x09\xf6\xb9\xac\x3a\xca\x15\x8f\xed\xb8\xb7\x34\x6d\x3e\x3a\xed\x0a\x75\xed\x43\x6c\x1a\x00\x6c\x96\x1a\xe7\x4c\x89\xf2\x54\xdf\x89\x1c\x9f\x55\x3e\x96\xdb\xb4\xb4\xc2\x5e\x1d\x95\xe9\x24\x35\x05\xa6\x32\x7d\xc4\x74\xfc\xf6\x4f\xf7\x2c\x9d\x82\xf9\x80\x2b\xb8\xde\xd5\xda\xdb\x80\x24\x32\x75\x6a\xc1\x91\xbd\xfa\x25\x8e\x57\xc8\x57\xd5\xde\x36\x3b\x0b\x32\x72\x60\x9b\x6c\xcc\xf8\x16\x36\xa0\x0b\x80\xbc\xcd\x45\x12\xe6\xb0\xae\x80\x59\x5f\x69\x99\x55\xf0\x2b\x09\xd4\x9d\xf0\xbb\x24\x94\x87\x84\x8c\x11\x57\x20\x0d\xec\x17\x47\x61\x73\xd1\x73\x11\xcb\x2d\xc5\xa5\x2a\xaf\xfa\xc5\x7d\xfe\xdd\xf7\xaa\x1c\x27\x97\x7b\xcf\x3e\x2c\x00\x3b\x37\x07\xe8\x59\x8a\xaf\x57\xab\x05\x2d\x12\x60\xf4\xf5\x62\x31\x6a\x5f\x57\xd8\x8f\xd5\xd6\x91\x61\x72\xbd\x8b\x11\x6b\x34\x0e\xcf\x38\x66\x6c\x49\x0f\x0f\x76\xf6\x03\xb4\xce\x08\xa7\xa3\x8a\xa7\x21\x94\xb6\xad\xf1\x69\xb5\xa2\xf8\x12\x9b\x73\xe1\xf0\xfc\x9a\x9c\x5a\xa9\xa5\x44\x6a\xeb\xc3\xc5\xbe\x7f\xfe\x0d\xed\xb0\x67\x81\xa1\x2a\x44\x6d\x2c\x8b\xa2\x25\xe8\x01\xb4\x08\x60\x5f\xdd\x57\x8c\xc8\xe1\x68\xde\x99\x7a\x29\x66\x8b\xf9\x73\x3a\x22\x7e\x9e\x06\xd4\x78\xb8\xe0\xd5\x65\x7c\x07\x42\x78\x55\xa6\x0a\x1c\x9c\xfb\xac\xa8\xdc\x83\x33\xac\x47\xaf\xa7\x3a\x5d\x91\x58\x41\x76\x8e\xd6\x8f\x48\xac\xba\x5c\xf7\x70\x7f\x50\xc3\xa2\xde\x20\xce\x5b\x15\xe6\x96\xf3\xcc\x97\x95\x60\x2b\x10\x4c\x3c\x8a\xed\x9a\x2b\xe6\xb2\x81\x06\xf2\xc5\x32\x4e\x5d\x86\x79\x31\x34\x7f\x25\x51\x44\x47\xfc\xe0\xa1\x89\xb9\x1c\x20\x3c\x0a\x15\x4e\x3e\x0f\xd1\xb0\x15\xaf\xe6\x06\x2a\xfb\x02\x28\xfc\x76\x51\xdf\x8d\x01\x25\xd7\xc8\x59\x54\x21\xe8\x07\x91\x2a\x77\xab\xd4\x1e\xfe\x64\xdb\xa0\xf6\x35\x87\x99\xbb\xf1\x85\x64\x56\x9a\x4f\x20\xdc\x3b\xca\x5b\x44\x64\x19\x43\x42\xcf\xae\xf8\x68\x4e\xbe\x1a\x55\x1e\x2f\x31\x06\x5a\x73\x7f\x6a\xe3\xb8\xa9\xc3\xf6\xf6\x27\x24\x5b\x67\xc6\x36\x67\x72\x20\x53\x95\xc9\xac\x45\x37\xd1\x43\x3e\xf6\x8b\x9c\xa1\x43\x1c\x5f\x21\x33\x24\x96\x6f\xd2\x9c\xf4\x21\x3e\x59\x59\xd6\xc7\x7d\x55\x69\x2a\x83\x12\x9e\x41\x6d\x76\x1f\xaf\x86\x50\xbe\xd8\x43\xea\x35\x64\xec\x03\x16\xd5\x5f\x2a\x73\x5b\x81\x33\x4a\xed\x7b\x47\xf0\xc0\x2e\x36\x61\xd3\x40\xdc\xbd\x64\x11\xe8\x2e\xa1\x1d\xff\x49\xbb\x47\xa0\x11\xc9\x9d\x9f\xac\x72\xb2\x16\xe9\x99\xe4\x35\xd8\x04\x50\x4a\xa6\xa2\xad\x0a\x0a\xaa\xa4\x82\xd9\x70\xc1\x84\x56\x81\xb2\x86\x44\xd1\x96\x8a\x69\x88\xde\xb1\xec\xf7\x5a\x89\x12\x38\x55\x3e\xea\x48\x2a\x5a\x6b\x1e\x1a\x59\xef\xbd\xb6\xb6\x45\xd0\x3c\x8b\xda\x30\x0b\x2e\xf8\xb6\x17\x6c\x14\xe8\xc0\xc6\xd6\x20\xda\xe7\x93\x0d\xce\x68\x24\xf2\xa9\x95\x9e\xcd\x0f\x40\xc5\x6e\x77\x18\x70\x68\xdb\xbf\x58\x03\xb4\x20\x4e\xb4\x40\xba\xaa\x5b\x37\x9d\x98\xda\x85\x6e\xc9\x87\x8b\xb4\xeb\x9d\x83\xbc\xf3\xfa\x2f\x54\xab\xe0\x4d\xbf\x95\x20\xd7\x87\x2f\x42\xe7\xbd\x6f\x91\xe9\xe1\xb3\x61\x53\xd1\x69\xf4\x88\x81\xa4\x13\x33\x9e\xad\x92\x43\x25\x35\x70\xfb\x00\x3e\x5b\x6d\x35\xe2\xfb\xd9\x3a\x9a\xf0\x2d\x00\x2b\x57\xb4\x9c\x65\x12\xdd\x96\x60\x0f\xdf\x84\x7a\xf3\xb1\x24\xfa\xc4\xe1\xc8\xbf\x00\xc8\x3c\x95\xd0\xcc\xd8\xdc\x74\x07\x0a\xa9\xc6\x87\x16\x0c\x1e\xaf\xb4\x95\x45\xeb\x3b\x9c\x5e\x47\x1c\x16\xf3\x62\xae\xef\xcb\xfe\x10\xa7\x07\x72\x72\xd0\x0d\xee\x4d\x05\x57\xef\x35\x75\x70\xf0\x7e\x8a\x2f\xae\x03\x67\x07\x0d\x4a\x44\xc2\x8b\xad\xd1\x39\xf7\x11\xff\x50\xe9\x2f\xd8\x61\xda\x32\xf9\xbc\xa8\xf2\x86\xde\x7c\x3b\x5f\xcc\xcf\xa8\xfc\x81\xdb\x9e\x5c\x3e\x05\x72\x49\x61\xee\x90\xca\x5c\xe7\x63\x77\x18\xe0\xc2\xe1\xe5\x7e\x35\x19\xf6\x89\x3d\x4c\xe8\x2a\xf6\x0f\x74\x13\xc5\xfb\xca\xa0\x1f\x6f\xbc\x50\xe4\xb4\x4d\x84\xdb\xd5\xca\x86\x46\x32\x22\x24\xb6\x20\xad\x8d\x06\x24\xb7\x3b\x21\x0d\xb2\x1d\x41\x32\x6b\x1b\x8b\x03\xf7\x76\x63\xdd\xd0\xa5\x5c\xab\x19\xd9\x0c\x7e\xc6\x65\xba\x42\x62\x0f\x82\xc8\x17\x56\x3c\xe8\x23\x4d\xfb\xf1\x91\x7f\xf1\xe8\x7a\xda\x41\xa4\xff\x96\xaa\x9f\xc3\xcb\xc1\x43\x48\x19\xf0\xd6\x23\xdf\x05\xc4\xfb\x2a\x59\xd7\x4a\x36\x74\x82\x20\x42\xb6\xd9\xc7\x96\xf7\x7e\x09\x41\xa5\x51\x95\x9a\x31\x84\xfd\xae\x66\xb6\x77\xb3\x17\xf6\x4b\x49\x3e\xca\x46\xdd\x9e\x3e\xc6\xd7\xd7\x83\xdd\x75\x4f\x3d\x1d\x84\xf5\x83\xf5\xab\xe3\xeb\x33\x5d\x55\xc0\x4a\x2d\xf3\x9c\x58\x8c\xe2\x17\x14\xd2\x80\x9f\x06\x2b\x2f\x97\x5c\x8f\xb9\x51\x9c\xf1\xfc\x10\xa7\xe3\x17\x87\xc2\x49\xc4\x95\x4e\x5c\x05\xec\x57\x24\x9a\x19\x11\xb8\xa0\x94\x78\xd6\xad\xe5\xd7\x5f\xe9\xb2\x36\x0d\x30\xc4\xf2\x97\x8c\x51\xf2\xff\xfd\x85\x4f\x19\xbb\x51\x5d\xb9\xe0\x79\xec\xd1\x40\xf1\xfa\xfb\xd6\x41\x6f\xec\x1e\x5b\x94\xee\xcc\x94\x25\xd5\x6f\x82\x93\x74\x62\x67\x5a\xb0\x47\xf5\x18\x32\xdf\xd1\x7d\x81\xdb\xe7\x3b\x56\x44\x8b\x3c\x28\x1a\x61\x27\x30\x04\xff\xfc\x0b\x7f\xd5\x9a\x1b\x65\xc9\x08\xf6\x5b\x13\x7c\xf6\xda\x65\x1a\x83\x33\xa5\x8b\xbb\x5b\x9d\xbb\x0d\x99\x8a\xc0\x0e\x21\xc2\x78\xe1\xc7\x77\x4c\xdc\x93\x3c\xc9\xf1\x75\x09\xa0\x3a\xbc\xd5\x89\x87\x3d\xa6\x75\x7b\xe9\x9e\x6d\x54\x76\x03\x93\x87\x30\x68\x24\x08\xe2\x51\xec\x85\xf7\xc2\xab\xeb\x8b\xef\xfa\x65\xa7\x77\xde\x23\x77\x56\xaf\xf5\x1d\x1f\xd7\x00\xe5\x94\xfb\xbe\x04\x3e\xe6\xba\xd0\x70\xf4\xce\x73\x15\x3e\x05\x9c\xd9\x39\x49\x1d\x91\xa1\x14\x4b\x7f\xf0\x49\x27\x88\x22\x8f\x13\x83\xf0\x4d\x91\xcf\x99\x29\x7a\x5b\x91\xc9\xd6\xe2\xcd\xd1\xb2\x7e\x28\xd9\x39\x5c\x89\x82\xec\x6f\xc8\xfd\x35\x4c\x12\xe0\x94\x70\xf4\x89\x94\x93\x91\xf8\x55\x2d\x7a\x9b\xe6\xd1\x35\x50\x11\xb8\x82\x21\x31\xb3\x10\xf5\xb3\x41\x55\x3f\x39\xc1\xb4\x6e\x38\x81\x83\x1e\x13\xf7\x50\x8c\xf6\x92\x24\xd9\xd3\xd2\xbe\x60\x5b\x3c\xc9\x36\x2b\xad\x8a\x3c\x39\x79\x27\x40\xe9\xd2\x59\xd9\x1f\x6f\xd2\xbb\xc3\x88\xc7\x3c\x6e\x24\x30\xf2\x13\x1c\x5c\xd1\x52\x77\x45\x9c\xc0\x91\x5a\xb5\x0e\x0d\xc3\x0c\xa4\x6f\x56\x4f\x47\x02\xe6\xfd\xbb\x87\x5f\x29\xd2\xec\xfe\xf0\xeb\x85\xfb\x7b\x83\x68\xe8\xd4\x90\x1e\x5b\x9f\xb4\xf2\x05\xac\xe9\x9d\x47\x72\xc1\xd3\x22\xe6\xcd\x58\x7c\xa2\x20\xe6\x60\x87\xb7\x4f\x30\x9e\x65\x89\x1f\xf2\x94\x0a\x68\x1d\x52\x74\x60\x90\xdb\x56\x7a\x1c\x96\xf3\x94\xb6\x91\x96\x8c\x45\x77\x9e\x70\xd5\xee\x82\x85\xaa\xd8\xc5\xf0\xbe\xb0\xef\xd3\xce\x3f\x99\xf9\xa5\x47\x53\xf5\xe4\x1c\x40\xc7\x34\x9e\xe8\x4e\x9d\xe2\xa1\x68\xf9\x89\xef\xe3\x45\x4f\xfe\xf9\x14\x9b\x06\x65\xea\xb5\x54\xd8\x5c\xd7\x7b\x62\x63\x94\x91\xfd\xfa\x42\x37\x71\xd6\x14\x10\x39\x5f\x67\x0b\xfa\x3f\xb8\x23\x12\x97\xf5\xdd\xb0\x00\xcd\x9f\x3d\xa7\x5b\xa3\xf9\x77\x97\xfe\xf7\xf7\xfb\x7b\x89\xa3\xaf\x13\x58\x57\x27\x63\x7a\x7f\x4c\xff\xf6\x38\x37\x1e\x8a\x35\xa2\x56\x06\x95\x40\xd0\x55\xec\x63\xa6\x21\xba\x5b\xa7\x1b\x78\xcd\xcc\xb3\x13\x9f\x54\x63\xf8\x41\xdc\x5e\xac\xf8\x85\x5a\xa3\x10\x1d\x14\xd3\x2f\x26\xd9\x83\x2f\xca\xba\x2f\x37\x6c\x06\x39\x57\xa4\xb2\x39\x21\xe1\x07\x8d\xc5\xe9\x1e\xf8\x27\xae\x89\x51\x50\x33\xdb\xf6\x6e\x2a\x13\xa4\x63\x5d\x17\x1a\x81\x42\x9f\x28\x45\xd3\x52\x08\x52\x98\xe3\x5c\x14\xbf\xbd\xfb\xf0\x6a\xc9\xb3\x3a\x05\x24\x2b\x0a\xb3\x95\x2b\x85\x36\x29\x55\x81\x7a\xf3\xfd\x97\x2e\x23\xed\x65\x70\x39\x76\x47\xf7\xde\x97\x8a\x0f\xa4\x90\x1e\xfe\x32\xbf\x34\x16\xdb\x40\xb3\x4e\xe7\x1f\xa5\xab\x63\x59\x0c\x5c\x15\xf1\x70\xfd\x93\xde\xcd\x32\x72\xad\x90\xb5\x55\x44\x42\xfe\xd3\xd5\xfe\x65\xb0\x17\xf5\x93\xcb\xd1\x00\x6e\x78\xf6\x40\x59\xfd\x2f\x00\x00\xff\xff\xae\xc8\xcd\xa3\x75\x1e\x00\x00")

func normalize_css_bytes() ([]byte, error) {
	return bindata_read(
		_normalize_css,
		"normalize.css",
	)
}

func normalize_css() (*asset, error) {
	bytes, err := normalize_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "normalize.css", size: 7797, mode: os.FileMode(420), modTime: time.Unix(1419877841, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _skeleton_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x3a\xdb\x8e\xdb\x38\xb2\xef\xfe\x0a\xa2\x07\x03\x4c\x02\x5b\xed\x7b\xba\xdd\x98\x83\x93\x4b\x77\x66\x81\x49\x82\x4d\xb2\xbb\x0f\x8b\x79\xa0\x25\xca\x22\x5a\x12\x35\x24\x15\xb7\x27\x18\x60\xff\x61\xff\x70\xbf\x64\xab\x28\x51\x37\x93\x8e\xf7\xc9\x71\xd2\xb0\x55\xc5\x2a\x16\x8b\x75\xa5\x78\xfd\x7c\xf4\x9c\x7c\x7a\x64\x29\xd3\x22\x27\x7f\x9f\x07\xd3\x60\x09\x90\xd7\xa2\x38\x48\xbe\x4b\x34\x99\x4f\x67\xcb\x31\x79\x43\xbf\x30\xf2\x96\x66\x34\x4c\x18\xa0\xf7\xfb\x7d\xb0\x63\x5a\xd5\x74\x41\x28\x32\x80\x3e\x48\xc6\x88\x16\xa4\x54\x8c\x94\x79\xc4\x24\xd1\x09\x23\xef\xfe\xf2\x99\xa4\x3c\x64\xb9\x62\x01\x0c\x4a\xb4\x2e\x36\xd7\xd7\xc8\x41\x14\x00\x14\xa5\x0c\x59\x20\xe4\xee\xba\x1e\xa4\xae\x33\xae\x27\x96\xa2\x48\x0a\x20\x9a\xcd\xaf\xe7\xb7\xd7\x28\xca\xe8\xf9\xf5\x68\x34\xba\x7e\x4e\x3e\xd3\x6d\xca\x88\x88\x49\x28\x72\xcd\x72\xad\x46\xff\xf9\xd7\xbf\xbf\xc3\xff\xa3\x09\x79\x2b\x79\x04\x5f\xaf\x28\xe8\xe5\x93\x3e\xa4\x4c\xc1\xd3\xe7\x43\x21\x76\x92\x16\xc9\x01\x1e\x7e\xe5\xf9\x23\x02\x5f\x95\x1a\xd4\x89\xbf\x1e\x84\xcc\x94\xc1\x28\x8d\xdf\xaf\x45\xc4\x90\x0a\x57\x8d\xcf\x9f\x0a\x1a\xf2\x7c\x07\xbf\xfe\xa6\x79\xca\x35\x37\xd0\xd7\x29\xa3\xb2\x02\xbf\x63\x11\xa7\xe4\xaf\x25\x93\x88\xb2\x5a\x33\xa2\x5c\x5c\x27\xce\xff\x04\x64\x0c\x70\x33\x29\xcf\xc1\x74\xbe\x8e\x08\x29\x84\x82\x95\x89\x7c\x43\x24\x4b\xa9\xe6\x5f\xd8\x1d\x40\xf7\x3c\xd2\xc9\x86\xcc\xa6\xd3\x1f\xf1\x31\xa3\x4f\x93\x1a\x74\xbb\x9e\x16\x4f\x15\x4c\xee\x38\x90\x4d\x09\x2d\xb5\x40\x48\x41\xa3\x08\x14\x83\xa0\x79\x3d\x68\x2b\x9e\x26\x8a\xff\x61\xa0\x5b\x21\xc1\x5e\x27\x00\xba\x23\x7f\xa2\x18\x69\x99\xe5\x63\xfb\x43\x19\x69\x06\xf3\xc6\xa9\xa0\x7a\x43\x52\x16\xeb\xd3\xdc\x50\xef\xb0\x9d\x24\x62\x5f\xc0\xa8\x15\x49\x41\x38\xe3\x1b\x34\x27\xcb\x29\x08\x83\x2b\xff\xff\xcc\xec\xd7\x4f\x19\xcf\xed\x6a\x0c\xee\x99\x99\x7a\xa0\x97\x46\x96\x9b\x95\x11\xa5\xbb\x3c\x9c\xf1\xf4\x9c\xab\x95\x7f\x4e\x83\xfb\xc6\x9c\xb0\x7c\x98\xc3\x0c\xa8\xb4\xd4\xfc\x54\xf5\xc8\x4a\xfd\x13\x54\x0d\x2c\xa3\x3f\x7c\x13\x73\xa9\xf4\x24\x4c\x78\x1a\x75\x49\xbb\x70\x17\x1b\xb3\x30\x1c\x2f\x72\xd6\x9d\xb9\x7d\x54\xc4\xf1\xf9\x6a\xc5\x5e\x06\xeb\xe6\xf3\xc2\x8a\xa4\xf7\xe2\x2c\xe2\xd9\x22\x58\x34\x9f\x86\x38\x81\x78\xe7\x27\x6f\x88\xe7\x73\xa0\xe8\x7c\x0c\x71\x0c\x71\xef\xc4\xd4\x0d\xf1\x62\x1a\x1c\x8b\x1d\x83\x27\x9c\x45\x7c\xeb\x10\x5b\xf1\xa7\xf3\x14\x76\xe3\x10\x5b\xb1\x2f\x2c\x3f\x63\xcd\xab\xb5\x43\x6c\x86\x09\xe5\x0c\xe2\xf5\xca\x21\x76\xce\x4f\x6e\x74\x43\xfc\x62\x19\x4c\xbb\x82\x57\x5b\x75\x4a\xe8\x0e\xf1\xcd\xdc\x25\x76\x7a\x72\xd1\x0d\xf1\xed\xcc\x65\x24\x7b\x96\x9e\xd8\xac\xaf\xbd\xa8\xe2\xb7\xf8\x89\x4e\xb8\x8c\x6a\x36\x3e\x1e\x4e\x5b\x01\x13\xaf\x88\x95\x93\xfa\x84\xd2\xed\xd4\x09\x4d\x63\xf7\xcc\x03\x63\x31\x24\x10\x78\x3e\xc4\xb1\x82\xe2\x00\x43\x0c\xb2\x30\x4f\x93\xed\x61\x32\xf4\x5c\x07\x42\x75\x79\xf7\xb4\x71\xe3\xf2\xdf\x96\x45\xeb\xc9\x63\x1f\xc2\xcf\x7b\xf6\xc2\xb1\x73\x1d\x16\x1d\x47\x1f\x72\x77\xc5\x80\x01\xf7\xf9\xda\xe1\x48\x2d\x8b\x4e\x24\x18\x7b\x31\xca\x27\xfa\x62\xe9\xd8\xf3\x0e\x8b\x36\x52\x0c\x99\xbb\x62\xc8\x80\xf9\xd2\x15\xf6\x5a\x16\x6d\x24\x19\xfb\x10\x7e\x9d\xaf\x5c\x51\xb1\xc3\xa2\xe3\x73\x43\xee\x2e\x77\x1c\x70\x5f\xbb\x5c\xa1\x65\xd1\x8d\x44\x63\x3f\xca\xcb\xdd\x15\x57\x5b\x16\x9d\x50\x35\xf6\x62\xbc\x4a\x7f\x71\xe3\x08\x60\x1d\x8b\xf3\xa9\xc5\x19\xe3\x86\x4e\xe4\x0a\xcb\x9d\xb5\xa7\x7e\xad\x3b\xa3\xe0\x80\xfd\xad\x3b\x86\x74\xfd\xbc\x17\xc8\x8e\xc3\x40\x0f\xad\xfe\x77\x7b\x3f\x8a\x76\xc7\xd1\xa0\x8f\x57\xe7\xec\xee\xb1\x98\x9d\x98\xe8\x58\x44\x07\xab\xbc\xb6\x0f\x6c\xff\xac\x8a\xf2\x6e\x63\x70\xf1\x32\xdc\x5b\x9b\x83\xa0\xef\x3f\x7c\xbe\x1f\x25\x3a\x4b\x09\x57\x04\xd6\x8b\xcd\xde\x7a\x1e\xac\x7e\x24\x4a\x60\x6d\xa9\x09\x4d\x53\xd3\xf4\x7d\xbc\x7f\x47\x32\x46\x55\x29\x59\x86\xcd\x19\x00\xa5\x28\x77\x89\x28\x75\xd3\x6d\x8e\xa8\x64\x64\x0b\x6b\x8f\x08\xb4\x9e\x33\x2c\x4a\xab\x02\x3a\x20\x9f\x04\x22\x78\x08\xec\x0e\x64\x16\xac\x80\x0b\xf9\x99\xcc\x56\x30\x64\xf3\x0c\x85\x31\x42\x60\x95\x18\x43\x91\x8a\x75\x37\xdb\x54\x92\xa0\x5e\xb7\x22\x3a\x0c\x91\xc0\x84\x65\x77\x98\x9d\xc2\x52\x4a\x10\x09\x18\xb3\x4c\x91\x90\x62\xb7\x1a\x82\x74\x19\x08\x53\xee\x48\xc6\x15\x87\x86\x52\x16\x92\x69\x90\x05\x3a\x0f\x18\x05\xf2\x19\xa6\x20\x38\x2e\xa7\xca\x6c\x29\xc7\xad\x36\xf1\x02\xf9\xaf\xef\xec\x8c\xfb\x1a\x06\x05\x7c\x03\x8b\x69\xc6\xd3\xc3\x86\x5c\x7d\xa4\x29\xdb\xd3\xc3\xd5\x98\x5c\xfd\x82\xa5\x81\x86\x55\xbe\x67\x25\xeb\x01\x48\x0d\x69\x00\x63\xf2\x52\x72\x9a\x8e\x89\xa2\xb9\x82\x08\x28\x79\x8c\xac\xc1\xc6\x84\xdc\x90\x1f\xe6\xf3\xb9\x31\x28\xd3\x18\xb7\x8d\xe5\xc5\xed\xc6\x6b\x4c\xc9\x6c\x4c\x92\x39\xfc\x2d\xe0\x6f\x09\x7f\x2b\xf8\x5b\x9b\x4d\xab\x3d\x45\x8b\x02\x4b\xa0\x16\xb0\x15\xd0\x1b\x67\x90\x4e\x61\x47\x8e\x54\xbd\x98\x9a\x6a\x29\x99\x81\xab\x75\x76\x1d\x6a\x41\x1c\x3d\xdc\x2a\x50\x16\xf4\x6d\x1a\x76\x79\xa2\xaa\x3e\x7a\x43\x26\xc1\x0c\x87\x02\x8f\x79\x9f\xc7\x22\x58\x3b\x79\xac\xee\x7c\x3c\x50\x90\xc5\x90\x89\x53\x90\x85\x5f\x10\x64\xb2\xec\x33\x99\x07\x4b\x27\x13\xa7\x24\xd3\x1b\xcb\x65\xd5\xe7\x32\x0b\x6e\x5c\x5c\x56\x4e\x51\xa6\x2b\xcb\x65\x3d\xe4\xb2\x72\x71\x59\x3b\xb8\x4c\x6d\x13\xfc\x6b\xa7\x09\x2d\x12\x3c\xc9\xd0\xdf\x6e\x43\x87\x5b\xba\xaa\x35\x89\x81\x7f\xb8\x55\xcb\x60\xde\xe2\x8e\x76\x60\xdd\xe2\x96\xee\xdd\x31\xb8\x95\x5b\xe9\x06\xe7\x51\x82\xe9\xb8\x0b\x87\xf5\x5a\x9f\xac\xce\x77\x2e\xee\x79\x5e\x77\xa4\x46\x76\x1b\x4d\x66\xf7\x2f\xef\xdf\xbc\x42\xe1\xe9\x26\x11\x5f\xea\x03\x00\x8b\x9d\x3e\xbc\x9c\xbe\xbe\x6f\x96\x66\x8f\xac\x2e\xbe\x0e\xef\xe2\x82\xad\x11\x71\x3c\xb2\xdf\x3c\x2f\x4a\xfd\x4f\x7d\x28\xd8\xcf\x57\xaa\xdc\x66\x5c\x5f\xfd\xd6\x87\x4a\x06\xd9\x6d\x08\xac\xc8\xaf\x7e\x33\xda\x88\xb8\x2a\x52\x0a\x01\x9d\xe7\xc6\x05\xb6\xa9\x08\x1f\x31\x2c\x59\x5f\x58\xdc\x54\xa7\x4c\x9d\x83\xa7\x45\x7d\xf0\x64\x35\xb9\x5a\xad\xf0\x51\xb3\x27\x3d\xa1\x29\xdf\xe5\x1b\x12\x32\xcc\x3d\x77\x83\xdc\x35\xab\xe8\x7a\x11\x6f\x5d\x25\x97\x9e\xff\xd9\x39\x87\x1e\x58\x45\x14\x3b\x95\x96\x90\x43\x62\x21\x21\x96\x96\x45\xc1\x64\x08\x29\xb8\x41\x46\x2c\x14\x92\x56\x87\x6e\x39\xd4\x32\xe6\xc0\x2d\xe1\x9a\x19\x6e\x0c\x81\x7b\x48\x2e\xe6\xc4\x8b\x86\x8f\x3b\xc8\xea\x79\x34\xa9\x57\x64\x38\x17\x14\xd3\x6b\x75\x24\x66\x8e\xc1\x24\x8d\x78\xa9\xc0\x3b\xed\xb1\x1b\x42\x61\x55\x98\xef\x45\xca\x23\xf2\xc3\x76\xbb\x35\x7a\x29\xa5\x42\x36\x85\xe0\x56\x0b\x27\xce\xe8\xaa\xed\xa8\x0c\xd4\xee\xad\x7d\x72\xed\xb0\x0b\x57\xef\xb3\x0b\x65\x77\xdb\xe2\xec\x7c\xb1\x08\x4b\xd5\xcc\x57\x3f\x39\xe7\x73\xe0\xec\x7c\x0e\x54\x33\x9f\xc1\xf5\x3c\x0e\xca\xd0\x8e\x3a\x2d\xf4\xe6\xe6\x06\xa1\x50\x51\xa1\x0d\xd4\xd1\xa6\x96\xb2\xfe\x9a\x14\x92\x43\x40\x3a\x58\x71\x8f\xc0\x2e\xb9\x4f\x0e\xaa\x17\x70\x72\x8c\x5d\xc9\x60\x50\x6f\x49\x0f\x0f\x0f\x6e\x13\x82\xc5\xbe\x5e\x3c\x4c\x1d\xeb\xad\x11\xde\x45\xf6\x2d\xc1\x83\x3c\x63\xc1\x27\xec\xe4\x8c\x91\x9e\xc5\x0f\xac\x68\x88\xed\x19\x95\x07\x79\x8e\xe8\x7e\x93\x3b\x63\xa4\x4f\xf4\x63\x83\xf4\xef\x5e\x9d\x3a\x8e\x77\xaf\xcd\x29\xf6\x6c\x3a\xfb\x9e\x73\x46\x57\x2f\x2c\xa3\x3c\x1d\x26\x83\xbc\xcc\xb6\x4c\x0e\xa1\x8a\x51\x19\x26\x43\x28\x06\xd6\x63\xd8\x11\xcb\x52\x1e\x81\x0a\xaa\xd4\x1e\x14\x89\x70\xe4\x02\xa1\x95\x8e\x47\x0a\xda\x90\x50\x57\xd5\x91\x2f\xdb\xac\x21\xba\x62\x4b\x65\x5a\x9e\xcf\xd0\x91\x21\x00\x4c\x50\xd7\x4d\x55\x95\x68\x94\x09\xfa\xd8\xde\x3c\x3c\x8c\x09\x64\x20\x21\xa1\x1b\xdb\x1e\xc8\x3f\xd8\xf6\x91\xd7\x7d\x8e\x63\x9b\xe3\x38\xf6\x84\xf2\x37\x33\xfc\x77\x2a\xfa\x43\x40\x4f\x68\x24\xf6\x6d\x7e\xf1\x07\x79\x90\xfd\x23\xcb\xc0\x75\x14\xa1\xfb\xc7\x3d\x95\x11\x89\x58\x4c\xcb\x54\x13\x65\xfa\x65\x14\x5d\x61\xeb\x66\xb4\xa6\x20\x45\x4a\xc2\x3f\x7c\xfa\xce\xb7\xd0\x6c\x1d\x64\x72\xd4\xf1\x84\x42\x0e\xa6\x90\x38\xab\xe4\x5a\x69\x04\x3e\x93\x4c\xfc\xe1\xc3\x99\xcf\x11\x0e\xf4\xd5\x9b\x00\xeb\x69\x6b\x1f\xeb\x55\xcf\x3e\xaa\xfa\x74\xdd\x87\xd9\x06\x0b\xc1\xc0\xcb\xa1\x3f\x57\xd4\xb0\x5a\x74\x86\xa9\x5a\x97\x2e\x5c\xa5\x51\x37\xc6\x33\x95\xd1\xae\x0b\xd1\xea\xd8\x62\xad\x22\xec\x73\xe5\x32\x9d\x50\xe6\x30\xdd\x36\xf3\xf4\x73\x6a\x4a\xb7\x2c\x1d\x8f\x52\xb6\x63\x79\xd4\x2f\xfe\x9a\xaa\x6f\xd0\xa3\x56\x7d\x81\xab\x64\x03\x7e\x31\x67\x69\x84\xa7\x27\x5f\x7b\xe5\x61\xc7\x67\xea\x2e\x68\x3a\xdc\x86\x30\x61\xe1\x23\xb8\xc6\x51\xb9\x0a\x3e\x26\xdc\x85\x69\xb3\x00\xf2\x7f\x24\x30\x3f\x26\xcd\xf1\x88\xb7\x86\xed\x1d\x56\xb9\xd7\x02\x91\x22\xa3\x69\xa7\xc1\x51\xdf\xed\x1b\x78\x0c\x06\x65\x75\x5a\x94\x82\x98\x13\x13\x39\xa0\xd4\xe6\x32\x4c\x31\x70\x28\x1e\x19\x35\x89\xe3\x31\x50\x0d\x43\x0e\x4c\x7b\x83\xc6\xa4\x66\x66\xfd\xc6\xbe\xad\x71\x75\x7f\x30\xb4\x04\xdb\x81\x2f\x20\xc4\x19\xea\xaf\x9a\x85\x7d\x4f\x5d\x1f\x71\x4d\xed\x8f\x45\x57\xe5\x55\x13\x70\x5b\xbd\x76\x4d\x79\xb7\xcd\xb4\xf6\x66\x4f\x0c\xcc\x5e\x98\xab\x02\x17\xd7\xba\x77\x2b\x42\x10\xaf\x6f\xfa\xa6\x6b\x6f\xed\xac\x7d\x79\x1f\xcc\xdd\x7a\x38\xa7\x27\xc1\x32\x65\x86\xff\x3c\x89\xea\x7e\x86\xff\x3c\x89\x0a\x54\x59\x48\x06\x2e\xd3\x08\x7b\xec\xf2\x8d\xf8\xa8\x7c\x7b\x14\x30\x94\x0c\xb8\xb4\x07\x73\xd5\xdd\x8d\x8b\x6f\x82\x77\x67\x74\x02\x81\x33\xea\xef\xcd\x6c\x8e\x85\x44\x9d\x3d\xba\x9d\x6a\x7b\xef\xa1\x4e\xd8\xb5\x25\x1e\x69\x18\xd3\x52\xd2\x7f\xeb\xaf\xa3\xa3\xb7\xfd\x43\x67\xaa\xa8\x52\xda\x23\x6a\x1f\x7b\x34\xb2\x0a\x49\xed\x71\x8b\xbd\x1c\x73\x71\xad\x7a\x55\x6d\x8f\x22\xea\x5a\xfb\x94\x4f\x9b\x40\x7f\x5c\xfe\x8d\xfb\x79\x64\x48\xdc\x9c\x4c\x81\x05\x42\x63\x81\x46\xfb\x7b\x29\x34\xfc\x8e\x52\x24\xdd\x95\x08\xd7\x68\x92\xe3\x51\x81\x11\x0a\xe3\x12\x60\x20\xae\xbb\x18\xce\x1b\x86\x46\xc3\xed\xa5\xa3\x8b\xab\xd3\xab\xe3\xa0\x9c\xc4\x65\x9a\x56\xd9\xd4\x75\x83\xe7\xc4\xe1\x42\x39\xc1\x6b\x45\x03\xf2\xce\x4d\xa3\xb3\x58\x14\x48\x5e\xdd\xa2\x33\x6f\x23\xaa\x0b\x43\x06\xd0\x1d\x81\x26\xdf\x1d\x60\x3c\xcb\x2a\xfa\x1d\x57\xe1\x77\xac\xe3\x44\x1e\x9d\x7a\x2e\x7a\x61\xbc\x31\xa0\x45\x13\x20\x87\x45\x4e\x0b\x32\xf4\xce\x08\x52\x25\x36\x7b\xbb\xed\xe2\x4b\xf7\xea\xc3\x44\x1f\x96\xc6\x8d\xac\xe4\xad\x10\x51\xce\x94\xea\x5f\x6f\xdb\xd0\x58\x9b\xf3\x00\x09\x8d\x90\xfd\x5d\x4e\xc2\xb8\xee\xb3\xcd\x95\xc6\x0d\xb9\xba\xba\xeb\xe6\x1e\xe3\xae\xe6\xa4\x0c\xb9\xa3\xc5\xe9\xa4\xb5\x94\xde\x85\xbf\x8b\x6b\xc3\xab\xa2\xeb\xe7\xa3\xf7\x10\x88\x36\xa6\x2b\xdd\x32\xa5\xc9\x9e\x1e\xf0\x4d\xa2\xd2\xb2\x0c\x35\xc4\x25\xf3\x06\x11\x5f\xcb\x89\x98\x54\xef\x06\x7e\xaf\x56\x85\xaf\x1d\x61\x60\x08\x61\x50\x57\xa3\x6a\xc4\x28\x07\x7d\x18\x80\xc4\xf7\xd4\x34\xd7\x26\x73\x07\xe6\x0e\x1c\x7b\xa2\x59\x01\x61\x8e\xf0\x98\x1c\x44\x09\xd3\x81\x76\x23\xc3\x28\xa1\xf9\xae\x62\x54\xb7\x93\xd8\x3f\x6e\xeb\x13\x6d\x6c\x2d\x33\x7c\xa1\x59\x5f\xa2\x1b\x43\xc2\x51\xf5\xbc\x99\xd8\xf2\xb4\x9a\xfe\x50\x15\x09\x65\x01\x55\xa2\xc1\xd5\xf4\x44\x41\x94\xe6\xc0\x84\x42\xd7\x60\xd8\x13\xae\x21\xa9\x31\x89\xf7\x60\xeb\xfb\x98\xdd\x57\x22\x35\xcf\xd3\x97\x01\xfd\x6f\x52\x7e\xa2\xa9\x12\xd5\xd9\x29\xd4\x20\x2c\x27\x3b\x09\x3e\xb4\x65\x21\xb4\xc7\xd0\x3e\x87\x78\x7d\xf2\xd9\x37\xde\xb7\x1c\x73\xd7\xa7\x5e\xd3\xbc\xf0\x92\x45\x4c\x3d\x82\x37\x7b\xe8\x20\x7a\xfa\x08\xdf\xd4\x84\xbf\xbc\xf1\xd1\xce\x1b\xda\xff\x06\x00\x00\xff\xff\xf0\x1b\x32\x3e\xbc\x2c\x00\x00")

func skeleton_css_bytes() ([]byte, error) {
	return bindata_read(
		_skeleton_css,
		"skeleton.css",
	)
}

func skeleton_css() (*asset, error) {
	bytes, err := skeleton_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "skeleton.css", size: 11452, mode: os.FileMode(420), modTime: time.Unix(1419877841, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"app.css": app_css,
	"app.js": app_js,
	"index.html": index_html,
	"normalize.css": normalize_css,
	"skeleton.css": skeleton_css,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"app.css": &_bintree_t{app_css, map[string]*_bintree_t{
	}},
	"app.js": &_bintree_t{app_js, map[string]*_bintree_t{
	}},
	"index.html": &_bintree_t{index_html, map[string]*_bintree_t{
	}},
	"normalize.css": &_bintree_t{normalize_css, map[string]*_bintree_t{
	}},
	"skeleton.css": &_bintree_t{skeleton_css, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

