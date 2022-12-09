#!/bin/zsh

builder=nerdctl # docker
DIR=$(git rev-parse --show-toplevel)
cd $DIR/images

for img in *
do
    TAG="terra-${img}:latest"
    set -x
    $builder build -t $TAG --namespace k8s.io $img
    set +x
done
$builder images
