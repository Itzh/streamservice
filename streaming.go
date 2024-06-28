package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func stream(fstream *os.File, ctx *gin.Context) {
	mp3_320kbpsframe := 320_000 / 8
	stat, err := fstream.Stat()
	if err != nil {
		panic(err)
	}
	seek_sec, e := strconv.Atoi(ctx.Query("seek"))
	if e != nil {
		seek_sec = 0
	}

	ofs := stat.Size() - int64(seek_sec)*int64(mp3_320kbpsframe)
	fmt.Printf("Ofset: %d / %d", ofs, stat.Size())

	header := ctx.Writer.Header()
	header.Set("Content-Type", "application/octet-stream")
	header.Set("Content-Length", strconv.FormatInt(ofs, 10))
	header.Set("Content-Disposition", "inline; filename="+stat.Name()+"")

	httpWr := ctx.Writer
	httpWr.WriteHeader(http.StatusOK)
	for i := int64(seek_sec * mp3_320kbpsframe); i < stat.Size(); {
		bufLen := mp3_320kbpsframe
		if stat.Size()-i < int64(bufLen) {
			bufLen = int(stat.Size() - i)
		}
		buf := make([]byte, bufLen)
		rByte, e := fstream.ReadAt(buf, i)
		if e != nil {
			httpWr.Flush()
			panic(e)
		}

		wByte, e := httpWr.Write(buf)
		if e != nil {
			httpWr.Flush()
			panic(e)
		}
		i += int64(wByte)
		fmt.Printf("%d - %d / %d, Total %d \n", i, rByte, wByte, stat.Size())
		httpWr.(http.Flusher).Flush()
		time.Sleep(time.Duration(1) * time.Millisecond)
	}
	httpWr.(http.Flusher).Flush()
}
