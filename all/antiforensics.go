package all

import "github.com/D3Ext/maldev/antiforensics"

func Wipe(src string) (error) {
  return antiforensics.Wipe(src)
}

func Timestomp(src string, count int) (error) {
  return antiforensics.Timestomp(src, count)
}

