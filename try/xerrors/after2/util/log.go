package util

import "golang.org/x/xerrors"

func Log(org error) error {
	if org == nil {
		return nil
	}
	return xerrors.Errorf("[Log] error occurred: %w", org)
}
