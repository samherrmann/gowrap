package zip

import (
	"archive/zip"
	"io"
	"os"
)

// Make packages the files referenced by 'filePaths' into
// a '.zip' file.
func Make(archivePath string, filePaths []string) error {
	// set up the output archive file
	file, err := os.Create(archivePath + ".zip")
	if err != nil {
		return err
	}
	defer file.Close()

	// set up the zip writer
	zw := zip.NewWriter(file)
	defer zw.Close()

	// add files to zip archive.
	for _, path := range filePaths {
		if err := addFile(zw, path); err != nil {
			return err
		}
	}
	return nil
}

// addFile adds a file to the zip archive.
func addFile(zw *zip.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// write header to zip archive.
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Method = zip.Deflate
	w, err := zw.CreateHeader(header)
	if err != nil {
		return err
	}

	// copy file data to zip archive.
	_, err = io.Copy(w, file)
	return err
}
