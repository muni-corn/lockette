package brightness

import (
	"image"
    "math"
)

// ImageBrightness represents image brightness. It is either
// light or dark.
type ImageBrightness string

// ImageBrightness definitions
const (
	Light ImageBrightness = "light"
	Dark  ImageBrightness = "dark"
)

// ImageSection represents a section of the image.
type ImageSection int

// Image section definitions
const (
	All ImageSection = iota
    Center
	LeftHalf
	UpperHalf
	RightHalf
	LowerHalf
	UpperLeft
	UpperRight
	LowerRight
	LowerLeft
)

func getBounds(img image.Image, section ImageSection) image.Rectangle {
	// get img's width and height
	w, h := img.Bounds().Dx(), img.Bounds().Dy()

	switch section {
	case All:
		return img.Bounds()
    case Center:
		return image.Rect(w/3, h/3, w*2/3, h*2/3)
	case LeftHalf:
		return image.Rect(0, 0, w/2, h)
	case UpperHalf:
		return image.Rect(0, 0, w, h/2)
	case RightHalf:
		return image.Rect(w/2, 0, w, h)
	case LowerHalf:
		return image.Rect(0, h/2, w, h)
	case UpperLeft:
		return image.Rect(0, 0, w/2, h/2)
	case UpperRight:
		return image.Rect(w/2, 0, w, h/2)
	case LowerRight:
		return image.Rect(w/2, h/2, w, h)
	case LowerLeft:
		return image.Rect(0, h/2, w/2, h)
	}

    return img.Bounds()
}

// GetImageBrightness returns an ImageBrightness pertaining
// to the brightness of an image `img` and section `section`
func GetImageBrightness(img image.Image, section ImageSection) (classification ImageBrightness, brightness float32) {
    bounds := getBounds(img, section)

    var darkCount, lightCount int

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, _ := img.At(x, y).RGBA()
            max := math.Max(math.Max(float64(r), float64(g)), float64(b))
            if (max > 0x8888) {
                lightCount++;
            } else {
                darkCount++;
            }
        }
    }

    brightness = float32(lightCount)/float32(darkCount+lightCount)

    // if more light pixels than dark pixels, return light
    if (lightCount > darkCount) {
        classification = Light
    }

    // else, return Dark
    classification = Dark

    return
}
