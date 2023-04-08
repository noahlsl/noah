package csvx

import (
	"bytes"
	"fmt"
	"net/http"
)

func WriteCsv(w http.ResponseWriter, b *bytes.Buffer, name string) {

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename="+name+".csv"))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b.Bytes())
}
