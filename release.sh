#!/bin/bash

# Prior to running the release script install tfplugindocs
#  - https://github.com/hashicorp/terraform-plugin-docs#installation
# Next run the following command to release a new version 
#  - ./release.sh v<release number> 

version=${1}
echo ${version}
if [ -z "${version}" ]
then
    echo "version was left empty"
else
    tfplugindocs generate
    cp -r guides docs/
    git add .
    git commit -m "kps client install"
    git tag -a ${version} -m "Release version ${version}" main
    git push origin ${version}
    echo "pushed release version ${version}"
fi