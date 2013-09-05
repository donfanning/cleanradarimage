# Clean Radar Image

Copyright 2013 Ardan Studios. All rights reserved.  
Use of this source code is governed by a BSD-style license that can be found in the LICENSE handle.

This program will take a NOAA radar image and remove the noise colors.

Ardan Studios  
12973 SW 112 ST, Suite 153  
Miami, FL 33186  
bill@ardanstudios.com

	-- Install Binary Package C Library Procedure
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

	-- Go Code
	export GOPATH=$HOME/goinggo
	go get github.com/gographics/imagick/imagick
	go get github.com/goinggo/cleanradarimage
	
	-- Run Program
	cd $GOPATH/bin
	./cleanradarimage source.gif out.gif