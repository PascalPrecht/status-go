// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1561059284_add_waku_keys.down.sql (22B)
// 1561059284_add_waku_keys.up.sql (109B)
// doc.go (373B)

package sqlite

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __1561059284_add_waku_keysDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x4f\xcc\x2e\x8d\xcf\x4e\xad\x2c\xb6\xe6\x02\x04\x00\x00\xff\xff\x4f\x00\xe6\x8e\x16\x00\x00\x00")

func _1561059284_add_waku_keysDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1561059284_add_waku_keysDownSql,
		"1561059284_add_waku_keys.down.sql",
	)
}

func _1561059284_add_waku_keysDownSql() (*asset, error) {
	bytes, err := _1561059284_add_waku_keysDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1561059284_add_waku_keys.down.sql", size: 22, mode: os.FileMode(0664), modTime: time.Unix(1604073469, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe5, 0x2a, 0x7e, 0x9, 0xa3, 0xdd, 0xc6, 0x3, 0xfa, 0xaa, 0x98, 0xa0, 0x26, 0x5e, 0x67, 0x43, 0xe6, 0x20, 0xfd, 0x10, 0xfd, 0x60, 0x89, 0x17, 0x13, 0x87, 0x1b, 0x44, 0x36, 0x79, 0xb6, 0x60}}
	return a, nil
}

var __1561059284_add_waku_keysUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xb1\x0a\xc2\x40\x0c\x06\xe0\xfd\x9e\xe2\x1f\x15\x7c\x03\xa7\xbb\x33\x6a\x30\x26\x12\x52\x6a\xa7\x52\xb4\xa0\xdc\xa8\x22\x7d\xfb\x7e\xd5\x29\x07\x21\x72\x11\xc2\x7f\x6a\xbf\xb1\xcd\xcb\x07\x9b\x04\x3c\x5e\xd3\x77\x7c\x3f\x11\x74\x0f\xdc\x9c\xaf\xd9\x07\x5c\x68\x80\x29\xaa\xe9\x51\xb8\x06\xf8\xa4\xe6\xb4\x4b\x40\x9b\x17\x14\xb1\x02\xb5\x80\x76\x22\x69\x8b\x9e\xe3\x6c\x5d\xc0\xad\xe7\xc3\x3e\xad\x01\x00\x00\xff\xff\xbc\x45\x31\x54\x6d\x00\x00\x00")

func _1561059284_add_waku_keysUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1561059284_add_waku_keysUpSql,
		"1561059284_add_waku_keys.up.sql",
	)
}

func _1561059284_add_waku_keysUpSql() (*asset, error) {
	bytes, err := _1561059284_add_waku_keysUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1561059284_add_waku_keys.up.sql", size: 109, mode: os.FileMode(0664), modTime: time.Unix(1604073469, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa9, 0x5c, 0x8, 0x32, 0xef, 0x12, 0x88, 0x21, 0xd, 0x7a, 0x42, 0x4d, 0xe7, 0x2d, 0x6c, 0x99, 0xb6, 0x1, 0xf1, 0xba, 0x2c, 0x40, 0x8d, 0xa9, 0x4b, 0xe6, 0xc4, 0x21, 0xec, 0x47, 0x6b, 0xf7}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x3d\x72\xeb\x30\x0c\x84\x7b\x9d\x62\xc7\x8d\x9b\x27\xb2\x79\x55\xba\x94\xe9\x73\x01\x98\x5a\x91\x18\x4b\xa4\x42\xc0\x7f\xb7\xcf\xc8\xe3\xc2\x5d\xda\x1d\x7c\x1f\x76\x63\xc4\x77\x51\xc3\xac\x0b\xa1\x86\xca\x44\x33\xe9\x0f\x9c\x98\xe4\x62\xc4\x21\xab\x97\xcb\x29\xa4\xb6\x46\x73\xf1\x8b\x8d\xba\xc6\x55\x73\x17\x67\xbc\xfe\x3f\x0c\x31\x22\x49\x3d\x3a\x8a\xd4\x69\xe1\xd3\x65\x30\x97\xee\x5a\x33\x6e\xea\x05\x82\xad\x73\xd6\x7b\xc0\xa7\x63\xa1\x98\xc3\x8b\xf8\xd1\xe0\x85\x48\x62\xdc\x35\x73\xeb\xc8\x6d\x3c\x69\x9d\xc4\x25\xec\xd1\xd7\xfc\x96\xec\x0d\x93\x2c\x0b\x27\xcc\xbd\xad\x4f\xd6\x64\x25\x26\xed\x4c\xde\xfa\xe3\x1f\xc4\x8c\x8e\x2a\x2b\x6d\xe7\x8b\x5c\x89\xda\x5e\xef\x21\x75\xfa\x7b\x11\x6e\xad\x9f\x0d\x62\xe0\x7d\x63\x72\x4e\x61\x18\x36\x49\x67\xc9\x84\xfd\x2c\xea\x1c\x86\x18\x73\xfb\xc8\xac\xdc\xa9\xf7\x8e\xe3\x76\xce\xaf\x2b\x8c\x0d\x21\xbc\xd4\xda\xaa\x85\xdc\x10\x86\xdf\x00\x00\x00\xff\xff\x21\xa5\x75\x05\x75\x01\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 373, mode: os.FileMode(0664), modTime: time.Unix(1604073469, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x23, 0x6a, 0xc1, 0xce, 0x94, 0xf6, 0xef, 0xf1, 0x97, 0x95, 0xb, 0x35, 0xaf, 0x5f, 0xe7, 0x5f, 0xac, 0x6e, 0xb8, 0xab, 0xba, 0xb5, 0x35, 0x97, 0x22, 0x36, 0x11, 0xce, 0x44, 0xfc, 0xfa, 0xac}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"1561059284_add_waku_keys.down.sql": _1561059284_add_waku_keysDownSql,

	"1561059284_add_waku_keys.up.sql": _1561059284_add_waku_keysUpSql,

	"doc.go": docGo,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"1561059284_add_waku_keys.down.sql": &bintree{_1561059284_add_waku_keysDownSql, map[string]*bintree{}},
	"1561059284_add_waku_keys.up.sql":   &bintree{_1561059284_add_waku_keysUpSql, map[string]*bintree{}},
	"doc.go":                            &bintree{docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
