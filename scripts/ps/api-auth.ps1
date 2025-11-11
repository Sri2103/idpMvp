function Set-ApiAuthMultDir {
    param(
        [string]$main = "main.go",
        [string]$dirs = "./cmd/auth-service,./internal/auth-service/v1/service",
        [string]$swagOut = "./internal/auth-service/v1/docs"
    )

    $dirArray = $dirs -split ','

    # Join the array back to a comma-separated string for the swag command (in case trimming is needed)
    $dirList = $dirArray -join ','

    & swag.exe i -g $main -d $dirList -o $swagOut
}

function Set-AuthCodeGen {
    &oapi-codegen.exe -config ./pkg/api/generated/auth-service/config.yaml ./internal/auth-service/v1/docs/swagger.yaml
}


Set-ApiAuthMultDir
Set-AuthCodeGen