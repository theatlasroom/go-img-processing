package types

type images []image.Images

type ActionFunction func(images)

type Action struct {
	Images  images
	Process ActionFunction
}
