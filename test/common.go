package hltv_test

import (
	"context"
	"fmt"
	"os"
)

func mockGetPage(ctx context.Context, body *string, htmlName string) error {
	fileloc := fmt.Sprintf("testdata/html/%v", htmlName)
	tmp, err := os.ReadFile(fileloc)
	if err != nil {
		return err
	}

	*body = string(tmp)

	return nil
}
