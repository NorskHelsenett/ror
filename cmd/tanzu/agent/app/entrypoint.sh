#!/bin/sh
export KUBECONFIG=/config/kubeconfig
export PATH=$PATH:/app

/app/tanzuAgent
