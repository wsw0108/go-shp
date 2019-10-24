package shp

type option struct {
	dbfEncoding string
}

// OptionFunc used to set options
type OptionFunc func(*option)

// WithDbfEncoding set the encoding of dbf file
func WithDbfEncoding(enc string) OptionFunc {
	return func(opts *option) {
		opts.dbfEncoding = enc
	}
}
