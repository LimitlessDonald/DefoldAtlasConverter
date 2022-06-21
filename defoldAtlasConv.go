package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	atlasFile:=flag.String("atlas","","The name of the atlas file with the extension, e.g images.atlas")
	//tableFlag:=flag.String("table","mytable","The name you would like to give your table variable, which would also be the name of the .lua file ")
	//jsonFlag:=flag.String("json","myjson","The name you would like to give your JSON file")
	outputFile:=flag.String("output","output.lua","The file name and extension you would like to save the file as \n For example myfile.lua myfile.json")
	//flag.Usage()
	flag.Parse()

	if strings.Contains(*atlasFile,".atlas"){


		if strings.Contains(*outputFile,".lua") {
			//	This also works If no output file is provided , saves it as output.lua, a lua table
			toLuaTable(*atlasFile, *outputFile)

		}else if strings.Contains(*outputFile,".json"){
			toJSON(*atlasFile,*outputFile)
		}else{
			// if just a file name is provided without an extension, default to .lua
			toLuaTable(*atlasFile, *outputFile+".lua")
		}
	} else{
		fmt.Println("Please include a valid .atlas file")
		flag.Usage()
	}

	//print(*atlasFileFlag)
	//print(*tableFlag)
}

func fileToString(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		print(err)

	}
	return string(b)
}

//toJSON converts the provided string , in this case a defold atlas file into JSON text file
// fileName is the name of the .atlas file INCLUDING THE EXTENSION
// JSONfile is the name of the .json file INCLUDING THE EXTENSION
func toJSON(fileName,JSONfile string)  {
	atlasString:= fileToString(fileName)

	// remove spaces
	atlasString=strings.ReplaceAll(atlasString," ","")
	//replace images{ with "images":
	atlasString=strings.ReplaceAll(atlasString,"images{","\"images\":{")

	atlasString=strings.ReplaceAll(atlasString,"image:","\"image\":")
	atlasString=strings.ReplaceAll(atlasString,"sprite_trim_mode:","\"sprite_trim_mode\":")
	atlasString=strings.ReplaceAll(atlasString,"SPRITE_TRIM_MODE_OFF","\"SPRITE_TRIM_MODE_OFF\"")
	atlasString=strings.ReplaceAll(atlasString,"margin:","\"margin\":")
	atlasString=strings.ReplaceAll(atlasString,"extrude_borders:","\"extrude_borders\":")
	atlasString=strings.ReplaceAll(atlasString,"inner_padding:","\"inner_padding\":")
	atlasString=strings.ReplaceAll(atlasString,"\"\n\"","\",\n\"")
	atlasString=strings.ReplaceAll(atlasString,"}\n\"","},\n\"")
	//append opening curly bracket for valid JSON
	atlasString="{\n"+atlasString+"\n}"
	//preparing to replace images with image1-2 etc
	//count number of images
	imagesCount := strings.Count(atlasString,"\"images\"")
	imagesCountLoop :=1

	for imagesCountLoop <= imagesCount {
		newImageObjectName := "\"image"+strconv.Itoa(imagesCountLoop) +"\""
		// Always replace the first occurrence
		atlasString=replaceNth(atlasString,"\"images\"", newImageObjectName,1)
		imagesCountLoop = imagesCountLoop + 1

	}
	// replace the last three properties, margin, extrude border and inner padding with regex and include a ","
	re := regexp.MustCompile(`(["]\w+["][:][\d])`)
	atlasString = re.ReplaceAllString(atlasString,`$1`+`,`)

	writeJSONerr := ioutil.WriteFile(JSONfile, []byte(atlasString),0644)
	if writeJSONerr != nil {
		fmt.Println("There was an error converting your atlas file to JSON")
		flag.Usage()
		return
	}
	fmt.Println("Your JSON file has been saved :"+JSONfile)

	//print(atlasString)
	//print("\n")


}

//toLuaTable converts the provided atlas file to a lua table, it actually extracts only the image names and converts it to a table
//
// For example if an atlas file contains the following
// images {
// image: "/images/cat.png"
// sprite_trim_mode: SPRITE_TRIM_MODE_OFF
// }
// images {
// image: "/images/dog.png"
// sprite_trim_mode: SPRITE_TRIM_MODE_OFF
// }
//
// It converts to a lua table like so
//todo It doesn't handle the image name properly if it starts with a number or contains "-"
func toLuaTable(fileName string, luaFile string ){
	// the quote string
	quote:="\""
	atlasString:= fileToString(fileName)
	// first let's trim all spaces
	atlasString = strings.ReplaceAll(atlasString," ", "")
	// regex to pick out , e.g image: "/images/dog.png"
	re := regexp.MustCompile(`\bimage\b[:]["].+["]`)
	//-1 indicates there is no limit to how many it should find
	arrayOfImages := re.FindAllString(atlasString,-1)
	//build the lua file content
	//init table name, by removing the extension from the name
	tableName:=strings.ReplaceAll(luaFile,".lua","")
	// init the file content with the name of the table and the beginning of the table
	luaFileString := tableName + " = {\n"

	for _, image := range arrayOfImages {
		//print("Image " + strconv.Itoa(i) + " " + image)
		//remove "
		image= strings.ReplaceAll(image,"\"","")
		// split the image object into an array seperated by the forward slash
		imageNameInit:= strings.Split(image,"/")
		// subtract one from the length of the array
		imageName:=imageNameInit[len(imageNameInit)-1]
		// remove the file extension from the image name since image name above produces something like cat.png or dog.jpeg
		imageExtRegex := regexp.MustCompile(`[.].+`)
		imageName = imageExtRegex.ReplaceAllString(imageName,"")
		// keys in table can't include an hypen "-" , so lets replace it with underscore it , if it has one
		tableKeyName := strings.ReplaceAll(imageName,"-","_")
		//print(strconv.Itoa(len(imageNameInit))  + " For Image "+strconv.Itoa(i) +" "+ imageName + "\n")
		luaFileString = luaFileString + "\t" + tableKeyName+"="+quote+imageName+quote+",\n"

	}
	luaFileString=luaFileString+"}\n"
	err := ioutil.WriteFile(luaFile, []byte(luaFileString), 0644)
	if err != nil {
		fmt.Println("There was an error converting your atlas file to lua table")
		flag.Usage()
		return
	}
	fmt.Println("Your lua file has been saved :"+luaFile)
	
}

func replaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

//func replacey(string,,index){
//	excludingLast := string[:i] + strings.Replace(mystring[i:], "optimism", "", 1)
//}