// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	This program will take a NOAA radar image and remove the noise colors.

	// Install Binary Package C Library Procedure
	mkdir ~/temp
	cd ~/temp
	curl -O http://www.imagemagick.org/download/ImageMagick.tar.gz
	tar -xzf ImageMagick.tar.gz
	rm -f ImageMagick.tar.gz
	cd ImageMagick-6.8.6-9/
	./configure
	make
	sudo make install
	sudo ldconfig /usr/local/lib   ** LINUX ONLY
	remove ImageMagick-6.8.6-9
	pkg-config --cflags --libs MagickWand

	-- For Development Environment Only
	export GOPATH=$HOME/<My New Folder Location>
	go get github.com/gographics/imagick/imagick

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

	colors := []string{
		"#3030CE",
		"#04e9e7",
		"#019ff4",
		"#0300f4",
		"#a9a879",
		"#777777",
		"#7a4679",
		"#aa7ca9",
		"#d7acd6",
		"#cccc99",
		"#999966",
		"#646464",
		"#663366",
	}

	for _, color := range colors {

		pixelWand := imagick.NewPixelWand()
		pixelWand.SetColor(color)
		mw.TransparentPaintImage(pixelWand, 0, fuzz, false)
		defer pixelWand.Destroy()
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
