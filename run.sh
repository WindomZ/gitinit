#!/usr/bin/env bash
go install

echo "------------------"
gitinit -h
#gitinit --help
#gitinit -v
#gitinit --version
echo "------------------"
gitinit --init
gitinit --init --remote=xxx
echo "------------------"
gitinit --bare
echo "------------------"
