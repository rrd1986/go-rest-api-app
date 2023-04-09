# Build REST API  with GoFiber and Postgress and Docker Compose


## Requirements
`
go get github.com/gofiber/fiber/v2
`

## To create the fact entries use the following similar payload
`
method: POST
url: http://localhost:3000/fact
payload: 
{"question": "what is your age",
 "answer": "38"
}
`