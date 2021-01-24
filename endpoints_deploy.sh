#!/bin/sh

gcloud endpoints services deploy ./proto/api_descriptor.pb ./proto/api_config.yaml
