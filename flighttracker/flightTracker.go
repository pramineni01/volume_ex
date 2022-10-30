package flighttracker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Server struct{}

func (s Server) Calculate(ctx echo.Context) error {
	var req CalculateJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		log.Fatal(err)
	}
	log.Printf("Server::Calculate(): req: %v\n", req)
	fr, err := calculateImpl(req)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "")
	}

	data, _ := json.Marshal(fr)
	return ctx.String(http.StatusOK, string(data))
}

func calculateImpl(inp []FlightRoute) (FlightRoute, error) {
	src := make(map[string]interface{})
	dst := make(map[string]interface{})
	for _, a := range inp {
		_, foundInSrc := src[a[1]]
		_, foundInDst := dst[a[0]]
		if (foundInSrc == false) && (foundInDst == false) {
			src[a[0]] = nil
			dst[a[1]] = nil
			continue
		}

		if foundInSrc && foundInDst {
			delete(src, a[1])
			delete(dst, a[0])
			continue
		}

		if foundInSrc {
			delete(src, a[1])
			src[a[0]] = nil
		}
		if foundInDst {
			delete(dst, a[0])
			dst[a[1]] = nil
		}
	}
	if (len(src) != 1) && (len(dst) != 1) {
		fmt.Println("No viable answer")
		return []string{"", ""}, echo.ErrBadRequest
	}

	// extract from sets
	var s, d string
	for s, _ = range src {
	}
	for d, _ = range dst {
	}
	log.Println("Returning: ", []string{s, d})
	return []string{s, d}, nil
}
