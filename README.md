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

	-- For Development Environment Only
	export GOPATH=$HOME/<My New Folder Location>
	go get github.com/gographics/imagick/imagick

	-- Make sure these environment variables are set
	MAGICK_HOME=$HOME/Spaces/PublicPackages/ImageMagick-6.8.6
	DYLD_LIBRARY_PATH=$MAGICK_HOME/lib/
	PKG_CONFIG_PATH=$HOME/Spaces/PublicPackages/ImageMagick-6.8.6/lib/pkgconfig

	cleanradarimage source.gif out.gif

