# Touchbase

[![Releases](https://img.shields.io/github/v/tag/TransformersAutoBots/touchbase.svg?label=latest)](https://github.com/TransformersAutoBots/touchbase/releases/latest)
[![GitHub issues](https://img.shields.io/github/issues/TransformersAutoBots/touchbase?color=red)](https://github.com/TransformersAutoBots/touchbase/issues)
[![GitHub forks](https://img.shields.io/github/forks/TransformersAutoBots/touchbase?color=violet)](https://github.com/TransformersAutoBots/touchbase/network)
[![License](https://img.shields.io/github/license/TransformersAutoBots/touchbase)](./LICENSE)

## Table of Contents
* [Introduction](#introduction)
* [Prerequisites](#prerequisites)
  * [Create GCP Project](#create-gcp-project)
  * [Enable Sheets and Gmail API](#enable-sheets-and-gmail-api)
  * [Create App Credentials](#create-app-credentials)
  * [Retrieve spreadsheet ID](#retrieve-spreadsheet-id)
  * [Spreadsheet Format](#spreadsheet-format)
* [Build Binary](#build-binary)
* [Getting Started](#getting-started)

## Introduction
Touchbase helps to connect with people, share your profile with short 
description about yourself and your resume/portfolio!

## Prerequisites
### Create GCP Project
1. Navigate to [GCP](https://console.cloud.google.com) and login using your 
   `Gmail` credentials
2. From the left navigation bar, select `IAM & Admin -> Manage Resources -> Create Project`
   and enter your desired project name and id
   
   ![Create GCP Project](https://media.giphy.com/media/xzWE8zydyhgWOmakdP/giphy.gif)

### Enable Sheets and Gmail API
1. Now from the left navigation bar, select `APIs & Services -> Enable APIs and Services`
2. In the search bar, search for following APIs one by one and enable it 
   1. Google Sheets API
   2. Gmail API
   
   ![Enable Sheets and Gmail APIs](https://media.giphy.com/media/ecP7K6SsGmnqHLtvPI/giphy.gif)

### Create App Credentials
1. Again from the left navigation bar, select `APIs & Services -> OAuth Consent Screen`.
   Select `External` and add the required details i.e. your `App Name and User support email`
2. Now again select `APIs & Services -> Credentials -> Create Credentials`.
   After creating the `OAuth Client ID` download the json format and remember 
   the saved file location.    
   
   ![Create App Credentials](https://media.giphy.com/media/jBWtwSYFtoJ00Twxyb/giphy.gif)

### Retrieve spreadsheet ID
1. In your Google Drive create a Google Sheets and retrieve its spreadsheet id.
   For Reference on how to retrieve spreadsheet id [Click here](https://developers.google.com/sheets/api/guides/concepts#spreadsheet_id)

### Spreadsheet Format
1. The sheets same should be the companies name
2. The first row of each sheet should have 3 columns
   
   |First Name | Last Name | Email|
   |-----------|-----------|------|
3. Sample Spreadsheet
   
   ![Sample Spreadsheet](https://media.giphy.com/media/Py0Uolt9MCo7pqEtkJ/giphy.gif)

## Build Binary
### Download latest
You can download the latest [Release](https://github.com/TransformersAutoBots/touchbase/releases/latest)

### Build from source code
```
go build
go install
```

### Getting Started
```
touchbase --help
  _                            _       _
 | |                          | |     | |
 | |_    ___    _   _    ___  | |__   | |__     __ _   ___    ___
 | __|  / _ \  | | | |  / __| | '_ \  | '_ \   / _` | / __|  / _ \
 | |_  | (_) | | |_| | | (__  | | | | | |_) | | (_| | \__ \ |  __/
  \__|  \___/   \__,_|  \___| |_| |_| |_.__/   \__,_| |___/  \___|
Touchbase helps to connect with people, share your profile with short
description about yourself and your resume/portfolio!

Usage:
  touchbase [command]

Available Commands:
  config      Config required for touchbase application
  help        Help about any command
  reach-out   Reach out to the company recruiters/managers

Flags:
  -X, --debug   Enable debug mode (default false)
  -h, --help    help for touchbase

Use "touchbase [command] --help" for more information about a command.
```

For all the commands you need to export two env variables
```
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/credntials
export TOUCHBASE_CONFIG_DIR=/path/to/config/dir
```
1. GOOGLE_APPLICATION_CREDENTIALS - This is the path to the credentials file created in [Prerequisites steps - Create app credentials](#create-app-credentials)
2. TOUCHBASE_CONFIG_DIR - Place where you want to store the touchbase application config

#### App config init
```
touchbase config init --email=<gmail_id@gmail.com> --full-name=<your_full_name> --spreadsheet-id=<spreadsheetid> --resume=/path/to/resume_file
```
**Note:**
  1. Email must be a gmail account
  2. spreadsheet-id is the spreadsheet id created in [Prerequisites steps - Retrieve spreadsheet id](#retrieve-spreadsheet-id)
  3. Resume file path must be a valid file path. (Only pdf format allowed for now)
  4. After config init successful run, add a file named `introduce.html` in your
     TOUCHBASE_CONFIG_DIR. This file will contain the body of the email that 
     you want to send to Recruiter/manager. [Click here for Sample](./templates/introduce.html)

E.g:
```
touchbase config init --email="test@gmail.com" --full-name="FirstName LastName" --spreadsheet-id="1234567890abcdefgh" --resume=/Users/testuser/Desktop/resume.pdf
```
```json5
{
  "spreadsheetid": "1234567890abcdefgh",
  "user": {
    "fullname": "FirstName LastName",
    "emailid": "test@gmail.com",
    "resume": "/Users/testuser/Desktop/resume.pdf"
  }
}
```
#### config update
```
touchbase config update --key=<key_to_update> --value=<updated_value_for_key>
```
E.g:
```
touchbase config update --key=user.resume --value=/Users/testuser/Desktop/new_resume.pdf
```
#### reach-out
```
touchbase reach-out
```

Select the Company from the list and enter the start and end row
#### Enable debug mode
Add -X at the end of any command to enable debug mode 
```
touchbase reach-out -X
```


