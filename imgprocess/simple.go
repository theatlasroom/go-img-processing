package imgprocess

import (
	"github.com/h2non/bimg"
)

// Size extracts the size of the image
func Size(src *bimg.Image) (bimg.ImageSize, error) {
	return bimg.Size(src.Image())
}

// // Resizes the image at path `path` and returns the buffer
// // TODO: probably better to take a pointer to the image
// func Resize(src bimg.Image, w, h int) {
// 	buffer, err := bimg.Read(src)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}

// 	newImage, err := bimg.NewImage(buffer).ForceResize(w, h)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}

// 	size := bimg.Size(newImage)
// 	if size.Width != w || size.Height != h {
// 		fmt.Fprintln(os.Stderr, "Incorrect image size")
// 	}
// }
