# generating swagger docs for the service

param (
    [string]$Path
)


function Invoke-Swagger {


    # & docker run --rm -v $(Get-Location):/code ghcr.io/swaggo/swag:latest 

    $fullPath = Join-Path $(Get-Location) -ChildPath $Path
    Write-Host $fullPath
    
}


# function Set-ApiGatewaySwag {
#     param (
#         [string]$main = "main.go",
#         [string]$basePath = "./cmd/api-gateway/",
#         [string]$swagOut = "./internal/api-gateway/v1/docs/",
#         [string]$swagdir = "./internal/api-gateway/v1/handlers/",
#         [string]$genconfig = "./pkg/api/generated/api-gateway/config.yaml",
#         [string]$apiConfig = "./internal/api-gateway/v1/docs/swagger.yaml"
#     )

#     $list = @{
#         $basePath, $swagdir
#     }

#     & swag.exe i -g $main -d "./cmd/api-gateway", $swagdir -o $swagOut
#     & oapi-codegen.exe -config $genconfig $apiConfig
# }

function Set-ApiGatewaySwagMultiDir {
    param (
        [string]$main = "main.go",
        [string]$dirs = "./cmd/api-gateway,./internal/api-gateway/v1/handlers",
        [string]$swagOut = "./internal/api-gateway/v1/docs/"
    )

    # Convert the comma-separated string of directories to an array
    $dirArray = $dirs -split ','

    # Join the array back to a comma-separated string for the swag command (in case trimming is needed)
    $dirList = $dirArray -join ','

    & swag.exe i -g $main -d $dirList -o $swagOut
}

function Set-ApiGatewayCodeGen {
    &oapi-codegen.exe -config ./pkg/api/generated/api-gateway/config.yaml ./internal/api-gateway/v1/docs/swagger.yaml
}


Set-ApiGatewaySwagMultiDir
Set-ApiGatewayCodeGen




