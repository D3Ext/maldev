package antiforensics

/*

References:
https://github.com/lemon-mint/wipe-file-go/blob/main/wiper/wiper.go
https://www.kroll.com/en/insights/publications/cyber/anti-forensic-tactics/anti-forensics-tactics-timestomping

*/

import (
  "os"
  "math"
  "io"
  "time"
  "path/filepath"
  "encoding/base32"
  mrand "math/rand"
  "crypto/rand"
)

const blockSize = 1024 * 1024

func WipeFile(filename string) error {
  randbuf := make([]byte, blockSize)
  f, err := os.OpenFile(filename, os.O_RDWR, 0666)
  if err != nil {
    return err
  }
  stat, err := f.Stat()
  if err != nil {
    return err
  }
	
  size := stat.Size()
	blockCount := int(math.Ceil(float64(size)/blockSize)) + 1
	for pass := 0; pass < 3; pass++ {
		for i := 0; i < blockCount; i++ {
			f.WriteAt(randbuf, int64(i*blockSize))
			if pass != 0 {
				io.ReadFull(rand.Reader, randbuf)
			}
		}
		io.ReadFull(rand.Reader, randbuf)
		f.Sync()
	}
	f.Close()
	
  dir, _ := filepath.Split(filename)
	newname := filepath.Join(dir, randb32(max(len(filename), 20)))
	
  for i := 0; i < 10; i++ {
		os.Rename(filename, newname)
    filename = newname
    newname = filepath.Join(dir, randb32(max(len(filename), 20)))
	}

	return os.Remove(filename)
}

func TimestompFile(filename string, count int) error {
  for i := 0; i < count; i++ {
    err := os.Chtimes(filename, time.Unix(mrand.Int63n(time.Now().Unix()), 0), time.Unix(mrand.Int63n(time.Now().Unix()), 0))
    if err != nil {
      return err
    }
  }
  return nil
}

func randb32(size int) string {
  buf := make([]byte, size)
  io.ReadFull(rand.Reader, buf)
  return base32.StdEncoding.EncodeToString(buf)
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

