package gziptar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

// Make packages the files referenced by 'filePaths' into
// a '.tar.gz' file.
func Make(archivePath string, filePaths []string) error {
	// set up the output archive file
	file, err := os.Create(archivePath + ".tar.gz")
	if err != nil {
		return err
	}
	defer file.Close()

	// set up the gzip writer
	gw := gzip.NewWriter(file)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// add files to tarball
	for _, path := range filePaths {
		if err := addFile(tw, path); err != nil {
			return err
		}
	}
	return nil
}

// addFile adds a file to the tarball.
func addFile(tw *tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	// write header to tarball
	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	// copy file data to tarball
	_, err = io.Copy(tw, file)
	return err
}
