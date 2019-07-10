package objectmodel

// Parser - Interface for parsers, similar to io.Reader
type Parser interface {
	Parse(data []byte) (*ObjectModel, error)
}
