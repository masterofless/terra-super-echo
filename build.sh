#!/bin/zsh

builder=nerdctl # docker
DIR=$(git rev-parse --show-toplevel)
cd $DIR/images

for img in *
do
    TAG="terra-${img}:latest"
    set -x
    $builder build -t $TAG $img
    set +x
done
$builder images
