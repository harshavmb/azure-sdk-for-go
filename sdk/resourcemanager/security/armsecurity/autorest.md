### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
azure-arm: true
tag: package-composite-v3
require:
- https://github.com/Azure/azure-rest-api-specs/blob/6080bb1fc0e219b72ed3c85966b54334e22e9980/specification/security/resource-manager/readme.md
- https://github.com/Azure/azure-rest-api-specs/blob/6080bb1fc0e219b72ed3c85966b54334e22e9980/specification/security/resource-manager/readme.go.md
license-header: MICROSOFT_MIT_NO_VERSION
module-version: 0.8.0
modelerfour:
  lenient-model-deduplication: true
```