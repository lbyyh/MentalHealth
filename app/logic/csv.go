package logic

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

func CountColumns(c *gin.Context) {
	columns := c.PostForm("columns")

	// 打开指定的 CSV 文件
	file, err := os.Open("D:/code/GO/MentalHealth-Platform/app/pymodel/Mental Health Dataset.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	count, _ := readCSV(file, strings.Split(columns, ","))
	c.JSON(http.StatusOK, count)
}

func readCSV(file *os.File, columns []string) (map[string]map[string]int, error) {
	r := csv.NewReader(file)

	header, err := r.Read()
	if err != nil {
		return nil, err
	}

	idx := make(map[string]int)
	for i, name := range header {
		for _, column := range columns {
			if strings.EqualFold(name, column) {
				idx[column] = i
			}
		}
	}

	result := make(map[string]map[string]int)
	for column := range idx {
		result[column] = make(map[string]int)
	}

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		for column, i := range idx {
			result[column][line[i]]++
		}
	}

	return result, nil
}
