package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Folder struct {
	Backet  string `json:"backet"`
	Rfolder string `json:"s3folder"`
	Lfolder string `json:"localfolder"`
}

func get_file(bucket string, item string) {
	fmt.Println("getting files\n")
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)
	file, err := os.Create("/tmp/" + item)
	if err != nil {
		fmt.Println("[MYERROR] is : ", err.Error())
	}
	// 	exitErrorf("Unable to open file %q, %v", err)
	// 	fmt.Println(exitErrorf)
	// }

	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	//if err != nil {
	//	exitErrorf("Unable to download item %q, %v", item, err)
	//}
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

}

func main() {

	content, err := ioutil.ReadFile("folders.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	var folders []Folder
	// json.Unmarshal(content, &friends)
	err2 := json.Unmarshal(content, &folders)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
	for _, x := range folders {
		fmt.Printf("%s %s  \n", x.Backet, x.Rfolder)
		dd := x.Backet
		ff := x.Rfolder
		//cc := x.Lfolder
		fmt.Println(" start copying from: ", dd)
		get_file(dd, ff)

	}
}
