package service

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"sync"
	"time"
)

type PicFile struct {
	dir      string // 图片目录
	filename string // 图片文件名
}

func (p *PicFile) ReadImage() ([][]uint32, string, error) {
	begin := time.Now()
	path := fmt.Sprintf("%v%v%v", p.dir, string(os.PathSeparator), p.filename)
	reader, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer reader.Close()
	img, t, err := image.Decode(reader)
	if err != nil {
		return nil, "", err
	}

	if t == "png" {
		reader.Seek(0, 0)
		img, _ = png.Decode(reader)
	} else {
		reader.Seek(0, 0)
		img, _ = jpeg.Decode(reader)
	}

	size := img.Bounds().Size()
	data := make([][]uint32, size.X)
	for i, _ := range data {
		data[i] = make([]uint32, size.Y)
	}
	wg := sync.WaitGroup{}
	wg.Add(size.X)
	for i := 0; i < size.X; i++ {
		go func(i int) {
			for j := 0; j < size.Y; j++ {
				r, g, b, a := img.At(i, j).RGBA()
				data[i][j] = a>>8 + (b>>8)*0x100 + (g>>8)*0x10000 + (r>>8)*0x1000000
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Printf("Read image: %v successfully, type:%v, cost: %.3fs", path, t, time.Since(begin).Seconds())
	return data, t, nil
}

// ExportFile Export image with specified type
func (p *PicFile) ExportFile(data [][]uint32, typ string) (string, error) {
	start := time.Now()
	x, y := len(data), len(data[0])
	paint := image.NewRGBA(image.Rect(0, 0, x, y))
	wg := sync.WaitGroup{}
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func(i int) {
			for j := 0; j < y; j++ {
				c := data[i][j]
				a := c & 0xff
				c = c >> 8
				b := c & 0xff
				c = c >> 8
				g := c & 0xff
				c = c >> 8
				r := c
				paint.SetRGBA(i, j, color.RGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(a),
				})
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	path := fmt.Sprintf("%v%vo_%v", p.dir, string(os.PathSeparator), p.filename)

	file, err := os.Create(path)
	if err != nil {
		return "", err
	}

	if typ == "png" {
		png.Encode(file, paint)
	} else if typ == "jpeg" || typ == "jpg" {
		jpeg.Encode(file, paint, &jpeg.Options{Quality: 100})
	}

	log.Printf("Successfully export file to %v, cost %.3fs", path, time.Since(start).Seconds())
	return path, nil
}
