package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// ディレクトリを再帰的にコピーする関数
func CopyDir(srcDir, dstDir string) error {
	return filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// コピー先のパスを決定
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)

		if d.IsDir() {
			// ディレクトリなら作成（なければ）
			return os.MkdirAll(dstPath, os.ModePerm)
		}

		// ファイルならコピー
		return copyFile(path, dstPath)
	})
}

// 1ファイルをコピーするヘルパー関数（上書き対応）
func copyFile(srcFile, dstFile string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dstFile) // Createは既にあれば上書き
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	// パーミッションもコピー
	fi, err := os.Stat(srcFile)
	if err != nil {
		return err
	}
	return os.Chmod(dstFile, fi.Mode())
}
