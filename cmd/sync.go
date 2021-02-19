/*
 * Copyright 2021 Teo Mrnjavac <teo.mrnjavac@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/teo/zcascade/cmd/replication"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Aliases: []string{},
	Short: fmt.Sprintf("synchronize dataset to inventory"),
	Long: `The sync command performs a cascaded replication of a given dataset, including all its snapshots.`,
	Run: func(*cobra.Command, []string) {
		const nTargets = 100
		targets := make([]string, nTargets)
		for i := 0; i < nTargets; i++ {
			targets[i] = fmt.Sprintf("target-%d", i)
		}
		replication.Sync("mySource", targets)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
