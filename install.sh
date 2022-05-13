#!/bin/bash
#Export all environment variables
echo "Exporting environment variables...."
source <(sed -E -n 's/[^#]+/export &/ p' .env)
echo "Finished exporting environment variables...."
echo "....................."
 #Build the executable first
 ./build.sh
 #Install it in the user bin folder so the app can be run from terminal in any directory
 cp $APP_NAME ~/bin/$APP_NAME