# generating swagger docs for the service

param (
    [string]$Path
)


function Invoke-Swagger {


    # & docker run --rm -v $(Get-Location):/code ghcr.io/swaggo/swag:latest 

    $fullPath = Join-Path $(Get-Location) -ChildPath $Path
    Write-Host $fullPath
    
}

Invoke-Swagger