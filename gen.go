// SPDX-License-Identifier: Apache-2.0

// +build ignore

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const notoEmojiVersion = "v2020-09-16-unicode13_1"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	if err := clean(); err != nil {
		return err
	}

	tmp, err := os.MkdirTemp("", "noto-emoji-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	if err := prepareNotoEmojiFiles(tmp); err != nil {
		return err
	}

	return nil
}

func clean() error {
	fmt.Printf("Cleaning Noto Emoji files\n")
	if err := os.RemoveAll("png"); err != nil {
		return err
	}
	if err := os.Remove("README-noto-emoji.md"); err != nil {
		return err
	}
	if err := os.Remove("LICENSE-noto-emoji"); err != nil {
		return err
	}
	return nil
}

func prepareNotoEmojiFiles(tmp string) error {
	fn := notoEmojiVersion + ".tar.gz"
	if e, err := exists(fn); err != nil {
		return err
	} else if !e {
		url := "https://github.com/googlefonts/noto-emoji/archive/refs/tags/" + fn
		fmt.Fprintf(os.Stderr, "%s not found: please download it from %s\n", fn, url)
		return nil
	}

	fmt.Printf("Copying %s to %s\n", fn, filepath.Join(tmp, fn))
	in, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(filepath.Join(tmp, fn))
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	fmt.Printf("Extracting %s\n", fn)
	cmd := exec.Command("tar", "-xzf", fn)
	cmd.Stderr = os.Stderr
	cmd.Dir = tmp
	if err := cmd.Run(); err != nil {
		return err
	}

	dir := filepath.Join(tmp, "noto-emoji-"+notoEmojiVersion[1:])
	fmt.Printf("Copying PNG images\n")
	srcpngdir := filepath.Join(dir, "png", "128")
	dstpngdir := "image"
	if err := os.MkdirAll(dstpngdir, 0755); err != nil {
		return err
	}
	es, err := os.ReadDir(srcpngdir)
	if err != nil {
		return err
	}
	for _, e := range es {
		if e.IsDir() {
			continue
		}
		src := filepath.Join(srcpngdir, e.Name())
		dst := filepath.Join(dstpngdir, e.Name())
		if err := os.Rename(src, dst); err != nil {
			return err
		}
	}

	fmt.Printf("Copying README.md and LICENSE\n")
	for _, f := range []string{"README.md", "LICENSE"} {
		infn := filepath.Join(dir, f)

		ext := filepath.Ext(f)
		outfn := f[:len(f)-len(ext)] + "-noto-emoji" + ext

		in, err := os.Open(infn)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.Create(outfn)
		if err != nil {
			return err
		}
		defer out.Close()

		if _, err := io.Copy(out, in); err != nil {
			return err
		}
	}

	return nil
}

func exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
