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
	"github.com/spf13/cobra"
	"fmt"
	"github.com/spf13/viper"
	"github.com/fatih/color"
	"github.com/teo/zcascade/app"

	"time"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:   "about",
	Aliases: []string{},
	Short: fmt.Sprintf("about %s", app.NAME),
	Long: `The about command shows some basic information on this utility.`,
	Run: func(*cobra.Command, []string) {
		color.Set(color.FgHiWhite)
		fmt.Print(app.PRETTY_SHORTNAME + " *** ")
		color.Set(color.FgHiGreen)
		fmt.Printf("%s\n", app.PRETTY_FULLNAME)
		color.Unset()
		fmt.Printf(`
version:         %s
config:          %s
`,
			color.HiGreenString(viper.GetString("version")),
			color.HiGreenString(func() string {if len(viper.ConfigFileUsed()) > 0 { return viper.ConfigFileUsed() }; return "builtin"}()), )

		color.Set(color.FgHiBlue)
		fmt.Printf("\nCopyright 2020-%d Teo Mrnjavac.\n" +
			"This program is free software: you can redistribute it and/or modify \n" +
			"it under the terms of the GNU General Public License as published by \n" +
			"the Free Software Foundation, either version 3 of the License, or \n" +
			"(at your option) any later version.\n", time.Now().Year())
		color.Unset()

		fmt.Printf(`
bugs:            %s
code:            %s
maintainer:      %s
`,

			color.HiBlueString("https://github.com/teo/zcascade/issues"),
			color.HiBlueString("https://github.com/teo/zcascade"),
			color.HiBlueString("Teo Mrnjavac <teo.mrnjavac@gmail.com>"))
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
