# Zip and Unzip Files

# [1](https://golang.org/pkg/archive/zip/) `https://golang.org/pkg/archive/zip/`

I have written code from the source above.  :point_up: :point_up:

- First of all, open gofile `Zip.go` to zip a file then declare a path and write a name **.zip** file in function main.go.

    - Example: 
func main()  { 

	<em>`path`</em> zipit("C:/Users/KHUSRAV/go/src/ZipArchive/index.html", <em>`aim`</em> "C:/Users/KHUSRAV/go/src/ZipArchive/indexx.zip")

}


------------------------------------------------------------------------------------------------------------------------------


- If you want to unzip the file you must open gofile `Unzip.go`, and rename zipRead, _ := zip.OpenReader("NameOfTheZipFile.zip").


Where **"NameOfTheZipFile.zip"** in my case is **indexx.zip**
