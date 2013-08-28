/*
	// Make sure these environment variables are set
	MAGICK_HOME=$HOME/Spaces/PublicPackages/ImageMagick-6.8.6
	DYLD_LIBRARY_PATH=$MAGICK_HOME/lib/
	PKG_CONFIG_PATH=$HOME/Spaces/PublicPackages/ImageMagick-6.8.6/lib/pkgconfig

	cleanradarimage source.gif out.gif
*/
package main

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
	"log"
	"os"
)

// CleanImage removes noise from the image
func main() {

	if len(os.Args) != 3 {

		fmt.Println("cleanradarimage source.gif out.gif")
		return
	}

	if FileExists(os.Args[2]) == true {

		os.Remove(os.Args[2])
	}

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImage(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	fuzz := float64(10) // should be 10%

	// our special colors we want to filter
	filter := []*imagick.PixelWand{
		imagick.NewPixelWand(),
		imagick.NewPixelWand(),
		imagick.NewPixelWand(),
		imagick.NewPixelWand(),
	}

	filter[0].SetColor("#3030CE")
	defer filter[0].Destroy()
	filter[1].SetColor("#04e9e7")
	defer filter[1].Destroy()
	filter[2].SetColor("#019ff4")
	defer filter[2].Destroy()
	filter[3].SetColor("#0300f4")
	defer filter[3].Destroy()

	for _, f := range filter {

		mw.TransparentPaintImage(f, 0, fuzz, false)
	}

	mw.BlurImage(2, 2)

	mw.WriteImage(os.Args[2])
}

func FileExists(path string) bool {

	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	return false
}
