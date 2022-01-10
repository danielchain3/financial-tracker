package main

import (
 	"context"
 	"encoding/json"
 	"fmt"
 	"io/ioutil"
 	"log"
 	"net/http"
 	"os"

 	"golang.org/x/oauth2"
 	"golang.org/x/oauth2/google"
 	"google.golang.org/api/sheets/v4"
 	"google.golang.org/api/option"

)

func createSheet(title_name string) {
	ss, err := service.CreateSpreadSheet(spreadsheet.Spreadsheet{
		Properties: spreadsheet.Properties{
			Title: title_name,
		},
	})
}


