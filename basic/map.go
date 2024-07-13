package basic

import (
	"fmt"

)

func TestMap() {
	fmt.Println("----------------------------------------Start Map----------------------------------------")
	
	fmt.Println("                                    ***********             A            *********\n")
	myMap := make(map[int]string)
	 //به شکل زیر هم میشود تعریف کرد
	// var myMap map[int]string = map[int]string{}
	// اگر مقداردهی نکنیم با خطای nil مواجه خواهیم شد
	// var myMap map[int]string
	myKey := 10
	myMap[myKey] = "ten"
	fmt.Println(myMap)        //map[13:thirteen]
	fmt.Println(myMap[myKey]) //thirteen

	fmt.Println("                                    ***********             B            *********\n")
	var sampleMap = map[string]bool{}
	var otherMap = make(map[string]uint)
	var nilMap map[bool]bool
	sampleMap["condition#1"] = true
	sampleMap["condition#2"] = false
	otherMap["foo"] = 1
	fmt.Println(len(sampleMap))		//2
	fmt.Println(len(otherMap))		//1
	fmt.Println(len(nilMap))		//0 (len nil is zero)

	fmt.Println("                                    ***********             C            *********\n")

	var personData = map[string]string{"name": "frank", "family": "colleti"}

	name, nameExist := personData["name"]
	organization, organizationExist := personData["organization"]

	fmt.Println(name, nameExist) //frank true
	fmt.Println(organization, organizationExist) // false

	fmt.Println("                                    ***********             D            *********\n")
	var companyProfile = map[string]string{
		"key1": "val1",
		"key2": "val2",
	}

	var editorMap = companyProfile
	fmt.Println(companyProfile["key1"], "\t", companyProfile["key2"]) //val1 	val2
	fmt.Println(editorMap["key1"], "\t", editorMap["key2"]) //val1 	val2

	editorMap["key1"] = "new val1"
	editorMap["key2"] = "new val2"

	fmt.Println(companyProfile["key1"], "\t", companyProfile["key2"]) //new val1   new val2
	fmt.Println(editorMap["key1"], "\t", editorMap["key2"]) //new val1 	 new val2

	fmt.Println("                                    ***********             E            *********\n")
	var companyProfile = map[string]string{
		"name":    "companyName",
		"address": "sampleAddress",
	}

	var editorMap = map[string]string{}

	//copy from map to another map
	for key, value := range companyProfile {
		editorMap[key] = value
	}	

	fmt.Println(companyProfile["name"], "\t", companyProfile["address"])
	fmt.Println(editorMap["name"], "\t", editorMap["address"])

	editorMap["name"] = "new address"
	editorMap["address"] = "new address"


	fmt.Println(companyProfile["name"], "\t", companyProfile["address"])
	fmt.Println(editorMap["name"], "\t", editorMap["address"])

	fmt.Println("                                    ***********             F            *********\n")
	fmt.Println("----------------------------------------End Map  ----------------------------------------")
}
