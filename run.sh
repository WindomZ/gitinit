#!/usr/bin/env bash
go install

echo "------------------"
gitinit -h
#gitinit --help
#gitinit -v
#gitinit --version
echo "------------------"
gitinit --init
echo "------------------"
gitinit --bare
echo "------------------"
gitinit repo
echo "------------------"
