#!/bin/bash
#This script automatically installs the app
#Build the executable first
echo "Building application...."
go build defoldAtlasConv.go && echo "Finished building application"
#Install it in the user bin folder so the app can be run from terminal in any directory
echo "Moving application to /usr/local/bin folder..."
sudo mv defoldAtlasConv /usr/local/bin/defoldAtlasConv && echo "Installation completed"