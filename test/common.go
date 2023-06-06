package hltv_test

import (
	"context"
	"os"
)

func mockGetPage(ctx context.Context, url string, body *string) error {
	tmp, err := os.ReadFile("testdata/html/get_team_map_stats.html")
	if err != nil {
		return err
	}

	*body = string(tmp)

	return nil
}
