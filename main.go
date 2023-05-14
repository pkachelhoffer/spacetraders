package main

import (
	"context"
	"fmt"
	"spacetraders/stclient"
)

func main() {
	ctx := context.Background()
	client := stclient.NewAPIClient(stclient.NewConfiguration())

	resp, httpResp, err := client.SystemsApi.GetSystemExecute(client.SystemsApi.GetSystem(ctx, "X1-AC10"))
	fmt.Print(resp, httpResp, err)
}
