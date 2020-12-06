module go-img-processing

go 1.15

require github.com/h2non/bimg v1.1.5 // indirect

require locals/utils v0.0.0

require locals/imgprocess v0.0.0

// locals/types => ./types
replace locals/utils => ./utils

replace locals/imgprocess => ./imgprocess
