package tag

import "errors"

type TagType string

const (
	TAG_DEV   TagType = "dev"
	TAG_RC    TagType = "rc"
	TAG_PATCH TagType = "patch"
	TAG_MINOR TagType = "minor"
	TAG_MAJOR TagType = "major"
)

func (t TagType) IsValid() error {
	switch t {
	case TAG_DEV, TAG_RC, TAG_PATCH, TAG_MINOR, TAG_MAJOR:
		return nil
	}
	return errors.New("an invalid or unsupported tag type was provided")
}
