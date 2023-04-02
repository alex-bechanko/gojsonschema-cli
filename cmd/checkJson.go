/*
Copyright © 2022 Alex Bechanko alexbechanko@gmail.com

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

var (
	checkJSONInputDocumentPath string
	checkJSONInputSchemaPath   string
)

var checkJSONCmd = &cobra.Command{
	Use:   "check-json",
	Short: "Checks if a JSON document is valid according to a JSON schema",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		inputDocumentAbsPath, err := filepath.Abs(checkJSONInputDocumentPath)
		if err != nil {
			fmt.Printf("Error occurred getting absolute path for document %s: %s\n", checkJSONInputDocumentPath, err)
			os.Exit(1)
		}
		inputDocumentURI := fmt.Sprintf("file://%s", inputDocumentAbsPath)
		documentLoader := gojsonschema.NewReferenceLoader(inputDocumentURI)

		inputSchemaAbsPath, err := filepath.Abs(checkJSONInputSchemaPath)
		if err != nil {
			fmt.Printf("Error occurred getting absolute path for file %s: %s\n", checkJSONInputSchemaPath, err)
			os.Exit(1)
		}
		inputSchemaURI := fmt.Sprintf("file://%s", inputSchemaAbsPath)
		schemaLoader := gojsonschema.NewReferenceLoader(inputSchemaURI)

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			fmt.Printf(
				"Error running validator with schema %s on document %s: %s\n",
				inputSchemaAbsPath,
				inputDocumentAbsPath,
				err,
			)
			os.Exit(1)
		}

		if !result.Valid() {
			fmt.Printf("Validation of document %s with schema %s: FAIL\n", inputDocumentAbsPath, inputSchemaAbsPath)
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}
			os.Exit(1)
		}

		fmt.Printf("Validation of document %s with schema %s: PASS\n", inputDocumentAbsPath, inputSchemaAbsPath)
	},
}

func init() {
	rootCmd.AddCommand(checkJSONCmd)
	checkJSONCmd.Flags().StringVar(&checkJSONInputDocumentPath, "document", "", "Document that will be validated")
	checkJSONCmd.Flags().StringVar(&checkJSONInputSchemaPath, "schema", "", "Schema that will be used for validation")

	checkJSONCmd.MarkFlagRequired("document")
	checkJSONCmd.MarkFlagRequired("schema")
}
