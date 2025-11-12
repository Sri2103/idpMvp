# generating swagger docs for the service

param (
    [string]$Path
)



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





function Set-GoSwagClient {
    &swagger generate client -c apiGateway -f ./internal/api-gateway/v1/docs/swagger.yaml -t ./pkg/api/generated/api-gateway/client
}


# function Set-GoSwagCLI {
#     &swagger generate cli -c apiGateway  --cli-app-name apiGatewayCli  -f ./internal/api-gateway/v1/docs/swagger.yaml -t ./pkg/api/generated/api-gateway/client

# }
function Set-GoSwagCLI {
    &swagger generate cli -A apiGateway  -f ./internal/api-gateway/v1/docs/swagger.yaml -t ./pkg/api/generated/api-gateway

}

function Set-GetSwagServer {
    &swagger generate server -s server -f ./internal/api-gateway/v1/docs/swagger.yaml -t ./pkg/api/generated/api-gateway/server
}

function Set-Swag {
    Set-ApiGatewaySwagMultiDir
    Set-GoSwagCLI
}

Set-Swag





