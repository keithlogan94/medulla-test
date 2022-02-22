package main

import (
	"bytes"
	"fmt"
	hlib "html"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type Component struct {
	Name     string
	Content  string
	Category string
}

type ComponentCategoryItem struct {
	Name      string
	Component string
}

type ComponentCategory struct {
	Category string
	Content  string
}

func GetComponentLinks() [][]byte {
	primerComponentsHtml := GetHtmlAt(ExternalComponentsLibrariesPrimer)
	links := GetLinks(primerComponentsHtml)
	links = FilterList(links, "/css/components/")
	links = PrependListItemsWith(links, []byte(ExternalComponentsLibrariesPrimerBase))
	return links
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetCategoryFromLink(link string) string {
	link = link[strings.LastIndex(link, "/")+1:]
	link = strings.Replace(link, "-", " ", -1)
	link = strings.Title(link)
	return link
}

func GetComponents() []Component {
	links := GetComponentLinks()

	var recordComponent bool = false
	var recordName bool = false
	var captureComponent []byte
	var captureName []byte
	var components [][]byte
	var names [][]byte
	var foundComponents []Component
	var currentName []byte
	var currentCategory []byte

	for i := 0; i < len(links); i++ {
		html := GetHtmlAt(string(links[i]))
		currentCategory = []byte(GetCategoryFromLink(string(links[i])))

		for htmlCharIndex := 0; htmlCharIndex < len(html); htmlCharIndex++ {
			if len(html)-htmlCharIndex > 10 {
				if recordComponent == false && recordName == false &&
					html[htmlCharIndex] == '<' &&
					html[htmlCharIndex+1] == 't' &&
					html[htmlCharIndex+2] == 'e' &&
					html[htmlCharIndex+3] == 'x' &&
					html[htmlCharIndex+4] == 't' &&
					html[htmlCharIndex+5] == 'a' &&
					html[htmlCharIndex+6] == 'r' &&
					html[htmlCharIndex+7] == 'e' &&
					html[htmlCharIndex+8] == 'a' {
					for recordComponent == false {
						if html[htmlCharIndex] == '>' {
							htmlCharIndex++
							recordComponent = true
							continue
						} else {
							htmlCharIndex++
						}
					}
				}
				if recordComponent {
					if html[htmlCharIndex] == '<' &&
						html[htmlCharIndex+1] == '/' &&
						html[htmlCharIndex+2] == 't' &&
						html[htmlCharIndex+3] == 'e' &&
						html[htmlCharIndex+4] == 'x' &&
						html[htmlCharIndex+5] == 't' &&
						html[htmlCharIndex+6] == 'a' &&
						html[htmlCharIndex+7] == 'r' &&
						html[htmlCharIndex+8] == 'e' &&
						html[htmlCharIndex+9] == 'a' {
						recordComponent = false
						components = append(components, []byte(hlib.UnescapeString(string(captureComponent))))
						foundComponents = append(foundComponents, Component{Name: string(currentName), Category: string(currentCategory), Content: hlib.UnescapeString(string(captureComponent))})
						captureComponent = []byte{}
						continue
					}
					captureComponent = append(captureComponent, html[htmlCharIndex])
				} else {
					if recordName == false &&
						html[htmlCharIndex] == '<' &&
						html[htmlCharIndex+1] == 'h' &&
						(html[htmlCharIndex+2] == '4' || html[htmlCharIndex+2] == '2' || html[htmlCharIndex+2] == '3' || html[htmlCharIndex+2] == '5' || html[htmlCharIndex+3] == '1' || html[htmlCharIndex+3] == '6') {
						for recordName == false {
							if html[htmlCharIndex] == '>' {
								htmlCharIndex++
								recordName = true
								continue
							} else {
								htmlCharIndex++
							}
						}
					}
					if recordName {
						if html[htmlCharIndex] == '<' &&
							html[htmlCharIndex+1] == '/' &&
							html[htmlCharIndex+2] == 'h' &&
							(html[htmlCharIndex+3] == '4' || html[htmlCharIndex+3] == '2' || html[htmlCharIndex+3] == '3' || html[htmlCharIndex+3] == '5' || html[htmlCharIndex+3] == '1' || html[htmlCharIndex+3] == '6') {
							recordName = false

							index := strings.Index(string(captureName), "</a>") + len("</a>")
							captureName = captureName[index:]

							names = append(names, []byte(string(captureName)))
							currentName = captureName
							captureName = []byte{}
							continue
						}
						captureName = append(captureName, html[htmlCharIndex])
					}
				}

			}
		}
	}

	var newNames [][]byte
	for i := 0; i < len(names); i++ {
		name := names[i]
		index := strings.Index(string(name), "</a>") + len("</a>")
		name = name[index:]
		newNames = append(newNames, name)
	}

	return foundComponents
}

func main() {
	components := GetComponents()
	fmt.Println("found " + strconv.Itoa(len(components)) + " components")

	fmt.Println("deleting tmp directory if exists")
	os.RemoveAll("tmp")
	err := os.Mkdir("tmp", 0755)
	CheckError(err)
	fmt.Println("created tmp directory for generating files")

	fmt.Println("creating directory structure")
	for i := 0; i < len(components); i++ {
		var dir string = strings.Replace(components[i].Category, " ", "", -1)
		os.Mkdir("tmp/"+dir, 777)
	}

	content, err := os.ReadFile("component.txt")
	CheckError(err)

	tmpl, err := template.New("component").Parse(string(content))
	CheckError(err)

	var uniqueDirs []string

	for i := 0; i < len(components); i++ {
		if !contains(uniqueDirs, strings.Replace(components[i].Category, " ", "", -1)) {
			uniqueDirs = append(uniqueDirs, strings.Replace(components[i].Category, " ", "", -1))
		}
		var buffer bytes.Buffer
		err = tmpl.Execute(&buffer, components[i])
		var fileContent string = buffer.String()
		CheckError(err)
		err := os.WriteFile("tmp/"+strings.Replace(components[i].Category, " ", "", -1)+"/"+strings.Replace(components[i].Category, " ", "", -1)+strconv.Itoa(i)+".razor", []byte(fileContent), 0777)
		CheckError(err)

	}

	componentCategoryItemContent, err := os.ReadFile("component-category-item.txt")
	CheckError(err)
	componentCategoryItemTmpl, err := template.New("component-category-item").Parse(string(componentCategoryItemContent))
	CheckError(err)
	componentCategoryContent, err := os.ReadFile("component-category.txt")
	CheckError(err)
	componentCategoryTmpl, err := template.New("component-category").Parse(string(componentCategoryContent))
	CheckError(err)

	var componentCategories string
	for i := 0; i < len(uniqueDirs); i++ {
		fmt.Println(uniqueDirs[i])
		var componentCategoryItems string
		for componentIndex := 0; componentIndex < len(components); componentIndex++ {
			if strings.ReplaceAll(components[componentIndex].Category, " ", "") == strings.ReplaceAll(uniqueDirs[i], " ", "") {
				fmt.Println("categories are equal")
				var buffer bytes.Buffer
				var cci ComponentCategoryItem = ComponentCategoryItem{components[componentIndex].Name, "Medulla.Frontend.Client.Components.RegisteredComponents." + uniqueDirs[i] + "." + uniqueDirs[i] + strconv.Itoa(componentIndex)}
				err = componentCategoryItemTmpl.Execute(&buffer, cci)
				CheckError(err)
				var value string = buffer.String()
				componentCategoryItems += value
			}
		}
		var buffer bytes.Buffer
		var componentCategory ComponentCategory = ComponentCategory{Category: uniqueDirs[i], Content: componentCategoryItems}
		err = componentCategoryTmpl.Execute(&buffer, componentCategory)
		CheckError(err)
		componentCategories += buffer.String()
	}

	err = os.WriteFile("tmp/component-categories.txt", []byte(componentCategories), 0777)
	CheckError(err)

}
