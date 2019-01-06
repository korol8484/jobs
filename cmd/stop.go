// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"github.com/spf13/cobra"
	rr "github.com/spiral/roadrunner/cmd/rr/cmd"
	"github.com/spiral/roadrunner/cmd/util"
)

func init() {
	rr.CLI.AddCommand(&cobra.Command{
		Use:   "jobs:stop",
		Short: "Destroy job consuming for Job service brokers",
		RunE:  stopHandler,
	})
}

func stopHandler(cmd *cobra.Command, args []string) error {
	client, err := util.RPCClient(rr.Container)
	if err != nil {
		return err
	}
	defer client.Close()

	if len(args) == 0 {
		util.Printf("<yellow>stop job consumption for all pipelines</reset>: ")

		var r string
		if err := client.Call("jobs.StopAll", true, &r); err != nil {
			return err
		}

		util.Printf("<green+hb>done</reset>\n")
		return nil
	}

	for _, pipe := range args {
		util.Printf("<yellow>stop job consumption for</reset> <white+hb>%s</reset><yellow>: </reset>", pipe)

		var r string
		if err := client.Call("jobs.Destroy", pipe, &r); err != nil {
			return err
		}

		util.Printf("<green+hb>done</reset>\n")
	}

	return nil
}
