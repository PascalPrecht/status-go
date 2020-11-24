// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 0001_accounts.down.sql (21B)
// 0001_accounts.up.sql (163B)
// 1605007189_identity_images.down.sql (29B)
// 1605007189_identity_images.up.sql (268B)
// 1606224181_drop_photo_path_from_accounts.down.sql (47B)
// 1606224181_drop_photo_path_from_accounts.up.sql (675B)
// doc.go (74B)

package migrations

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

var __0001_accountsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\x96\x1e\x13\xa1\x15\x00\x00\x00")

func _0001_accountsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_accountsDownSql,
		"0001_accounts.down.sql",
	)
}

func _0001_accountsDownSql() (*asset, error) {
	bytes, err := _0001_accountsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_accounts.down.sql", size: 21, mode: os.FileMode(0644), modTime: time.Unix(1585751575, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd2, 0x61, 0x4c, 0x18, 0xfc, 0xc, 0xdf, 0x5c, 0x1f, 0x5e, 0xd3, 0xbd, 0xfa, 0x12, 0x5e, 0x8d, 0x8d, 0x8b, 0xb9, 0x5f, 0x99, 0x46, 0x63, 0xa5, 0xe3, 0xa6, 0x8a, 0x4, 0xf1, 0x73, 0x8a, 0xe9}}
	return a, nil
}

var __0001_accountsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcc\xb1\x6e\x83\x30\x14\x46\xe1\xdd\x4f\xf1\x8f\xad\xe4\x37\xe8\x64\xa8\x5b\xae\x42\x00\x99\x4b\x80\xd1\x02\x04\x56\x82\x8d\xc0\x19\xf2\xf6\x51\x58\x8f\xf4\x9d\xd4\x68\xc5\x1a\xac\x92\x5c\x83\xfe\x50\x94\x0c\xdd\x51\xcd\x35\xec\x30\x84\xa7\x8f\x07\xbe\xc4\x7d\x7a\x35\x6e\xc4\x4d\x99\x34\x53\x06\x95\xa1\xab\x32\x3d\x2e\xba\x97\xc2\xdb\x75\x02\xeb\x8e\x4f\x5b\x34\x79\x2e\xc5\x23\xcc\xce\xb3\x5b\xa7\x23\xda\x75\x43\x42\xff\xa0\x82\xa5\xd8\x96\x10\x43\x65\xe3\x72\x02\xf9\xf9\x0e\x76\x1f\x2b\xeb\x76\xe7\xe7\x33\x8a\x6f\xb4\xc4\x59\xd9\x30\x4c\xd9\xd2\xef\x8f\x78\x07\x00\x00\xff\xff\xab\xcf\xa2\xbd\xa3\x00\x00\x00")

func _0001_accountsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_accountsUpSql,
		"0001_accounts.up.sql",
	)
}

func _0001_accountsUpSql() (*asset, error) {
	bytes, err := _0001_accountsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_accounts.up.sql", size: 163, mode: os.FileMode(0644), modTime: time.Unix(1585751575, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf2, 0xfa, 0x99, 0x8e, 0x96, 0xb3, 0x13, 0x6c, 0x1f, 0x6, 0x27, 0xc5, 0xd2, 0xd4, 0xe0, 0xa5, 0x26, 0x82, 0xa7, 0x26, 0xf2, 0x68, 0x9d, 0xed, 0x9c, 0x3d, 0xbb, 0xdc, 0x37, 0x28, 0xbc, 0x1}}
	return a, nil
}

var __1605007189_identity_imagesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\x4c\x49\xcd\x2b\xc9\x2c\xa9\x8c\xcf\xcc\x4d\x4c\x4f\x2d\xb6\xe6\xe5\x02\x04\x00\x00\xff\xff\xa1\x22\x72\x37\x1d\x00\x00\x00")

func _1605007189_identity_imagesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1605007189_identity_imagesDownSql,
		"1605007189_identity_images.down.sql",
	)
}

func _1605007189_identity_imagesDownSql() (*asset, error) {
	bytes, err := _1605007189_identity_imagesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1605007189_identity_images.down.sql", size: 29, mode: os.FileMode(0644), modTime: time.Unix(1604935057, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2f, 0xcf, 0xa7, 0xae, 0xd5, 0x4f, 0xcd, 0x14, 0x63, 0x9, 0xbe, 0x39, 0x49, 0x18, 0x96, 0xb2, 0xa3, 0x8, 0x7d, 0x41, 0xdb, 0x50, 0x5d, 0xf5, 0x4d, 0xa2, 0xd, 0x8f, 0x57, 0x79, 0x77, 0x67}}
	return a, nil
}

var __1605007189_identity_imagesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xc1\x6a\xc3\x30\x10\x04\xd0\xbb\xc1\xff\x30\xc7\x04\xf2\x07\x3d\xc9\xaa\x42\x44\x55\x29\x28\x4a\xd3\x9c\x84\x40\x5b\x7b\x69\xe2\x96\x58\xa5\xb8\x5f\x5f\xea\xfa\x60\x72\xdc\xc7\xec\x30\xd2\x2b\x11\x14\x82\x68\x8c\x82\xde\xc2\xba\x00\xf5\xaa\x0f\xe1\x00\xce\xd4\x17\x2e\x63\xe4\x6b\x6a\x69\x58\xd5\x15\x00\xbc\xd3\x18\xbf\x38\xe3\x45\x78\xb9\x13\x7e\xf3\xaf\x7d\xba\xd2\x1d\x4d\x5f\xf1\x33\x8d\x97\x8f\x94\xd1\x18\xd7\x4c\xe5\xf6\x68\xcc\x9c\xf8\xe6\x5c\x3a\x70\x5f\xe6\xbb\x23\x6e\xbb\xb2\x80\x37\xbe\x50\x1c\xf8\x87\x16\x76\xa3\x3f\x88\x25\xdd\x5a\x5a\x66\xf7\x5e\x3f\x0b\x7f\xc6\x93\x3a\x63\x35\x8f\xdc\x4c\xbb\xd6\x70\x16\xd2\xd9\xad\xd1\x32\xc0\xab\xbd\x11\x52\xd5\xd5\x1a\x27\x1d\x76\xee\x18\xe0\xdd\x49\x3f\x3e\xd4\xd5\x6f\x00\x00\x00\xff\xff\x8c\x6a\x0a\x57\x0c\x01\x00\x00")

func _1605007189_identity_imagesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1605007189_identity_imagesUpSql,
		"1605007189_identity_images.up.sql",
	)
}

func _1605007189_identity_imagesUpSql() (*asset, error) {
	bytes, err := _1605007189_identity_imagesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1605007189_identity_images.up.sql", size: 268, mode: os.FileMode(0644), modTime: time.Unix(1605999808, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x50, 0xb6, 0xc1, 0x5c, 0x76, 0x72, 0x6b, 0x22, 0x34, 0xdc, 0x96, 0xdc, 0x2b, 0xfd, 0x2d, 0xbe, 0xcc, 0x1e, 0xd4, 0x5, 0x93, 0xd, 0xc2, 0x51, 0xf3, 0x1a, 0xef, 0x2b, 0x26, 0xa4, 0xeb, 0x65}}
	return a, nil
}

var __1606224181_drop_photo_path_from_accountsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xc8\xc8\x2f\xc9\x0f\x48\x2c\xc9\x50\x08\x71\x8d\x08\xb1\x06\x04\x00\x00\xff\xff\xc2\xe6\xd6\xb1\x2f\x00\x00\x00")

func _1606224181_drop_photo_path_from_accountsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1606224181_drop_photo_path_from_accountsDownSql,
		"1606224181_drop_photo_path_from_accounts.down.sql",
	)
}

func _1606224181_drop_photo_path_from_accountsDownSql() (*asset, error) {
	bytes, err := _1606224181_drop_photo_path_from_accountsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1606224181_drop_photo_path_from_accounts.down.sql", size: 47, mode: os.FileMode(0644), modTime: time.Unix(1606258709, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdc, 0x18, 0x37, 0x1a, 0x2, 0x2b, 0x46, 0x72, 0xe0, 0x25, 0xa6, 0x7c, 0x59, 0xa4, 0xde, 0x3a, 0x78, 0x30, 0x9a, 0xc6, 0xbf, 0x24, 0x96, 0xa6, 0x4, 0xa4, 0x73, 0x92, 0xc6, 0x18, 0xe4, 0xfc}}
	return a, nil
}

var __1606224181_drop_photo_path_from_accountsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd2\xc1\xae\xa2\x30\x14\x06\xe0\x3d\x4f\xf1\x2f\x95\x90\xf0\x00\xae\x10\x3a\x4a\x06\x81\xd4\x32\xea\x6a\xac\xd0\x48\x23\xb4\x0d\x96\x4c\x7c\xfb\x89\x88\xde\x48\xee\x5d\xdc\xdc\xed\x81\xff\x9c\x9f\x2f\xf8\x2e\x42\x6d\x6e\xb0\xb5\x00\x2f\x4b\xdd\x2b\x7b\x85\xe5\xa7\x46\x40\x2a\xab\xc1\x61\x45\x6b\x1e\x13\x0f\x64\x1f\x26\x45\x44\x86\xb7\x8f\xa6\xd6\x56\xe7\xdc\xd6\x47\x94\xba\xe9\x5b\x05\xd7\x77\x42\x4a\x02\x46\xc0\xc8\x26\xcf\x68\x40\x0f\x60\xc1\x32\x21\xaf\xd5\x7f\x4f\xbc\xbc\xf4\x66\xe6\x00\xc0\x45\xdc\x0a\x59\xe1\x4f\x40\xc3\x75\x40\x91\xd3\x78\x73\x4f\xfc\x26\x07\x6f\x78\xae\x78\x2b\xc0\xc8\x9e\x21\xcd\x18\xd2\x22\x49\x1e\xf3\x46\x9f\xa5\x62\xb2\x15\x57\xcb\x5b\x83\x65\xbc\x42\x9c\x32\xef\xb9\xb3\xe4\x5d\x95\x73\xd9\x49\x75\x1e\xd2\xce\x1c\xbb\x98\xad\xb3\x82\x81\x66\xbb\x38\x5a\x38\x71\xba\x25\x94\xdd\x43\xd9\xb4\x19\xb6\x24\x21\x21\x1b\xbb\x79\x43\x07\x6f\x72\xd1\x9b\x5e\xf9\x45\xb3\xcd\x6b\xd1\xc2\x71\x7c\x17\x51\xa7\xcd\xc0\xa4\x9b\x6a\x0a\xcb\x55\x85\x4e\x94\x9d\xe0\x56\xe0\x9f\xb4\x35\x78\xd3\x8c\x86\xd7\xd1\x38\x4e\x57\x6f\xc2\xae\xef\x44\x34\xcb\x27\x9c\x8b\x97\xf7\xdb\x78\xf6\x9d\xf6\xf3\xcf\x3d\x7e\x0e\x31\x8a\x3e\x3c\x98\xac\x6e\xe8\x8d\x87\xea\x09\xf3\xf1\x5f\x7d\xf1\x71\xcf\xfc\xff\x00\x00\x00\xff\xff\x50\xd5\xcd\xce\xa3\x02\x00\x00")

func _1606224181_drop_photo_path_from_accountsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1606224181_drop_photo_path_from_accountsUpSql,
		"1606224181_drop_photo_path_from_accounts.up.sql",
	)
}

func _1606224181_drop_photo_path_from_accountsUpSql() (*asset, error) {
	bytes, err := _1606224181_drop_photo_path_from_accountsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1606224181_drop_photo_path_from_accounts.up.sql", size: 675, mode: os.FileMode(0644), modTime: time.Unix(1606225159, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa7, 0x3a, 0x8e, 0xc7, 0x1a, 0xfc, 0x95, 0x2c, 0x3, 0xf4, 0x6b, 0xd0, 0x1f, 0xb6, 0xf9, 0x48, 0x99, 0xef, 0xcc, 0x97, 0xc4, 0x63, 0x94, 0x4, 0x82, 0xb8, 0xf5, 0xc2, 0xe0, 0x3c, 0x3e, 0x18}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x0d\xc4\x20\x0c\x05\xd0\x9e\x29\xfe\x02\xd8\xfd\x6d\xe3\x4b\xac\x2f\x44\x82\x09\x78\x7f\xa5\x49\xfd\xa6\x1d\xdd\xe8\xd8\xcf\x55\x8a\x2a\xe3\x47\x1f\xbe\x2c\x1d\x8c\xfa\x6f\xe3\xb4\x34\xd4\xd9\x89\xbb\x71\x59\xb6\x18\x1b\x35\x20\xa2\x9f\x0a\x03\xa2\xe5\x0d\x00\x00\xff\xff\x60\xcd\x06\xbe\x4a\x00\x00\x00")

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

	info := bindataFileInfo{name: "doc.go", size: 74, mode: os.FileMode(0644), modTime: time.Unix(1585751575, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xde, 0x7c, 0x28, 0xcd, 0x47, 0xf2, 0xfa, 0x7c, 0x51, 0x2d, 0xd8, 0x38, 0xb, 0xb0, 0x34, 0x9d, 0x4c, 0x62, 0xa, 0x9e, 0x28, 0xc3, 0x31, 0x23, 0xd9, 0xbb, 0x89, 0x9f, 0xa0, 0x89, 0x1f, 0xe8}}
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
	"0001_accounts.down.sql": _0001_accountsDownSql,

	"0001_accounts.up.sql": _0001_accountsUpSql,

	"1605007189_identity_images.down.sql": _1605007189_identity_imagesDownSql,

	"1605007189_identity_images.up.sql": _1605007189_identity_imagesUpSql,

	"1606224181_drop_photo_path_from_accounts.down.sql": _1606224181_drop_photo_path_from_accountsDownSql,

	"1606224181_drop_photo_path_from_accounts.up.sql": _1606224181_drop_photo_path_from_accountsUpSql,

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
	"0001_accounts.down.sql":                            &bintree{_0001_accountsDownSql, map[string]*bintree{}},
	"0001_accounts.up.sql":                              &bintree{_0001_accountsUpSql, map[string]*bintree{}},
	"1605007189_identity_images.down.sql":               &bintree{_1605007189_identity_imagesDownSql, map[string]*bintree{}},
	"1605007189_identity_images.up.sql":                 &bintree{_1605007189_identity_imagesUpSql, map[string]*bintree{}},
	"1606224181_drop_photo_path_from_accounts.down.sql": &bintree{_1606224181_drop_photo_path_from_accountsDownSql, map[string]*bintree{}},
	"1606224181_drop_photo_path_from_accounts.up.sql":   &bintree{_1606224181_drop_photo_path_from_accountsUpSql, map[string]*bintree{}},
	"doc.go": &bintree{docGo, map[string]*bintree{}},
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
