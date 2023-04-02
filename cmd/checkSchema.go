/*
Copyright © 2023 Alex Bechanko

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
)

var checkSchemaInputSchemaPath string

var checkSchemaCmd = &cobra.Command{
	Use:   "check-schema",
	Short: "Check if the schema document is a valid JSON schema",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		inputSchemaAbsPath, err := filepath.Abs(checkSchemaInputSchemaPath)
		if err != nil {
			fmt.Printf("Error occurred getting absolute path for schema %s: %s\n", checkSchemaInputSchemaPath, err)
			os.Exit(1)
		}
		inputSchemaURI := fmt.Sprintf("file://%s", inputSchemaAbsPath)

		schemaLoader := gojsonschema.NewSchemaLoader()
		schemaLoader.Validate = true

		err = schemaLoader.AddSchemas(gojsonschema.NewReferenceLoader(inputSchemaURI))

		if err != nil {
			fmt.Printf("Validation of schema %s: FAIL\n", inputSchemaAbsPath)
			fmt.Printf("%s", err)
		} else {
			fmt.Printf("Validation of schema %s: PASS\n", inputSchemaAbsPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkSchemaCmd)

	checkSchemaCmd.Flags().StringVar(&checkSchemaInputSchemaPath, "schema", "", "Schema that will be validated")
	checkSchemaCmd.MarkFlagRequired("schema")
}
