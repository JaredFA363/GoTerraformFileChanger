package main

import (
    "fmt"
    "os"
    "regexp"
    "github.com/hashicorp/hcl/v2/hclwrite"
    hcl "github.com/hashicorp/hcl/v2"
    "github.com/zclconf/go-cty/cty"
)

func getImageRunnerName(file_content string) string {
    fmt.Println(file_content)
    runner_image_pattern := `"([0-9]{8}.[0-9]{4}){1}"`
    runner_image_regex := regexp.MustCompile(runner_image_pattern)

    runner_image_match := runner_image_regex.FindStringSubmatch(file_content)
    fmt.Println(runner_image_match)

    if len(runner_image_match) > 1 {
        return runner_image_match[1]
    }
    return "No matches found"
}

func main() {
    path := "/home/runner/work/cc-spikes/cc-spikes/runner-automation/main.tf"

    // Load the current terrafrorm file
    file, err := os.ReadFile(path)
    if err != nil {
        panic("Step 1: " + err.Error())
    }
    fmt.Println(string(file))

    image_runner_name := getImageRunnerName(string(file))
    fmt.Println("Image runner name:", image_runner_name)

    // Parse the file into an hclwrite.File object and get the diagnostics
    hclFile, diags := hclwrite.ParseConfig(file, "main.tf", hcl.Pos{Line: 1, Column: 1})
    if diags.HasErrors() {
        panic("Step 2: failed to parse HCL file:" + diags.Error())
    }
