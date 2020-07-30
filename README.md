# AppStoreConnect-OpenApi-Go

This is a generated Go project for Apple's [App Store Connect API](https://developer.apple.com/documentation/appstoreconnectapi).

Apple provided an OpenAPI
[downloable specification](https://developer.apple.com/sample-code/app-store-connect/app-store-connect-openapi-specification.zip) which allows use to generate client api codes.

Client code generation was done with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator)

## Motivation

When I was generating the code for my own use, I experience some difficulties with OpenAPI Generator.
After fixing all the bugs, I thought it would be helpful to share this generated code with others.

## Installation

Add the generated code to your project

```bash
go get github.com/xr1337/appstoreconnect-openapi-go/generated
```

## Usage

Heres a non-production recommended code to test the API out
```golang
import 	api "github.com/xr1337/appstoreconnect-openapi-go/generated"

func main() {
	cfg := api.NewConfiguration()
	auth := context.WithValue(context.Background(), api.ContextAccessToken, signedToken)
	client := api.NewAPIClient(cfg)
	response, _, _ := client.UsersApi.UsersGetCollection(auth, nil)
	for _, user := range response.Data {
		fmt.Println(user.Attributes.Username)
	}
}
```

Checkout out the [example folder](https://github.com/xr1337/appstoreconnect-openapi-go/tree/master/example) for more examples ( includes reading Apples .p8 file )

Alternatively, OpenAPI generates a README.md file ([View here](https://github.com/xr1337/appstoreconnect-openapi-go/blob/master/generated/README.md)) to see what APIs that you can use

## Regenerate

The Makefile comes with a example docker command to regenerate the project.
Requirements: [docker](https://www.docker.com/get-started)

```bash
make generate
```

The output will generate a new folder called `out`.
You may want to remove the generated go.mod and go.sum files.

## License
[MIT](https://choosealicense.com/licenses/mit/)
