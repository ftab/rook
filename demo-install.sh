#!/bin/bash

# Disable default CSI driver
# Enable flexvolume driver
# as of a month ago, the version of cephcsi in use by rook did not support volume expansion, but flexvolume driver has since v1.1
# rook may have been updated to have cephcsi v2.0.0 by now which does support resize. if so, that might be a better bet

helm repo add rook-release https://charts.rook.io/release
helm install --name rook-ceph --namespace rook-ceph --version v1.2.4 rook-release/rook-ceph --set csi.enableRbdDriver=false --set csi.enableCephfsDriver=false --set enableFlexDriver=true
