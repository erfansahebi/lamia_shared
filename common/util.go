package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

type ResponseStruct struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}

type ContextKey int

const (
	ContextUser ContextKey = iota
)

var (
	ErrUnsupportedMIMEType = errors.New("unsupported MIME type")
)

// TODO: handle this with UUID
//func GetUserIDFromContext(ctx context.Context) uint64 {
//	userData := ctx.Value(ContextUser)
//	if userData == nil {
//		return 0
//	}
//
//	userID, _ := UserIDFromClaims(userData.(map[string]interface{}))
//
//	return uint64(userID)
//}

func Response(writer http.ResponseWriter, detail ResponseStruct, code int) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(code)

	json.NewEncoder(writer).Encode(detail)
}

func HTTPErr(w http.ResponseWriter, statusCode int, errs ...error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	detail := map[string]interface{}{
		"error": errs[0].Error(),
	}

	if len(errs) > 1 {
		detail["errors"] = errs[1:]
	}

	json.NewEncoder(w).Encode(detail)
}

type ctxDebugKey int

func SetDebug(ctx context.Context, debug bool) context.Context {
	return context.WithValue(ctx, ctxDebugKey(0), debug)
}

func IsDebug(ctx context.Context) bool {
	val, ok := ctx.Value(ctxDebugKey(0)).(bool)
	return ok && val
}

func UploadFile(r *http.Request, fileKey string) (tempPath string, mimeType string, err error) {
	file, _, err := r.FormFile(fileKey)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	dst, err := os.CreateTemp("", "upload-*")
	if err != nil {
		return "", "", nil
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return "", "", err
	}

	_, err = dst.Seek(0, 0)
	if err != nil {
		return "", "", err
	}

	incipit := make([]byte, 512)
	_, err = io.ReadAtLeast(dst, incipit, 512)
	if err != nil {
		return "", "", err
	}

	//mimeType = mimeFromIncipit(incipit)
	mediaMimeType := mimetype.Detect(incipit)
	if mediaMimeType == nil {
		return "", "", ErrUnsupportedMIMEType
	}

	tempPath, err = filepath.Abs(dst.Name())
	if err != nil {
		return "", "", err
	}

	return tempPath, mediaMimeType.String(), nil
}

// image formats and magic numbers
var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
}

func mimeFromIncipit(incipit []byte) string {
	incipitStr := string(incipit)
	for magic, mime := range magicTable {
		if strings.HasPrefix(incipitStr, magic) {
			return mime
		}
	}

	return ""
}

func ContainsUint64(array []uint64, needle uint64) bool {
	for _, item := range array {
		if item == needle {
			return true
		}
	}

	return false
}

func HasPrefix(mainString string, prefixes []string) (bool, string, error) {
	for _, prefix := range prefixes {
		regxString, err := regexp.Compile(fmt.Sprintf("^%s", prefix))
		if err != nil {
			return false, "", err
		}

		if regxString.MatchString(mainString) {
			return true, prefix, nil
		}
	}

	return false, "", nil
}
