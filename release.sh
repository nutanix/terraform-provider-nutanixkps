#!/bin/bash

version=${1}
echo ${version}
if [ -z "${version}" ]
then
    echo "version was left empty"
else
    git tag -a ${version} -m "Release version ${version}" main
    git push origin ${version}
    echo "pushed release version ${version}"
fi