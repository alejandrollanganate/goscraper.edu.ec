/*
Copyright © 2021 Club de Software EPN

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"

	"github.com/Club-de-Software-EPN/goscraper.edu.ec/cmd"
	scraping "github.com/Club-de-Software-EPN/goscraper.edu.ec/scraping/ESPE"
	"github.com/Club-de-Software-EPN/goscraper.edu.ec/sites"
)

func main() {
	cmd.Execute()
	fmt.Println(sites.GetVulnerableCollageURLs()["ESPE"])
	scraping.GetResearcherAndTeachersInfo()
}
