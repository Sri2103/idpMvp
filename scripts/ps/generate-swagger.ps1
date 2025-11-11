# generating swagger docs for the service

param (
    [string]$Path
)


function Invoke-Swagger {


    # & docker run --rm -v $(Get-Location):/code ghcr.io/swaggo/swag:latest 

    $fullPath = Join-Path $(Get-Location) -ChildPath $Path
    Write-Host $fullPath
    
}



function Set-ApiGatewaySwag {
    param(
        [string]$main = "cmd/api-gateway/main.go",
        [string]$swagOut = "./internal/api-gateway/v1/docs/",
        [string] $genconfig = "./pkg/api/generated/api-gateway/config.yaml",
        [string] $apiConfig = "./internal/api-gateway/v1/docs/swagger.yaml" 
    )


    & swag.exe i -g $main -o $swagOut
    & oapi-codegen.exe -config $genconfig $apiConfig
}

