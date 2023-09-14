package main

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
)

const (
	swaggerUIVersion = "5.7.0"
)

var (
	distDirPattern = regexp.MustCompile(`^swagger-ui-\d+\.\d+\.\d+/dist/(.*)`)
	filesToIgnore  = []string{"swagger-initializer.js"}
	workingDir     string
)

func init() {
	if wd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		workingDir = wd
	}
}

func DownloadSwaggerUi(ctx context.Context) (err error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://github.com/swagger-api/swagger-ui/archive/refs/tags/v%s.tar.gz", swaggerUIVersion),
		nil,
	)

	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("requested failed: %s", resp.Status)
	}

	defer func() {
		err = errors.Join(err, resp.Body.Close())
	}()

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		if !distDirPattern.MatchString(header.Name) {
			continue
		}

		submatches := distDirPattern.FindStringSubmatch(header.Name)

		if len(submatches) != 2 || submatches[1] == "" {
			continue
		}

		if slices.Contains(filesToIgnore, submatches[0]) {
			continue
		}

		if err := os.MkdirAll(filepath.Join(workingDir, "assets", "swagger-ui"), 0o755); err != nil {
			return err
		}

		target := filepath.Join(workingDir, "assets", "swagger-ui", submatches[1])

		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tarReader); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}
